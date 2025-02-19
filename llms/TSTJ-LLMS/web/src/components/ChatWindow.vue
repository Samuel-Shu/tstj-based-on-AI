<template>
  <div class="chat-container">
    <div class="messages" ref="messagesRef">
      <MessageItem
        v-for="(msg, index) in messages"
        :key="index"
        :message="msg"
      />
    </div>
    <div class="input-area">
      <el-input
        v-model="inputMessage"
        placeholder="请输入您的问题..."
        :disabled="disabled || loading"
        @keyup.enter="handleSendMessage"
      >
        <template #append>
          <el-button
            type="primary"
            :loading="loading"
            :disabled="disabled"
            @click="handleSendMessage"
          >
            发送
          </el-button>
        </template>
      </el-input>
      <div v-if="disabled" class="disabled-tip">
        请先上传图书数据
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { sendChatMessage } from '../api'
import MessageItem from './MessageItem.vue'

const props = defineProps({
  disabled: {
    type: Boolean,
    default: false
  }
})

const messages = ref([])
const inputMessage = ref('')
const loading = ref(false)
const messagesRef = ref(null)

const scrollToBottom = async () => {
  await nextTick()
  const container = messagesRef.value
  container.scrollTop = container.scrollHeight
}

const handleSendMessage = async () => {
  if (!inputMessage.value.trim()) {
    return
  }

  const userMessage = inputMessage.value
  messages.value.push({
    type: 'user',
    content: userMessage
  })
  
  inputMessage.value = ''
  loading.value = true
  
  try {
    const response = await sendChatMessage(userMessage)
    messages.value.push({
      type: 'assistant',
      content: response.response
    })
  } catch (error) {
    ElMessage.error('发送消息失败')
  } finally {
    loading.value = false
    scrollToBottom()
  }
}
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
}

.messages {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
}

.input-area {
  padding: 20px;
  border-top: 1px solid #dcdfe6;
}

.disabled-tip {
  text-align: center;
  color: #909399;
  margin-top: 10px;
  font-size: 14px;
}
</style> 