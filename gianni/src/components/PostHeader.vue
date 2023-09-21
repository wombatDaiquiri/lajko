<script setup>
defineProps({
  post: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="header">
    <img :src="post.avatar" alt="post.avatar" class="avatar"/>
    <div class="author">{{ post.author }}</div>
    <div class="likes">{{ post.likes }}</div>
  </div>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: start;
  align-items: center;
  margin-bottom: 1rem;
  font-weight: bold;
}

.avatar {
  max-height: 34px;
  max-width: 34px;
  border-radius: 100%;
}

.author {
  margin-left: 1rem;
  height: 100%;
}

.likes {
  margin-left: auto;
  margin-right: 1rem;
}
</style>