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
import FileGifBox from 'vue-material-design-icons/FileGifBox.vue'
import Emoticon from 'vue-material-design-icons/Emoticon.vue'
import ArrowLeft from 'vue-material-design-icons/ArrowLeft.vue'
import MenuItem from '@/components/MenuItem.vue'

let randImg1 = ref(`https://picsum.photos/id/${(Math.random() * 200).toFixed(0)}/100`)
let randImg2 = ref(`https://picsum.photos/id/${(Math.random() * 200).toFixed(0)}/100`)

let textarea = ref(null)
let createTweet = ref(false)
let tweet = ref('')

let file = ref('')
let showUpload = ref('')
let uploadType = ref('')

const getFile = () => {
    file.value = e.target.files[0]
    showUpload.value = URL.createObjectURL(e.target.files[0])
    uploadType.value = file.value.name.split('.').pop();
}

const closeMessageBox = () => {
    createTweet.value = false
    tweet.value = ''
    showUpload.value = ''
    uploadType.value = ''
}

const textareaInput = (e) => {
    textarea.value.style.height = "auto"
    textarea.value.style.height = `${e.target.scrollHeight}px`;
}

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
                            <div class="flex items-center justify-center hover:bg-opacity-30 cursor-pointer hover:bg-gray-500 p-4 text-white font-extrabold text-[17px] w-full h-[60px]">
                                <div class="inline-block text-center border-b-4 h-[60px] rounded-sm border-b-[#1C9CEF]">
                                    <div class="my-auto mt-4">For you</div>
                                </div>
                            </div>

                            <div class=" hover:bg-gray-500 hover:bg-opacity-30 cursor-pointer p-4 text-gray-500 font-extrabold text-[17px] w-full h-[60px]">
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
            <div class="lg:block hidden lg:w-4/12 h-screen border-l border-gray-800 pl-4 pr-8">
                <div class="w-full p-1 mt-2 px-4 lg:flex items-center rounded-full hidden bg-[#212327]">
                    <Magnify fillColor="#5E5C5C" :size="25" />
                    <input type="text" placeholder="Search" 
                    class="appearance-none w-full border-0 py-2 text-gray-100 bg-[#212327] placeholder-gray-500 leading-tight focus:outline-none
                    ">
                </div>

                <div class="w-full mt-4 rounded-lg lg:block hidden border border-gray-800">
                    <div class="w-full p-4 text-white font-extrabold pb-6 text-[20px]">
                        Trends for you
                    </div>

                    <div class="hover:bg-neutral-950 cursor-pointer transition duration-200 ease-in-out h-[80px]">
                        <div class="flex justify-between py-3 p-3 h-[80px]">
                            <div>
                                <div class="text-gray-400 text-[13px]">Trending in Uganda</div>
                                <div class="w-full text-white font-extrabold mb-6 text-[15px]">
                                    Tertiary Institution curriculum overhaul
                                </div>
                            </div>
                            <img class="rounded-xl" :src="randImg2" alt="">
                        </div>
                    </div>
                    <div class="hover:bg-neutral-950 cursor-pointer transition duration-200 ease-in-out">
                        <div class="flex justify-between p-3 ">
                            <div>
                                <div class="text-gray-400 text-[13px]">Trending in Uganda</div>
                                <div class="w-full text-white font-extrabold text-[15px]">
                                    Christiano Ronaldo
                                </div>
                                <div class="text-gray-400 text-[14px]">3,000 posts</div>
                            </div>
                            <DotsHorizontal fillColor="#FFFFFF" />
                        </div>
                    </div>
                    <div class="hover:bg-neutral-950 cursor-pointer transition duration-200 ease-in-out">
                        <div class="flex justify-between p-3 ">
                            <div>
                                <div class="text-gray-400 text-[13px]">Trending in Uganda</div>
                                <div class="w-full text-white font-extrabold text-[15px]">
                                    Makerere
                                </div>
                                <div class="text-gray-400 text-[14px]">7k posts</div>
                            </div>
                            <DotsHorizontal fillColor="#FFFFFF" />
                        </div>
                    </div>
                    <div class="hover:bg-neutral-950 cursor-pointer transition duration-200 ease-in-out">
                        <div class="flex justify-between p-3 ">
                            <div>
                                <div class="text-gray-400 text-[13px]">Manchester United . Trending</div>
                                <div class="w-full text-white font-extrabold text-[15px]">
                                    Ten Hag
                                </div>
                                <div class="text-gray-400 text-[14px]">23k posts</div>
                            </div>
                            <DotsHorizontal fillColor="#FFFFFF" />
                        </div>
                    </div>
                    <div class="hover:bg-neutral-950 cursor-pointer transition duration-200 ease-in-out">
                        <div class="flex justify-between p-3 ">
                            <div>
                                <div class="text-gray-400 text-[13px]">Trending in East Africa</div>
                                <div class="w-full text-white font-extrabold text-[15px]">
                                    THE GOAT
                                </div>
                                <div class="text-gray-400 text-[14px]">238k posts</div>
                            </div>
                            <DotsHorizontal fillColor="#FFFFFF" />
                        </div>
                    </div>  
                    <!-- <div class="hover:bg-neutral-950 cursor-pointer transition duration-200 ease-in-out">
                        <div class="flex justify-between p-3 ">
                            <div>
                                <div class="text-gray-400 text-[13px]">Trending in Uganda</div>
                                <div class="w-full text-white font-extrabold text-[15px]">
                                    Namboole
                                </div>
                                <div class="text-gray-400 text-[14px]">2k posts</div>
                            </div>
                            <DotsHorizontal fillColor="#FFFFFF" />
                        </div>
                    </div>   -->
                    <!-- <div class="hover:bg-neutral-950 cursor-pointer transition duration-200 ease-in-out">
                        <div class="flex justify-between p-3 ">
                            <div>
                                <div class="text-gray-400 text-[13px]">Trending in Kampala</div>
                                <div class="w-full text-white font-extrabold text-[15px]">
                                    Gashumba
                                </div>
                                <div class="text-gray-400 text-[14px]">13k posts</div>
                            </div>
                            <DotsHorizontal fillColor="#FFFFFF" />
                        </div>
                    </div> -->
                </div>
            </div>
        </div>
    </div>

     <div id="OverlaySection" class="fixed top-0 left-0 w-full h-screen bg-black md:bg-slate-500 md:bg-opacity-30 md:p-3">
        <div class="md:max-w-2xl md:mx-auto md:mt-10 md:rounded-xl bg-black">
            <div class="flex justify-between items-center md:inline-block p-2 m-2 rounded-full cursor-pointer">
                <div class="hover:bg-gray-800 inline-block cursor-pointer rounded-full p-2">
                    <Close fillColor="#FFFFFF" :size="28" class="md:block hidden"/>
                    <ArrowLeft fillColor="#FFFFFF" :size="28" class="md:hidden block"/>
                </div>

                <button
                :disabled="!tweet"
                :class="tweet ? 'bg-[#1C9CEF] text-white' : 'bg-[#124D77] text-gray-400'"
                 class="md:hidden font-extrabold px-4 rounded-full cursor-pointer text-[16px] p-1.5">
                    Post
                </button>
            </div>

            <div class="w-full flex">
                <div class="mr-2 ml-3.5">
                    <img :src="randImg1" alt="" class="rounded-full" width="55">
                </div>
                <div class="w-[calc(100%-100px)]">
                    <!-- <div class="inline-block">
                        <div class="flex items-center border border-gray-700 rounded-full">
                            <span class="font-extrabold p-0.5 pl-3.5 text-[#1C9CEF]">Everyone</span>
                            <ChevronDown class="pr-2" :size="25" fillColor="#1C9CEF" />
                        </div>
                    </div> -->

                    <div>
                        <textarea 
                        :oninput="textareaInput"
                        placeholder="What is happening?!"
                        v-model="tweet"
                        ref="textarea"
                        cols="30"
                        rows="10"
                        class="
                        w-full
                    bg-black
                    border-0
                    mt-0
                    focus:outline-none
                    text-gray-300
                    text-[19px]
                    min-h-[120px]">
                </textarea>
                    </div>

                    <div class="w-full">
                        <video controls v-if="uploadType === 'mp4'" :src="showUpload" class="rounded-xl overflow-auto"></video>
                        <img v-else :src="showUpload" alt="" class="rounded-xl min-w-full">
                    </div>
                    <div class="flex py-2 items-center font-extrabold text-[#1C9CEF]">
                        <Earth fillColor="#1C9CEF" :size="20" class="pr-2"/> Everyone can reply
                    </div>
                    <div class="border-b border-gray-700"></div>
                    <div class="flex items-center justify-between py-2">
                        <div class="flex items-center">
                            <div class="inline-block rounded-full cursor-pointer p-2 hover:bg-gray-800">
                                <label for="fileUpload" class="cursor-pointer">
                                    <ImageOutline :size="25" fillColor="#1C9CEF"/>
                                </label>
                                <input type="file" id="fileUpload" class="hidden" @change="getFile">
                            </div>
                            <div class="inline-block rounded-full cursor-pointer p-2 hover:bg-gray-800">
                                    <FileGifBox :size="25" fillColor="#1C9CEF"/>
                            </div>
                            <div class="inline-block rounded-full cursor-pointer p-2 hover:bg-gray-800">
                                    <Emoticon :size="25" fillColor="#1C9CEF"/>
                            </div>
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