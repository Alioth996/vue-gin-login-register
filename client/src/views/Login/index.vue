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
import { onBeforeUnmount, onMounted, ref } from 'vue';
import Register from './components/register.vue';
import Login from './components/login.vue';
import '@/assets/style/login.less'
import { userLogin, userRegister } from '@/api/user';
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router';


const router = useRouter()
let timer = null

const slideContainerState = ref(false)

const loginUser = async (v) => {
    const { data, msg } = await userLogin(v)
    localStorage.setItem('token', data)
    ElMessage.success(msg)
    timer = setTimeout(() => {
        router.replace({
            name: 'home',
            // query正常
            // query: {
            //     username: v.username
            // },
            // // params: 似乎必须使用动态路由才能传参,不然则为空,知识盲区
            // params: {
            //     username: v.username
            // }
        })
        timer = null
        clearTimeout(timer)
    }, 500)


}
const registerUser = async (v) => {
    const { msg } = await userRegister(v)
    ElMessage.success(msg)
    // 注册完成直接登录 
    loginUser(v)
    timer = setTimeout(() => {
        router.replace({
            name: 'home',
            // query: {
            //     username: v.username
            // },

            // params: {
            //     username: v.username
            // }
        })
        timer = null
        clearTimeout(timer)
    }, 500)
}
onBeforeUnmount(() => {
    timer = null
    clearTimeout(timer)
})


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