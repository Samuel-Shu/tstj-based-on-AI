<template>
  <el-container style="height: 100vh;">
    <el-container>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  // 移除 isAuthenticated 的计算属性
};
</script>

<style>
/* 移除与标题相关的样式 */
</style> 