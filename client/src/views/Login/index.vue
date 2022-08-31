<template>
    <div id="login">
        <div class="container" :class="{ 'right-panel-active': slideContainerState }" id="container">
            <Register @registerCheck="registerUser" />
            <Login @loginCheck="loginUser" />
            <!-- 遮罩 -->
            <div class="overlay-container">
                <div class="overlay">
                    <div class="overlay-panel overlay-left">
                        <button class="ghost" @click="slideContainerState = false">登录</button>
                    </div>
                    <div class="overlay-panel overlay-right">
                        <button class="ghost" @click="slideContainerState = true">注册</button>
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
import { userLogin, userRegister } from '@/api/user';
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router';


const router = useRouter()

const slideContainerState = ref(false)

const loginUser = async (v) => {
    const { data, msg } = await userLogin(v)
    localStorage.setItem('token', data)
    ElMessage.success(msg)
    router.replace({
        name: 'home',
        // query: {
        //     username: v.username
        // },
        params: {
            username: v.username
        }
    })
}
const registerUser = async (v) => {
    const { msg } = await userRegister(v)
    ElMessage.success(msg)
    // 注册完成直接登录 
    router.replace({
        name: 'home',
        query: {
            username: v.username
        },
        params: {
            username: v.username
        }
    })
}



</script>

<style  lang='less' scoped>
#login {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-size: 20px 25px;
    background-image: linear-gradient(90deg, #cccccc60 1px, transparent 0), linear-gradient(180deg, #cccccc60 1px, transparent 0)
}
</style>