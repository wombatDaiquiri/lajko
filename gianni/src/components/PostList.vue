<template>
  <div>
    <div v-for="post in posts" :key="post.id">
      <h2>{{ post.author.username }}</h2>
      <MarkdownRenderer :markdownContent="post.content" />
    </div>
  </div>
</template>

<script>
import { GraphQLClient } from 'graphql-request';
import { GET_POSTS } from '../graphql/queries';
import MarkdownRenderer from './MarkdownRenderer.vue';

export default {
  data() {
    return {
      posts: [],
    };
  },
  async mounted() {
    const client = new GraphQLClient('lajko/graphql');
    try {
      const data = await client.request(GET_POSTS);
      this.posts = data.posts;
    } catch (error) {
      console.error(error);
    }
  },
  components: {
    MarkdownRenderer
  },
};
</script>
