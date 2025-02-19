from dotenv import load_dotenv
import os

# 加载环境变量
load_dotenv()

# OpenAI API配置
OPENAI_API_BASE = ""
OPENAI_API_KEY = ""
MODEL_ID = ""

# 向量数据库配置
CHROMA_PERSIST_DIRECTORY = "./data/chroma_db"

# embedding模型配置
EMBEDDING_MODEL_NAME = "paraphrase-multilingual-MiniLM-L12-v2"  # 使用默认多语言模型
