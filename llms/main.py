import uvicorn
from api import app

if __name__ == "__main__":
    try:
        port = 7001  # 固定使用 7001 端口
        print(f"后端服务将启动在 http://127.0.0.1:{port}")
        
        uvicorn.run(
            "api:app",
            host="127.0.0.1",
            port=port,
            reload=True
        )
    except Exception as e:
        print(f"启动服务时发生错误: {str(e)}")
