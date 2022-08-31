<template>
    <!-- 注册 -->
    <div class="form-container sign-up-container">
        <div class="form-box">
            <input type="text" placeholder="e-mail" v-model="userForm.email" />
            <input type="text" placeholder="username" v-model="userForm.username" />
            <input type="password" placeholder="password" v-model="userForm.password" @keyup.enter="registerBefore" />
            <button @click="registerBefore">注册</button>
        </div>
    </div>
</template>
<script setup>

import { ElMessage } from 'element-plus';
import { reactive } from 'vue';

const RegExp = /^\w+@[a-zA-Z0-9]{2,10}(?:\.[a-z]{2,4}){1,3}$/
const emit = defineEmits(['registerCheck'])

const userForm = reactive({
    username: '',
    password: 'admin',
    email: '3$6#xx@qq.com'
})

const registerBefore = () => {
    const re = RegExp.test(userForm.email)
    if (userForm.password.length < 5 || userForm.username.length < 5) {
        // 邮箱特殊校验
        ElMessage.error("请正确填写表单信息!!")
        return
    }
    // 子传父
    emit('registerCheck', userForm)
}
</script>
<style 
lang="less" scoped>
.form-container {
    position: absolute;
    top: 0;
    height: 100%;
    transition: all 0.6s ease-in-out;
}

div.form-box {
    background-color: #ffffff;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    padding: 0 50px;
    height: 100%;
    text-align: center;
}

.sign-in-container {
    left: 0;
    width: 50%;
    z-index: 2;
}

.container.right-panel-active .sign-in-container {
    transform: translateX(100%);
}

.sign-up-container {
    left: 0;
    width: 50%;
    opacity: 0;
    z-index: 1;
}
</style>