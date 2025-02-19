<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <h1>图书推荐系统</h1>
        <div class="status-container">
          <el-tag :type="systemInitialized ? 'success' : 'warning'">
            {{ systemInitialized ? '系统已初始化' : '请先上传图书数据' }}
          </el-tag>
          <FileUpload @upload-success="handleUploadSuccess" />
        </div>
      </el-header>
      <el-main>
        <ChatWindow :disabled="!systemInitialized" />
      </el-main>
    </el-container>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import FileUpload from './components/FileUpload.vue'
import ChatWindow from './components/ChatWindow.vue'
import { checkHealth } from './api'

const systemInitialized = ref(false)

const checkSystemStatus = async () => {
  try {
    const { system_initialized } = await checkHealth()
    systemInitialized.value = system_initialized
  } catch (error) {
    console.error('检查系统状态失败:', error)
  }
}

const handleUploadSuccess = async (result) => {
  ElMessage.success(`成功上传图书数据：${result.book_count} 本`)
  await checkSystemStatus()
}

onMounted(checkSystemStatus)
</script>

<style>
.app-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.el-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
}

.status-container {
  display: flex;
  align-items: center;
  gap: 20px;
}

.el-main {
  padding: 20px;
  height: calc(100vh - 200px);
}
</style> 