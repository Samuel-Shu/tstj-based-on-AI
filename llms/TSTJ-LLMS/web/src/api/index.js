import axios from 'axios'
import { ElMessage } from 'element-plus'

const api = axios.create({
  baseURL: '/api',
  timeout: 240000,  // 增加到120秒
})

// 添加响应拦截器
api.interceptors.response.use(
  response => response,
  error => {
    if (error.code === 'ECONNABORTED' && error.message.includes('timeout')) {
      ElMessage.error('请求超时，请稍后重试')
    } else {
      const message = error.response?.data?.detail || '请求失败'
      ElMessage.error(message)
    }
    return Promise.reject(error)
  }
)

export const uploadBookFile = async (file) => {
  try {
    const formData = new FormData()
    formData.append('file', file)
    const response = await api.post('/admin/upload', formData)
    return response.data
  } catch (error) {
    throw error
  }
}

export const sendChatMessage = async (message) => {
  try {
    const response = await api.post('/chat', { message })
    return response.data
  } catch (error) {
    throw error
  }
}

// 添加健康检查
export const checkHealth = async () => {
  try {
    const response = await api.get('/health')
    return response.data
  } catch (error) {
    throw error
  }
} 