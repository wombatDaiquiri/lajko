<script setup>
import posts from '../../data-snapshots/2023-09-20T21:46:31.638Z.json';
import Post from './components/Post.vue'
// posts.forEach(x => { console.log(x, x.author, x.content); });

</script>

<template>
  <main>
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
