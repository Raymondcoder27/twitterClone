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

                <MenuItem iconString="Home" />
                <MenuItem iconString="Explore" />
                <MenuItem iconString="Notifications" />
                <MenuItem iconString="Messages" />
                <!-- <MenuItem iconString="Grok" /> -->
                <MenuItem iconString="Profile" />
                <MenuItem iconString="More" />


                <button class="lg:w-full mt-8 ml-2 text-white font-extrabold p-3 px-3 rounded-full cursor-pointer bg-[#1C9CEF]">
                    <span class="lg:block hidden">
                        Post
                    </span>
                    <span class="block lg:hidden">
                        <Feather />
                    </span>
                </button>
            </div>


            <div class="border-x border-gray-800 relative lg:w-7/12 w-11/12">
                <div class="bg-black bg-opacity-50 backdrop-blur-md z-10 absolute w-full">
                    <div class="border-gray-800 border-b w-full">
                        <div class="w-full text-white font-extrabold p-4 text-[22px]">
                            Home
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
    body{
        background-color: black;
    }
</style>