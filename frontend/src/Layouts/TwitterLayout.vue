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
                        <div class="flex">
                            <div class="flex items-center justify-center hover:bg-gray-500 p-4 text-white font-extrabold text-[17px] w-full h-[60px]">
                                <div class="inline-block text-center border-b-4 h-[68px] border-b-[#1C9CEF]">
                                    <div class="my-auto mt-4">For you</div>
                                </div>
                            </div>

                            <div class=" hover:bg-gray-500 p-4 text-gray-500 font-extrabold text-[17px] w-full h-[60px]">
                                <div class="text-center">Following</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="absolute top-0 z-0 h-full overflow-auto scrollbar-hide">
                <div class="mt-[126px]"></div>
                <slot />
                <div class="pb-4"></div>
            </div>

            <!-- <div class="lg:block hidden lg:w-4"> -->
            <div class="lg:block hidden lg:w-4/12 h-screen border-l border-gray-800 pl-4">
                <div class="w-full p-1 mt-2 px-4 lg:flex items-center rounded-full hidden bg-[#212327]">
                    <Magnify fillColor="#5E5C5C" :size="25" />
                    <input type="text" placeholder="Search Twitter" 
                    class="appearance-none w-full border-0 py-2 text-gray-100 bg-[#212327] placeholder-gray-500 leading-tight focus:ring-0
                    ">
                </div>

                <div class="w-full mt-4 rounded-lg lg:block hidden bg-[#212327]">
                    <div class="w-full p-4 text-white font-extrabold pb-6 text-[20px]">
                        What's happening
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