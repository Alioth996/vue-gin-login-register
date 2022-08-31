<template>
    <div id="login">
        <div class="container" :class="{ 'right-panel-active': slideContainerState }" id="container">
            <!-- 注册 -->
            <Register />
            <Login />

            <div class="overlay-container">
                <div class="overlay">
                    <div class="overlay-panel overlay-left">
                        <button class="ghost" id="signIn" @click="slideContainerState = false">登录</button>
                    </div>
                    <div class="overlay-panel overlay-right">
                        <button class="ghost" id="signUp" @click="slideContainerState = true">注册</button>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script setup>
import { ref } from 'vue';
import Register from './components/register.vue';
import Login from './components/login.vue';
import '@/assets/style/login.less'

const slideContainerState = ref(false)

</script>

<style  lang='less'>
#login {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>