<script>
export default {
  // Properties returned from data() become reactive state and will be exposed on this.
  // data() {
  //   return {
  //     post: {}
  //   }
  // },
  props: ['post'],
  // Methods are functions that mutate state and trigger updates.
  // They can be bound as event listeners in templates.
  methods: {
    // increment() {
    //   this.count++
    // }
  },
  // Lifecycle hooks are called at different stages of a component's lifecycle.
  // This will be called when the component is mounted.
  mounted() {
    console.log(`mounted header with post`, this.post.value)
  }
}
</script>

<template>
  <div class="header">
    <img :src="post.author" alt="post.avatar" class="avatar"/>
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