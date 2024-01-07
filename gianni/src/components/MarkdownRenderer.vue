<template>
  <div v-html="renderedMarkdown"></div>
</template>

<script>
import MarkdownIt from 'markdown-it';

export default {
  props: {
    markdownContent: {
      type: String,
      required: true,
    },
  },
  computed: {
    renderedMarkdown() {
      const md = new MarkdownIt();
      return md.render(this.markdownContent);
    },
  },
};
</script>
