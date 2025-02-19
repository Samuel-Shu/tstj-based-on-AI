<template>
  <el-card class="register-card">
    <h2 class="register-title">注册</h2>
    <el-form :model="form" ref="formRef" label-width="80px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="form.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="form.email" placeholder="请输入邮箱"></el-input>
      </el-form-item>
      <el-form-item label="验证码" prop="verificationCode">
        <el-button type="primary" :disabled="isCounting" @click="getVerificationCode">
          {{ isCounting ? `${countdown}s` : '获取验证码' }}
        </el-button>
        <el-input v-model="form.verificationCode" placeholder="请输入验证码" style="margin-top: 10px;"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" :disabled="!form.verificationCode" @click="register">
          注册
        </el-button>
        <span v-if="!form.verificationCode" class="error-message">请先输入验证码</span>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>
import { registerUser } from '../services/api'; // 引入 API 服务
import axios from 'axios'; // 引入 axios

export default {
  data() {
    return {
      form: {
        username: '',
        password: '',
        email: '',
        verificationCode: '',
      },
      isCounting: false, // 控制按钮是否可用
      countdown: 60, // 倒计时初始值
      timer: null, // 定时器
    };
  },
  methods: {
    async getVerificationCode() {
      try {
        // 发送获取验证码的请求
        const response = await axios.post('http://localhost:8300/api/register/vlcode', { email: this.form.email });
        
        // 显示状态消息
        this.$message.success(response.data.status_msg);

        // 检查状态码
        if (response.data.status_code === 200) {
          // 启动倒计时
          this.isCounting = true;
          this.countdown = 60;
          this.timer = setInterval(() => {
            this.countdown--;
            if (this.countdown <= 0) {
              clearInterval(this.timer);
              this.isCounting = false;
            }
          }, 1000);
        }
      } catch (error) {
        // 处理错误响应
        if (error.response && error.response.data) {
          const { status_msg } = error.response.data;
          this.$message.error(status_msg || '获取验证码失败，请重试');
        } else {
          this.$message.error('获取验证码失败，请重试');
        }
      }
    },
    async register() {
      try {
        const response = await registerUser(this.form);
        
        // 检查状态码
        if (response.data.status_code === 200) {
          // 显示注册成功的消息
          this.$message.success(response.data.status_msg);
          this.$router.push('/'); // 注册成功后跳转到登录页面
        } else {
          this.$message.error(response.data.status_msg || '注册失败，请重试');
        }
      } catch (error) {
        // 处理错误响应
        if (error.response && error.response.data) {
          const { status_msg } = error.response.data;
          this.$message.error(status_msg || '注册失败，请重试');
        } else {
          this.$message.error('注册失败，请重试');
        }
      }
    },
  },
  beforeDestroy() {
    // 清理定时器
    if (this.timer) {
      clearInterval(this.timer);
    }
  },
};
</script>

<style scoped>
.register-card {
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.register-title {
  text-align: center;
  margin-bottom: 20px;
  color: #409EFF; /* 主题颜色 */
}

.el-form-item {
  margin-bottom: 15px; /* 固定间距 */
}

.error-message {
  color: red;
  font-size: 12px;
  margin-top: 5px;
}
</style> 