<template>
    <div id="login">
        <div class="container">
            <div class="login">
                <div class="title">
                    <img src="../../assets/vue.svg" alt="">
                    <strong>VueGinLR</strong>
                </div>
                <a-form :model="userInfo" name="login" autocomplete="off" @finish="onFinish"
                    @finishFailed="onFinishFailed">
                    <a-form-item name="username" :rules="[{ required: true, message: '请输入账号' }]">
                        <a-input v-model:value="userInfo.username" placeholder="username">
                            <template #prefix>
                                <user-outlined style="font-size: 16px;" />
                            </template>
                        </a-input>
                    </a-form-item>

                    <a-form-item name="password" :rules="[{ required: true, message: '请输入密码' }]">
                        <a-input v-model:value="userInfo.password" placeholder="password" type="password">
                            <template #prefix>
                                <lock-outlined style="font-size: 16px;" />
                            </template>
                        </a-input>
                    </a-form-item>
                    <a-form-item name="verfiyCode" :rules="[{ required: true, message: '请输入验证码' }]">
                        <a-input v-model:value="userInfo.verfiyCode" placeholder="验证码" style="height: 40px;">
                            <template #prefix>
                                <safety-outlined style="font-size: 16px;" />
                            </template>
                            <template #suffix>
                                <img src="../../assets/vue.svg" alt="">
                            </template>
                        </a-input>
                    </a-form-item>
                    <a-form-item>
                        <a-button type="primary" html-type="submit">登录</a-button>
                    </a-form-item>
                    <button @click="router.push({ name: 'home', params: { userInfo } })">主页</button>
                </a-form>
            </div>
        </div>
    </div>
</template>

<script setup>
import { reactive } from 'vue';
import {
    UserOutlined,
    LockOutlined,
    SafetyOutlined

} from '@ant-design/icons-vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const userInfo = reactive({
    username: 'admin',
    password: 'admin',
    verfiyCode: '4396',

})
const onFinish = (values) => {
    // 表单校验成功的回调,拿到表单信息
    // console.log('Success:', values);
    // 1. 拿到表单数据
    router.push({
        name: 'home',
        params: {
            id: '333'
        }
    })
    // 2. RSA密码加密
    // 3. 请求后端接口
    // 4. 登录失败 提示用户原因
    // 5. 成功
    // 5.1 vuex和localstorge存入token信息
    // 5.2 跳转 /home路由  router.replace('/home'), replace是替换路由,不是push到history栈.登录成功之后就不能返回登录页了
    // 5.3 提示登录成功
};

const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
};

</script>

<style scoped lang='less'>
/* 网格背景 */
#login {
    width: 100%;
    height: 100%;
    background-size: 25px 25px;
    background-image: linear-gradient(90deg, rgba(37, 82, 110, 0.1) 1px, transparent 0),
        linear-gradient(180deg, rgba(37, 82, 110, 0.1) 1px, transparent 0);


}

.container {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;

    .login {
        padding: 20px;
        width: 400px;
        background: #eee;

        // 样式穿透 
        & :deep(.ant-input) {
            padding: 3px;
        }

        &>div.title {
            height: 40px;
            text-align: center;
            padding: 10px 0 15px 0;
            box-sizing: content-box;

            img {
                vertical-align: bottom;
                padding-right: 10px;
            }

            strong {
                display: inline-block;
                height: 100%;
                font-size: 24px;
            }
        }
    }
}

// antd reset css
</style>