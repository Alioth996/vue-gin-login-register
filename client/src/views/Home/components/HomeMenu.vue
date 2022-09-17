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
            <el-menu-item index="/usermange" v-if="GET_ROLE != 0">
                <!-- 超管和管员可见 -->
                <span>用户管理</span>
            </el-menu-item>
            <el-menu-item index="3">
                <span>个人信息</span>
            </el-menu-item>
        </el-menu>
    </div>
</template>

<script setup >
import { userStore } from '@/store/user';


const { GET_ROLE } = userStore()

console.log(GET_ROLE);



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