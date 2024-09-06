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
        </div>
    </div>
</template>