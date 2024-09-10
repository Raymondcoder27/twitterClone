<script setup>
    import {ref, onMounted} from 'vue'
    import HeartOutline from 'vue-material-design-icons/HeartOutline.vue'
    import ChartBar from 'vue-material-design-icons/ChartBar.vue'
    import MessageOutline from 'vue-material-design-icons/MessageOutline.vue'
    import Sync  from 'vue-material-design-icons/Sync.vue'
    import DotsHorizontal from 'vue-material-design-icons/DotsHorizontal.vue'
    import BookMarkOutline from 'vue-material-design-icons/BookmarkOutline.vue'
    import TrashCanOutline from 'vue-material-design-icons/TrashCanOutline.vue'
    import api from '@/config/api'


    const emit = defineEmits()

    const props = defineProps({
        tweet: Object
    })

    let openOptions = ref(false)
    const tweets = ref([])
  

    const fetchTweets = async()=> {
  try{
    const response = await api.get('/tweets')
    tweets.value = response.data
  }catch(error){
    console.error('Error fetching tweets', error)
  }
}


    const deleteTweet = async(id)  => {
        try{
        await api.delete('/tweet/'+id)
        emit('tweetDeleted')  // Emit event when tweet is deleted
        // window.location.reload()

        // closeMessageBox()
        // window.location.reload()
        console.log(`Tweet with ID ${id} deleted`);
         fetchTweets()
    // alert(JSON.stringify(tweets))

    }catch(error){
        console.error('error deleting tweet:', error)
    }
    }

    onMounted(()=>{
        fetchTweets()
    })
</script>

<template>
    <div class="min-w-[60px]">
        <img :src="tweet.image" width="50" alt="" class="rounded-full m-2 mt-3">
    </div>
    <div class="p-2 w-full">
        <div class="font-extrabold flex items-center justify-between mt-0.5 mb-1.5">
            <div class="flex items-center">
                <div>{{tweet.name}}</div>
                <span class="font-[300] text-[15px] pl-2 text-gray-500">{{ tweet.handle }}</span>
                <!-- <span class="font-[300] text-[15px] pl-2 text-gray-500">{{ tweet.created_at }}h</span> -->
            </div>
            <div class="hover:bg-gray-800 rounded-full cursor-pointer relative">
                <button @click="openOptions = !openOptions"  type="button" class="block p-2">
                    <DotsHorizontal />
                </button>

                <div
                v-if="openOptions" class="absolute mt-1 p-3 right-0 bg-black border border-gray-700 rounded-lg shadow-lg w-[300px]">
                    <div
                    class="flex"
                    as="button"
                    @click="deleteTweet(tweet.ID)">
                    <TrashCanOutline class="pr-3 " fillColor="#DC2626" :size="18"/>
                    <span class="text-red-600 font-extrabold">Delete</span>
                </div>
                </div>
            </div>
        </div>

        <div class="pb-3 text-sm">
            {{tweet.tweet}}
        </div>

        <div
        v-if="tweet.file">
            <div
            v-if="!tweet.is_video" class="rounded-xl">
                <img :src="tweet.file" alt="" class="mt-2 object-fill rounded-xl-w-full">
            </div>
            <div v-else>
                <video :src="tweet.file" class="rounded-xl" controls></video>
            </div>
        </div>


        <div class="flex items-center justify-between mt-4 w-4/5">
            <div class="flex">
                <MessageOutline fillColor="#5e5c5c" :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.comments }}</span>
            </div>
            <div class="flex">
                <Sync fillColor="#5e5c5c" :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.retweets }}</span>
            </div>
            <div class="flex">
                <HeartOutline fillColor="#5e5c5c" :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.likes }}</span>
            </div>
            <div class="flex">
                <ChartBar fillColor="#5e5c5c" :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.analytics }}</span>
            </div>
            <div class="flex">
                <BookMarkOutline fillColor="#5e5c5c" :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3"></span>
            </div>
        </div>
    </div>
</template>