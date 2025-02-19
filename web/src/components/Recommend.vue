<template>
  <div class="recommend-page">
    <h2>AI 聊天助手</h2>
    <div class="chat-container" ref="messages">
      <div class="messages">
        <div
          v-for="(msg, index) in chatMessages"
          :key="index"
          :class="['message', msg.sender === '用户' ? 'user-message' : 'ai-message']"
        >
          <strong>{{ msg.sender }}:</strong> {{ msg.text }}
        </div>
      </div>
    </div>
    <div class="input-container">
      <el-input
        v-model="userInput"
        placeholder="请输入您的消息"
        @keyup.enter="sendMessage"
        class="input-box"
      ></el-input>
      <el-button type="primary" @click="sendMessage">发送</el-button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      userInput: '',
      chatMessages: [],
    };
  },
  methods: {
    async sendMessage() {
      if (this.userInput.trim() === '') return;

      // 添加用户消息到聊天记录
      this.chatMessages.push({ sender: '用户', text: this.userInput });

      try {
        // 发送请求到聊天接口
        const response = await axios.post('http://127.0.1.1:7001/chat', {
          message: this.userInput,
        });

        // 添加 AI 的回复到聊天记录
        this.chatMessages.push({ sender: 'AI', text: response.data.reply });
      } catch (error) {
        console.error('聊天请求失败:', error);
        this.chatMessages.push({ sender: 'AI', text: '抱歉，我无法处理您的请求。' });
      }

      // 清空输入框
      this.userInput = '';

      // 滚动到最新消息
      this.$nextTick(() => {
        const messagesContainer = this.$refs.messages;
        messagesContainer.scrollTop = messagesContainer.scrollHeight;
      });
    },
  },
};
</script>

<style scoped>
.recommend-page {
  padding: 20px;
  background-color: #f9f9f9; /* 背景颜色 */
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.chat-container {
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 10px;
  height: 400px;
  overflow-y: auto;
  background-color: #fff; /* 聊天框背景颜色 */
}

.messages {
  margin-bottom: 10px;
}

.message {
  margin: 5px 0;
  padding: 10px;
  border-radius: 15px;
  max-width: 70%;
  word-wrap: break-word;
}

.user-message {
  background-color: #d1e7dd; /* 用户消息气泡颜色 */
  align-self: flex-end; /* 用户消息右对齐 */
}

.ai-message {
  background-color: #f8d7da; /* AI消息气泡颜色 */
  align-self: flex-start; /* AI消息左对齐 */
}

.input-container {
  display: flex;
  margin-top: 10px;
}

.input-box {
  flex: 1; /* 输入框占据剩余空间 */
  margin-right: 10px; /* 输入框与按钮之间的间距 */
}
</style>