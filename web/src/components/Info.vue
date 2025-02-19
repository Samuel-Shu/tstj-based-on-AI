<template>
    <div class="user-page">
      <header class="header">
        <div class="title">图书AI推荐系统</div>
        <div class="user-info" @click="toggleDropdown">
          {{ username }} <el-button type="text">▼</el-button>
        </div>
        <el-button type="danger" @click="logout" class="logout-button" v-if="dropdownVisible">登出</el-button>
        <div v-if="dropdownVisible" class="dropdown-menu">
          <router-link to="/user-info" class="dropdown-item">用户信息管理</router-link>
          <el-button type="danger" @click="logout" class="dropdown-item">登出</el-button>
        </div>
      </header>
      <nav class="nav">
        <router-link to="/user" class="nav-item">首页</router-link>
        <router-link to="/info" class="nav-item">资讯信息</router-link>
        <router-link to="/recommend" class="nav-item">AI智能推荐</router-link>
      </nav>
      <el-card class="content-card">
        <h2>欢迎，用户！</h2>
        <p>这里是图书资讯。</p>
      </el-card>
    </div>
  </template>
  
  <script>
  export default {
  data() {
    return {
      dropdownVisible: false,
    };
  },
  computed: {
    username() {
      // 从 localStorage 获取用户名
      const name = localStorage.getItem('name'); // 假设用户信息存储在 'user' 中
      return name; // 获取用户名
    },
  },
  methods: {
    toggleDropdown() {
      this.dropdownVisible = !this.dropdownVisible; // 切换下拉菜单的显示状态
    },
    logout() {
      localStorage.removeItem('isAuthenticated');
      localStorage.removeItem('userRole');
      localStorage.removeItem('user'); // 清除用户信息
      this.$router.push('/'); // 登出后重定向到登录页面
    },
  },
};
  </script>
  
<style scoped>
.user-page {
  width: 100%; /* 扩展到整个屏幕宽度 */
  height: 100vh; /* 设置高度为100vh */
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

.user-info {
  cursor: pointer;
  display: flex;
  align-items: center;
}

.logout-button {
  margin-left: auto; /* 将登出按钮推到右侧 */
}

.nav {
  display: flex;
  justify-content: space-around;
  margin-bottom: 20px;
}

.nav-item {
  color: #409EFF; /* 主题颜色 */
  text-decoration: none;
  font-size: 16px;
}

.dropdown-menu {
  position: absolute;
  background: white;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  z-index: 1000;
}

.dropdown-item {
  display: block;
  padding: 10px;
  color: #409EFF;
  text-decoration: none;
}

.dropdown-item:hover {
  background: #f5f5f5;
}

</style>