<template>
  <div class="upload-container">
    <el-upload
      class="upload"
      action="#"
      :auto-upload="false"
      :on-change="handleFileChange"
      :on-remove="handleFileRemove"
      :limit="1"
      accept=".xlsx"
      :file-list="fileList"
      :disabled="uploading"
    >
      <template #trigger>
        <el-button type="primary" :disabled="uploading">选择文件</el-button>
      </template>
      <el-button
        type="success"
        :disabled="!currentFile || uploading"
        :loading="uploading"
        @click="handleUpload"
      >
        {{ uploading ? '初始化中...' : '上传数据' }}
      </el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { uploadBookFile } from '../api'

const emit = defineEmits(['upload-success'])
const currentFile = ref(null)
const fileList = ref([])
const uploading = ref(false)

const handleFileChange = (uploadFile) => {
  currentFile.value = uploadFile.raw
  fileList.value = [uploadFile]
}

const handleFileRemove = () => {
  currentFile.value = null
  fileList.value = []
}

const handleUpload = async () => {
  if (!currentFile.value) {
    ElMessage.warning('请先选择文件')
    return
  }

  uploading.value = true
  ElMessage.info('正在上传并初始化系统，请稍候...')
  
  try {
    const result = await uploadBookFile(currentFile.value)
    emit('upload-success', result)
    // 清空文件选择
    currentFile.value = null
    fileList.value = []
    ElMessage.success('上传成功')
  } catch (error) {
    console.error('上传失败:', error)
    ElMessage.error('上传失败，请重试')
  } finally {
    uploading.value = false
  }
}
</script>

<style scoped>
.upload-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

:deep(.el-upload-list) {
  margin: 0 10px;
}

:deep(.el-upload-list__item) {
  transition: none;
}
</style> 