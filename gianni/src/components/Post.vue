<script setup>
import PostHeader from './PostHeader.vue'
defineProps({
  post: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <div class="post">
    <PostHeader :post=post />
    <div class="content">
      <div class="text" v-html="post.content"></div>
      <img v-for="imageURL in post.images" :src="imageURL" :alt="imageURL" :key="imageURL"/>
    </div>
    <div class="comments">
      <div v-for="comment in post.comments" class="comment">
        <PostHeader :post=comment />
        <div v-html="comment.content"></div>
        <div v-for="subcomment in comment.subcomments" class="subcomment">
          <PostHeader :post=subcomment />
          <div v-html="subcomment.content"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.post {
  margin: 2rem;
  padding: 1rem;
  background-color: rgb(27, 26, 26);
  color: rgb(158, 158, 158);
  max-width: 930px;
  border-radius: 10px;
}

.comment {
  background-color: rgb(33, 33, 33);
  margin: 1rem 0 1rem 0;
  padding: 1rem;
  border-radius: 10px;
}

.text {
  margin-bottom: 1rem;
}

img {
  max-width: 100%;
  max-height: 400px;
}

.subcomment {
  background-color: rgb(39, 39, 39);
  margin: 1rem 0 1rem 0;
  padding: 1rem;
  border-radius: 10px;
}
</style>