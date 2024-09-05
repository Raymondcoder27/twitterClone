<script setup>
import { ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import Twitter from 'vue-material-design-icons/Twitter.vue'
import Magnify from 'vue-material-design-icons/Magnify.vue'
import DotsHorizontal from 'vue-material-design-icons/DotsHorizontal.vue'
import Feather from 'vue-material-design-icons/Feather.vue'
import Close from 'vue-material-design-icons/Close.vue'
import ChevronDown from 'vue-material-design-icons/ChevronDown.vue'
import Earth from 'vue-material-design-icons/Earth.vue'
import ImageOutline from 'vue-material-design-icons/ImageOutline.vue'
import FileGiftBox from 'vue-material-design-icons/FileGifBox.vue'
import Emoticon from 'vue-material-design-icons/Emoticon.vue'
import ArrowLeft from 'vue-material-design-icons/ArrowLeft.vue'
import MenuItem from '@/components/MenuItem.vue'
</script>

<template>
    <div class="fixed w-full ">
        <div class="flex mx-auto max-w-[1400px]">
            <div class="lg:w-3/12 w-[60px] h-[100vh] lg:px-4 max-w-[350px] lg:mx-auto">
                <div class="p-2 px-3 mb-4">
                    <img src="../assets/xlogo-removebg-preview.png" class="w-[50px]" alt="">
                </div>

                <MenuItem iconString="home" />
            </div>
        </div>
    </div>
</template>

<style>
    body{
        background-color: black;
    }
</style>