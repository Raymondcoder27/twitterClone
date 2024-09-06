<script setup>
import { RouterLink, RouterView } from 'vue-router'
import TwitterLayout from './Layouts/TwitterLayout.vue';
import Tweet from './components/Tweet.vue';
// import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <TwitterLayout>
    <div class="text-white">
      <div class="border-b border-b-gray-800 mt-2"></div>
     <div class="flex">
      <Tweet :tweet="{
        name: 'Ronald Mpagi',
        handle: '@Mpagi',
        image: 'https://randomuser.me/api/portraits/men/40.jpg',
        tweet: 'We went rock climbing this weekend. Here is the video, climbing is way more fun than cycling.',
        file: '/videos/Sportsman.mp4',
        is_video: true,
        comments: '34',
        retweets: '54',
        likes: '87',
        analytics: '81',
      }" />
     </div>
    </div>
  </TwitterLayout>

</template>
