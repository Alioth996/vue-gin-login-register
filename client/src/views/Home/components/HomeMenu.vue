<template>
    <div class="home-menu">
        <el-menu active-text-color="#ffd04b" router background-color="#34495e" class="el-menu-vertical-demo"
            default-active="2" text-color="#fff" @open="handleOpen" @close="handleClose">
            <el-sub-menu index="1">
                <template #title>
                    <span>系统设置</span>
                </template>
                <el-menu-item index="1-1">item one</el-menu-item>
                <el-menu-item index="1-2">item two</el-menu-item>
            </el-sub-menu>
            <el-menu-item index="/usermange" v-if="role">
                <!-- 超管和管员可见 -->
                <span>用户管理</span>
            </el-menu-item>
            <el-menu-item index="3">
                <span>个人信息</span>
            </el-menu-item>
        </el-menu>
    </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter()
const role = ref(1)

onMounted(() => {
    pushRoute(role.value)
})

// 根据用户角色添加路由
const pushRoute = (role) => {
    if (!role) {
        return
    }
    router.addRoute('home', {
        path: '/usermange',
        component: () => import('@/views/用户管理/index.vue')
    })

}

const handleOpen = () => {

}
const handleClose = () => {

}



</script>

<style scoped lang='less'>
.home-menu {
    height: 100%;
    background: #34495e;
}
</style>