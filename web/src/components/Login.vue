<template>
   <header class="header">
      <div class="title">图书AI推荐系统</div>
    </header>
  <el-card class="login-card">
    <h2 class="login-title">登录</h2>
    <el-form :model="form" ref="formRef" label-width="80px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="form.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="login">登录</el-button>
        <router-link to="/register">
          <el-button>注册账户</el-button>
        </router-link>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>
import { loginUser } from '../services/api'; // 引入 API 服务

export default {
  data() {
    return {
      form: {
        username: '',
        password: '',
      },
    };
  },
  methods: {
    async login() {
      try {
        const response = await loginUser(this.form);
        
        // 检查状态码
        if (response.data.status_code === 200) {
          localStorage.setItem('isAuthenticated', true);
          localStorage.setItem('userRole', response.data.role);
          localStorage.setItem('name', response.data.name);

          // 显示状态消息
          this.$message.success(response.data.status_msg);

          if (response.data.role === 0) {
            this.$router.push('/admin');
          } else {
            this.$router.push('/user');
          }
        } else {
          this.$message.error(response.data.status_msg || '登录失败，请重试');
        }
      } catch (error) {
        this.$message.error('登录失败，请重试');
      }
    },
  },
};
</script>

<style scoped>
.title {
  font-size: 24px;
  color: #409EFF; /* 主题颜色 */
}

.login-card {
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.login-title {
  text-align: center;
  margin-bottom: 20px;
  color: #409EFF; /* 主题颜色 */
}

.el-form-item {
  margin-bottom: 15px; /* 固定间距 */
}
</style> 
