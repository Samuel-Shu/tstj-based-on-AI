<template>
  <div class="admin-page">
    <header class="header">
      <div class="title">图书AI推荐系统</div>
      <el-button type="danger" @click="logout">登出</el-button>
    </header>
    <el-card class="content-card">
      <h2>管理员页面</h2>
      <p>这里是管理员的控制中心。</p>
    </el-card>
  </div>
</template>

<script>
export default {
  methods: {
    logout() {
      localStorage.removeItem('isAuthenticated');
      localStorage.removeItem('userRole');
      this.$router.push('/'); // 登出后重定向到登录页面
    },
  },
};
</script>

<style scoped>
.admin-page {
  width: 100%; /* 扩展到整个屏幕宽度 */
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.title {
  font-size: 24px;
  color: #409EFF; /* 主题颜色 */
}

.content-card {
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  width: 100%; /* 扩展到整个屏幕宽度 */
}
</style> 