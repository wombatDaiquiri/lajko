<template>
  <div>
    <button
        v-for="category in categories"
        :key="category"
        :class="{ active: activeCategory === category }"
        @click="fetchPosts(category)"
    >{{ category }}</button>


    <div v-for="post in posts" :key="post.id">
      <h2>{{ post.author.username }}</h2>
      <MarkdownRenderer :markdownContent="post.content" />
    </div>
  </div>
</template>

  <script>
import { GraphQLClient } from 'graphql-request';
import { GET_POSTS_BY_CATEGORY } from '../graphql/queries';
import MarkdownRenderer from './MarkdownRenderer.vue';

export default {
  data() {
    return {
      posts: [],
      activeCategory: 'p.createdAt',
      categories: ["p.createdAt", "p.numLikes", "p.numComments", "p.hot", "p.hotness", "rand", "p.promoted"]
    };
  },

  methods: {
    async fetchPosts(category) {
      this.activeCategory = category;
      const client = new GraphQLClient('/lajko/graphql');
      const variables = {
        category: category
      };
      try {
        const data = await client.request(GET_POSTS_BY_CATEGORY, variables);
        this.posts = data.posts;
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    }
  },

  mounted() {
    this.fetchPosts('p.createdAt');
  },
  components: {
    MarkdownRenderer
  },
};
</script>

<style>
.active {
  /* Style for active category button */
  background-color: blue;
  color: white;
}
</style>
