import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8300/api', // 后端 API 基础 URL
  headers: {
    'Content-Type': 'application/json',
  },
});

// 用户注册
export const registerUser = (userData) => {
  return apiClient.post('/register', userData);
};

// 用户登录
export const loginUser = (userData) => {
  return apiClient.post('/login', userData);
}; 