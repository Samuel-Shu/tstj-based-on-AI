from openai import OpenAI

model_id = 'Qwen/Qwen2-7B-Instruct-GGUF'

client = OpenAI(
    base_url='https://ms-fc-7963d553-445c.api-inference.modelscope.cn/v1',
    api_key='077335bc-7b0e-42c3-8162-3d8dbce56f61'
)

response=client.chat.completions.create(
    model=model_id,
    messages=[{"role":"user", "content":"你好，能帮我介绍一下杭州吗？"}],
    stream=True
)

for chunk in response:
    print(chunk.choices[0].delta.content, end='', flush=True)