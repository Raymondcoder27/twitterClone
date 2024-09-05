<script>
    import { toRefs } from 'vue';
    import Home from 'vue-material-design-icons/Home.vue'
    import Pound from 'vue-material-design-icons/Pound.vue'
    import BellOutline from 'vue-material-design-icons/BellOutline.vue'
    import AccountOutline from 'vue-material-design-icons/AccountOutline.vue'
    import EmailOutline from 'vue-material-design-icons/EmailOutline.vue'

    const props = defineProps({
        iconString: String,
    })

    const {iconString} = toRefs(props)

    let icon = null

    if (iconString.value === 'Home') icon = Home
    if (iconString.value === 'Pound') icon = Pound
    if (iconString.value === 'BellOutline') icon = BellOutline
    if (iconString.value === 'AccountOutline') icon = AccountOutline
    if (iconString.value === 'EmailOutline') icon = EmailOutline


    
</script>

<template>
    
</template>