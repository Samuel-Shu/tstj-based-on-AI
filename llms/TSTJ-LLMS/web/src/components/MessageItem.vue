<template>
  <div class="message-item" :class="message.type">
    <div class="avatar">
      {{ message.type === 'user' ? '👤' : '🤖' }}
    </div>
    <div class="content">
      {{ message.content }}
    </div>
  </div>
</template>

<script setup>
defineProps({
  message: {
    type: Object,
    required: true
  }
})
</script>

<style scoped>
.message-item {
  display: flex;
  margin-bottom: 20px;
  gap: 10px;
}

.message-item.user {
  flex-direction: row-reverse;
}

.avatar {
  font-size: 24px;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.content {
  background: #f4f4f5;
  padding: 10px 15px;
  border-radius: 8px;
  max-width: 70%;
  white-space: pre-wrap;
}

.user .content {
  background: #ecf5ff;
}
</style> 