from fastapi import FastAPI, UploadFile, File, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from book_loader import BookLoader
from rag_system import BookRecommendationSystem
from typing import Dict, Optional
import shutil
import os
import asyncio
from concurrent.futures import ThreadPoolExecutor

app = FastAPI()
executor = ThreadPoolExecutor()

# 配置CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# 全局变量存储推荐系统实例
recommendation_system = None

async def init_recommendation_system(file_path: str):
    """异步初始化推荐系统"""
    try:
        # 初始化图书加载器和推荐系统
        loader = BookLoader()
        documents = loader.load_books()
        
        global recommendation_system
        if recommendation_system is None:
            print("正在初始化推荐系统...")
            recommendation_system = BookRecommendationSystem()
            recommendation_system.initialize_vectorstore(documents)
        else:
            print("推荐系统已存在，正在重新初始化...")
            recommendation_system.initialize_vectorstore(documents)
            
        print("推荐系统初始化完成！")
        
    except Exception as e:
        print(f"初始化推荐系统时发生错误: {str(e)}")

@app.post("/admin/upload")
async def upload_books(file: UploadFile = File(...)):
    """管理员上传图书数据接口"""
    try:
        # 验证文件格式
        if not file.filename.endswith('.xlsx'):
            raise HTTPException(status_code=400, detail="只支持Excel文件(.xlsx)")
        
        # 保存上传的文件
        file_path = "./data/books.xlsx"
        os.makedirs("./data", exist_ok=True)
        
        # 如果文件已存在且内容相同，直接使用
        file_exists = os.path.exists(file_path)
        if not file_exists:
            # 保存新文件
            with open(file_path, "wb") as buffer:
                shutil.copyfileobj(file.file, buffer)
            print("新数据文件已保存")
        else:
            print("检测到已存在的数据文件")
        
        # 异步启动初始化过程
        asyncio.create_task(init_recommendation_system(file_path))
        
        return {
            "message": "文件上传成功，系统正在后台初始化",
            "status": "initializing"
        }
        
    except Exception as e:
        print(f"上传文件时发生错误: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/chat")
async def chat(request: Dict[str, str]):
    """用户聊天接口"""
    try:
        if not recommendation_system:
            raise HTTPException(status_code=500, detail="系统未初始化，请先上传图书数据")
        
        user_input = request.get("message")
        if not user_input:
            raise HTTPException(status_code=400, detail="消息不能为空")

        # 创建自定义回调处理器
        class OutputCollector:
            def __init__(self):
                self.output = []
                self.raise_error = False
                self.ignore_chat_model = False
                self.ignore_llm = False
                
            def on_llm_new_token(self, token: str, **kwargs):
                self.output.append(token)
                
            def on_llm_start(self, *args, **kwargs):
                pass
                
            def on_llm_end(self, *args, **kwargs):
                pass
                
            def on_llm_error(self, *args, **kwargs):
                pass
                
            def on_chain_start(self, *args, **kwargs):
                pass
                
            def on_chain_end(self, *args, **kwargs):
                pass
                
            def on_chain_error(self, *args, **kwargs):
                pass
                
            def on_chat_model_start(self, *args, **kwargs):
                pass

        collector = OutputCollector()
        
        # 设置回调处理器
        original_callbacks = recommendation_system.llm.callbacks
        recommendation_system.llm.callbacks = [collector]
        
        try:
            # 判断是否是图书推荐查询
            is_recommendation = recommendation_system.is_book_recommendation_query(user_input)
            
            # 根据查询类型处理消息
            if is_recommendation:
                response = recommendation_system.get_recommendation(user_input)
            else:
                response = recommendation_system.chat_prompt.pipe(recommendation_system.llm).invoke(
                    {"input": user_input}
                )
            
            # 获取收集到的输出
            response_text = "".join(collector.output) if collector.output else response
            
            return {"response": response_text}
            
        finally:
            # 恢复原始回调处理器
            recommendation_system.llm.callbacks = original_callbacks
            
    except Exception as e:
        print(f"聊天接口发生错误: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@app.get("/health")
async def health_check():
    """健康检查接口"""
    return {
        "status": "ok",
        "system_initialized": recommendation_system is not None
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=7001) 