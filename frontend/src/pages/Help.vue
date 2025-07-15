<script setup lang="ts">
import { computed } from "vue";
import markdownIt from "markdown-it";
import FloatingNavBar from "../components/layouts/FloatingNavBar.vue";

const md = markdownIt({
  html: true,
  linkify: true,
  typographer: true,
});

const markdown = `
# Ro-Wails Help

Welcome to the **Ro-Wails** help page! Here you'll find useful information to get started.

## Features

- ⚡ Fast & Lightweight
- 🛠️ Built with Go and Vue
- 🎨 Customizable UI with Tailwind and Shadcn

## Getting Started

1. Click **Get Started** on the home page.
2. Explore the documentation.
3. Join our community on [GitHub](https://github.com/your-org/ro-wails).

\`\`\`
const codeBlocks: string = "Works great"
\`\`\`

## FAQ

**Q:** How do I switch themes?  
**A:** Use the theme switcher in the right controls panel.

**Q:** Where can I report issues?  
**A:** Please open an issue on our [GitHub repository](https://github.com/your-org/ro-wails/issues).

---

Happy coding! 🚀
`;

const rendered = computed(() => md.render(markdown));
</script>

<template>
  <FloatingNavBar />
  <div class="flex justify-center">
    <div
      class="prose dark:prose-dark max-w-2xl w-full p-6 sm:p-12 dark:max-w-2xl dark:w-full dark:p-6 dark:sm:p-12"
    >
      <div v-html="rendered"></div>
    </div>
  </div>
</template>
