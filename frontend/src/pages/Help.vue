<script setup lang="ts">
import { computed } from 'vue'
import markdownIt from 'markdown-it'
import FloatingNavBar from '../components/layouts/FloatingNavBar.vue';

const md = markdownIt({
  html: true,
  linkify: true,
  typographer: true,
});

const markdown = `
# Aide

Coming soon !

`;

const rendered = computed(() => md.render(markdown));
</script>

<template>
  <FloatingNavBar />
  <div class="flex justify-center">
    <div class="prose dark:prose-dark max-w-2xl w-full p-6 sm:p-12 dark:max-w-2xl dark:w-full dark:p-6 dark:sm:p-12">
      <div v-html="rendered"></div>
    </div>
  </div>
</template>
