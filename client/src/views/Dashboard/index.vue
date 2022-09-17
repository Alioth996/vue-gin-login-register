<template>
    <div class="dashboard">
        <h1>Welcome</h1>
        <br>
    </div>
</template>

<script setup>
import { userStore } from '@/store/user';
import { onMounted } from 'vue';

const { username } = userStore()

// tts语音欢迎
const speakWelcome = (username) => {
    const ttsIntance = new SpeechSynthesisUtterance(`尊敬的用户${username},欢迎使用云信息管理系统`)

    ttsIntance.addEventListener('end', e => console.log('welcome tts end..'))
    ttsIntance.addEventListener('error', e => new Error(e))
    ttsIntance.addEventListener('start', e => console.log("welcome tts start.."))
    // speechSynthesis.speak(ttsIntance)
}

onMounted(() => speakWelcome(username))
</script>

<style scoped>
.dashboard {
    width: 100%;
    height: 100%;
    background: url('../../assets/images/background.5825f033.svg') no-repeat center center/cover;

}
</style>