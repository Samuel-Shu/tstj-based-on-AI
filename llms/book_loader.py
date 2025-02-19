from typing import List
import pandas as pd
import os
from langchain.schema import Document

class BookLoader:
    def __init__(self, data_path: str = "./data/books.xlsx"):
        self.data_path = data_path

    def load_books(self) -> List[Document]:
        """从Excel文件加载图书数据并转换为Document格式"""
        if not os.path.exists(self.data_path):
            raise FileNotFoundError(f"找不到图书数据文件: {self.data_path}")

        try:
            # 读取Excel文件
            df = pd.read_excel(
                self.data_path,
                # 指定列名
                names=['title', 'author', 'category', 'description'],
                # 如果第一行是标题，设置header=0；如果第一行就是数据，设置header=None
                header=0
            )

            documents = []
            # 遍历DataFrame的每一行
            for _, row in df.iterrows():
                # 将书籍信息组合成文本
                content = f"书名：{row['title']}\n作者：{row['author']}\n"
                content += f"类别：{row['category']}\n简介：{row['description']}\n"
                
                # 创建Document对象
                doc = Document(
                    page_content=content,
                    metadata={
                        "title": row['title'],
                        "author": row['author'],
                        "category": row['category']
                    }
                )
                documents.append(doc)

            return documents
            
        except Exception as e:
            raise Exception(f"读取Excel文件时发生错误: {str(e)}") 