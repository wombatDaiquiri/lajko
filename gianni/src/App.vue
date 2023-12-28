<script>
import Post from './components/Post.vue'
import gql from 'graphql-tag';

export default {
  data() {
    return {
      posts: null,
      loading: false,
      error: null
    };
  },

  apollo: {
    // Define the Apollo query
    posts: {
      query: gql`query posts {
          posts {
            author {
              username
              avatarUrl
            }
            title
            content
          }
        }`,
      // Update loading and error states
      update(data) {
        console.log('data came in: ', data.posts)
        this.loading = false;
        return data.posts;
      },
      loadingKey: 'loading',
      error(error) {
        this.error = error;
        this.loading = false;
      }
    }
  },

  created() {
    console.log('created')
    this.loading = true;
  },

  components: { Post },
};
</script>

<template>
  <main>
    <h1>rendering posts below</h1>
    <Post :post=post v-for="post in posts"/>
  </main>
</template>

<style>
@font-face {
  font-family: Poppins-medium;
  src: url('./assets/fonts/Poppins-Medium.ttf');
}
@font-face {
  font-family: Poppins;
  src: url('./assets/fonts/Poppins-Regular.ttf');
}
</style>
