from langchain_community.vectorstores import Chroma
from langchain_huggingface import HuggingFaceEmbeddings
from langchain.chat_models import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables import RunnablePassthrough, RunnableConfig
from langchain.callbacks.streaming_stdout import StreamingStdOutCallbackHandler
from sentence_transformers import SentenceTransformer
from config import *
import sys
from typing import Optional

class BookRecommendationSystem:
    def __init__(self):
        # 初始化embedding模型
        try:
            self.embedding_model = HuggingFaceEmbeddings(
                model_name="paraphrase-multilingual-MiniLM-L12-v2",
                model_kwargs={'device': 'cpu'},
                encode_kwargs={'normalize_embeddings': True}
            )
        except Exception as e:
            print(f"加载在线模型失败，使用本地模型: {str(e)}")
            model = SentenceTransformer('paraphrase-multilingual-MiniLM-L12-v2')
            self.embedding_model = HuggingFaceEmbeddings(
                model=model,
                model_kwargs={'device': 'cpu'},
                encode_kwargs={'normalize_embeddings': True}
            )
        
        # 初始化LLM，使用回调处理器
        self.llm = ChatOpenAI(
            openai_api_base=OPENAI_API_BASE,
            openai_api_key=OPENAI_API_KEY,
            model_name=MODEL_ID,
            temperature=0.7,
            streaming=True,
            callbacks=[StreamingStdOutCallbackHandler()]
        )
        
        self.vectorstore = None
        
        # 修改基础对话提示模板
        self.chat_prompt = ChatPromptTemplate.from_messages([
            ("system", """你是一个专业的图书推荐助手。请注意：
            1. 如果用户询问你能做什么，请简洁地介绍你是一个图书推荐助手，可以根据用户兴趣推荐合适的书籍
            2. 如果用户描述了想看的书籍类型或主题，你会推荐合适的书籍
            3. 在对话中保持友好专业，直接回答用户问题，不要解释你的思考过程"""),
            ("human", "{input}")
        ])
        
        # 修改推荐提示模板
        self.recommend_prompt = ChatPromptTemplate.from_messages([
            ("system", """你是一个专业的图书推荐助手。请严格遵守以下规则：
            1. 只能推荐检索到的书籍信息中实际包含的书籍
            2. 必须严格匹配用户的需求类型，不相关的书籍不要推荐
            3. 如果检索到的书籍与用户需求不匹配，必须给出补充推荐
            4. 推荐理由要包含简要的故事梗概，300字以内
            5. 如果没有匹配的馆藏书籍，直接从补充推荐开始
            6. 严格按照给定格式输出，保留标题但不要添加序号
            7. 如果既无馆藏推荐又无补充推荐，只返回"抱歉，目前没有找到符合您需求的书籍。请尝试描述其他类型的书籍。"
            """),
            ("human", """
            基于以下书籍信息：
            {context}
            
            用户需求：{question}
            
            严格按照以下格式回复：

            馆藏推荐：
            - 书名：xxx
            - 作者：xxx
            - 推荐理由：（包含300字以内的故事梗概）

            补充推荐：
            - 书名：xxx
            - 作者：xxx
            - 推荐理由：（包含300字以内的故事梗概）
            注：此书暂未收录在馆藏中

            建议阅读顺序：
            （简要说明）
            """)
        ])

    def initialize_vectorstore(self, documents):
        """初始化向量数据库"""
        try:
            self.vectorstore = Chroma(
                collection_name="book_collection",
                embedding_function=self.embedding_model,
                persist_directory=CHROMA_PERSIST_DIRECTORY
            )
            
            self.vectorstore.add_documents(documents)
            print(f"成功添加 {len(documents)} 个文档到向量数据库")
            
            # 创建检索链
            retriever = self.vectorstore.as_retriever(
                search_kwargs={"k": 3}
            )
            
            # 构建推荐链
            self.chat_chain = (
                {
                    "context": retriever, 
                    "question": RunnablePassthrough()
                }
                | self.recommend_prompt
                | self.llm
            ).with_config(
                {"configurable": {"temperature": "float"}}
            )
            
            # 构建普通对话链
            self.general_chat_chain = (
                self.chat_prompt
                | self.llm
            ).with_config(
                {"configurable": {"temperature": "float"}}
            )
            
        except Exception as e:
            raise Exception(f"初始化向量数据库失败: {str(e)}")

    def is_book_recommendation_query(self, query: str) -> bool:
        """判断是否是图书推荐查询"""
        # 定义意图识别提示
        intent_prompt = ChatPromptTemplate.from_messages([
            ("system", """判断用户输入是否在寻求图书推荐。
            如果用户在描述想看的书籍类型、主题、或请求推荐书籍，返回"True"
            如果是其他问题或对话，返回"False"
            只返回True或False，不要其他内容"""),
            ("human", query)
        ])
        
        # 关闭意图识别过程的流式输出
        intent_llm = ChatOpenAI(
            openai_api_base=OPENAI_API_BASE,
            openai_api_key=OPENAI_API_KEY,
            model_name=MODEL_ID,
            temperature=0.7,
            streaming=False  # 关闭流式输出
        )
        
        response = intent_prompt | intent_llm | StrOutputParser()
        return response.invoke({}).strip().lower() == "true"

    def get_recommendation(self, user_query: str) -> Optional[str]:
        """处理用户查询"""
        try:
            if not self.vectorstore:
                return "抱歉，图书数据库未初始化，无法提供推荐"
            
            # 获取相关文档
            retriever = self.vectorstore.as_retriever(
                search_kwargs={"k": 3}
            )
            
            # 构建推荐链
            recommend_chain = (
                {"context": retriever, "question": RunnablePassthrough()}
                | self.recommend_prompt
                | self.llm
            )
            
            # 执行推荐链
            return recommend_chain.invoke(user_query)
            
        except Exception as e:
            print(f"处理推荐请求时发生错误: {str(e)}")
            return f"抱歉，处理您的请求时发生错误: {str(e)}" 