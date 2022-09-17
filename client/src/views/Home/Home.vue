<template>
    <div id="home">
        <el-container>
            <el-aside width="220px">
                <home-menu />
            </el-aside>
            <el-container>
                <el-header>
                    <home-header />
                </el-header>
                <el-main>
                    <router-view />
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<script setup>
import HomeHeader from "./components/HomeHeader.vue";
import HomeMenu from "./components/HomeMenu.vue";
import { getUserRole } from "@/api/user";
import { userStore } from "@/store/user";
import { onMounted } from "vue";
import { useRouter } from 'vue-router'

const router = useRouter()

const { SET_ROLE } = userStore()

onMounted(async () => {
    const { data } = await getUserRole()
    SET_ROLE(data.role)
    pushRoute(data.role)

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

</script>

<style scoped lang='less'>
#home {
    width: 100%;
    height: 100%;

    .el-container {
        width: 100%;
        height: 100%;
    }

    .el-header {
        padding: 0;
    }
}
</style>