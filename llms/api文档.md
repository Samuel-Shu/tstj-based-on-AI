# 图书推荐系统 API 文档

## 基础信息
- 后端基础URL: `http://127.0.0.1:7001`
- 前端开发端口: `5173`

## API 接口说明

### 1. 健康检查
检查系统初始化状态

`GET /health`

**响应示例:**

```json
{
    "status": "ok",
    "system_initialized": true // true表示系统已初始化，false表示未初始化
}
```

### 2. 上传图书数据
上传并初始化图书数据库

`POST /admin/upload`

**请求参数:**
- Content-Type: `multipart/form-data`
- 参数:
  - `file`: Excel文件(.xlsx)，包含以下列：
    - title: 书名
    - author: 作者
    - category: 类别
    - description: 简介

**响应示例:**

```json
{
    "message": "文件上传成功，系统正在后台初始化",
    "status": "initializing"
}
```

### 3. 聊天接口
与AI助手进行对话

`POST /chat`

**请求参数:**
- Content-Type: `application/json`

- 请求体:

  ```json
  {
      "message": "string" // 用户输入的消息
  }
  ```

**响应示例:**

```json
{
    "response": "string" // AI助手的回复
}
```

## 前端组件说明

### 1. FileUpload 组件
文件上传组件，用于上传图书数据文件。

**事件:**
- `upload-success`: 文件上传成功时触发
  - 参数: `result` - 上传结果对象

### 2. ChatWindow 组件
聊天窗口组件，用于显示对话内容。

**属性:**
- `disabled`: Boolean - 是否禁用聊天功能
  - `true`: 系统未初始化
  - `false`: 系统已初始化

## 使用流程

1. **系统初始化:**
   - 启动后端服务: `python main.py`
   - 启动前端服务: `cd TSTJ-LLMS/web && npm run dev`

2. **上传图书数据:**
   - 准备Excel文件，包含必要的列（title, author, category, description）
   - 通过前端界面上传文件
   - 等待系统初始化完成

3. **开始对话:**
   - 系统初始化完成后，可以开始与AI助手对话
   - 可以询问系统能做什么
   - 可以请求图书推荐

## 注意事项

1. **超时设置:**
   - 前端请求超时时间: 240秒
   - 代理超时时间: 240秒

2. **文件格式要求:**
   - 仅支持.xlsx格式
   - Excel文件必须包含指定列名

3. **系统状态:**
   - 使用健康检查接口监控系统状态
   - 上传文件后系统会在后台异步初始化

4. **错误处理:**
   - 所有接口都包含错误处理
   - 错误信息会通过前端提示展示

## 配置说明

1. **后端配置 (config.py):**
   - LLM模型: Qwen2-7B-Instruct-GGUF
   - 向量数据库: Chroma
   - Embedding模型: paraphrase-multilingual-MiniLM-L12-v2

2. **前端配置 (vite.config.js):**
   - 开发端口: 5173
   - API代理: `/api` -> `http://127.0.0.1:7001`

## 项目结构

TSTJ-LLMS/

├── api.py # 后端API实现

├── config.py # 配置文件

├── main.py # 后端入口文件

├── rag_system.py # RAG系统实现

├── book_loader.py # 图书数据加载器

├── docs/ # 文档目录

│ └── api-documentation.md # API文档

├── data/ # 数据目录

│ ├── books.xlsx # 图书数据文件

│ └── chroma_db/ # 向量数据库文件

└── web/ # 前端项目目录

├── src/

│ ├── api/ # API调用

│ ├── components/# Vue组件

│ ├── App.vue # 主应用组件

│ └── main.js # 前端入口文件

├── package.json # 前端依赖配置

└── vite.config.js # Vite配置文件

