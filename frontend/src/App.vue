<script setup>
import {onMounted, ref} from 'vue'
import { RouterLink, RouterView } from 'vue-router'
import TwitterLayout from './Layouts/TwitterLayout.vue';
import Tweet from './components/Tweet.vue';
import api from '@/config/api'
// import HelloWorld from './components/HelloWorld.vue'
defineProps({tweets: Array})


const tweets = ref([])



const fetchTweets = async()=> {
  try{
    const response = await api.get('/tweets')
    tweets.value = response.data
    // alert(JSON.stringify(response))
  }catch(error){
    console.error('Error fetching tweets', error)
  }
}

const handleTweetDeleted = () => {
  fetchTweets()  // Re-fetch tweets when a tweet is deleted
}

const handleTweetAdded = () => {
  fetchTweets()
}

onMounted(()=> {
  fetchTweets()
})
</script>

<template>
  <TwitterLayout @tweetAdded="handleTweetAdded">
    <div class="text-white">
      <div class="border-b border-b-gray-800 mt-2"></div>
     <div class="flex" v-for="tweet in tweets" :key="tweet">
      <!-- <Tweet :tweet="{
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
      }" /> -->
      <Tweet :tweet="tweet" @tweetDeleted="handleTweetDeleted"/>
     </div>
    </div>
  </TwitterLayout>

</template>
