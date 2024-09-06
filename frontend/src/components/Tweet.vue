<script setup>
    import {ref, onMounted} from 'vue'
    import HeartOutline from 'vue-material-design-icons/HeartOutline.vue'
    import ChartBar from 'vue-material-design-icons/ChartBar.vue'
    import MessageOutline from 'vue-material-design-icons/MessageOutline.vue'
    import Sync  from 'vue-material-design-icons/Sync.vue'
    import DotsHorizontal from 'vue-material-design-icons/DotsHorizontal.vue'
    import TrashCanOutline from 'vue-material-design-icons/TrashCanOutline.vue'

    const props = defineProps({
        tweet: Object
    })

    let openOptions = ref(false)
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
            </div>
            <div class="hover:bg-gray-800 rounded-full cursor-pointer relative">
                <button type="button" class="block p-2">
                    <DotsHorizontal @click="openOptions = !openOptions" />
                </button>
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
            <div
            v-else>
                <video :src="tweet.file" class="rounded-xl" controls></video>
            </div>
        </div>


        <div class="flex items-center justify-between mt-4 w-4/5">
            <div class="flex">
                <MessageOutline text-[#5e5c5c] :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.comments }}</span>
            </div>
            <div class="flex">
                <MessageOutline text-[#5e5c5c] :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.comments }}</span>
            </div>
            <div class="flex">
                <MessageOutline text-[#5e5c5c] :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.comments }}</span>
            </div>
            <div class="flex">
                <MessageOutline text-[#5e5c5c] :size="18" />
                <span class="text-xs text-[#5e5c5c] font-extrabold ml-3">{{ tweet.comments }}</span>
            </div>
        </div>
    </div>
</template>