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