<template>
    <!-- 登录 -->
    <div class="form-container sign-in-container">
        <div class="form-box">
            <input type="text" placeholder="username" v-model="userForm.username" required />
            <input type="password" placeholder="password" v-model="userForm.password" required
                @keyup.enter="loginBefore" />
            <br>
            <button @click="loginBefore">登录</button>
        </div>

    </div>
</template>

<script setup>
import { ElMessage } from 'element-plus';
import { reactive } from 'vue';


const emit = defineEmits(['loginCheck'])

const userForm = reactive({
    username: 'admin',
    password: 'admin',
})


const loginBefore = () => {
    if (userForm.password.length < 5 || userForm.username.length < 5) {
        ElMessage.error('请正确填写用户名或密码!!')
        return
    }
    // 子传父
    emit('loginCheck', userForm)
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
</style>