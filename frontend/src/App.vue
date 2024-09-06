<script setup>
import { RouterLink, RouterView } from 'vue-router'
import TwitterLayout from './Layouts/TwitterLayout.vue';
// import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <TwitterLayout>
    <div class="text-white">
      <div class="border-b border-b-gray-800 mt-2"></div>
      <Tweet />
    </div>
  </TwitterLayout>

</template>
