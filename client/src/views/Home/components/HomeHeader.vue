<template>
    <div class="home-header">
        <el-dropdown trigger="click">
            <span class="el-dropdown-link">
                <el-avatar shape="square"
                    src="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80" />
                <span class="arrow"></span>
            </span>
            <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item @click="logout">logout</el-dropdown-item>
                    <el-dropdown-item @click="delCurrentUser">注销</el-dropdown-item>
                </el-dropdown-menu>
            </template>
        </el-dropdown>
    </div>
    <!--  -->
</template>

<script setup>
import { ElMessageBox } from 'element-plus';
import { useRouter } from 'vue-router';
import { delUser } from '@/api/user';


const router = useRouter()

let timer = null
const logout = () => {
    timer = setTimeout(() => {
        localStorage.clear()
        router.replace('/login')
    }, 500)
    timer = null
    clearTimeout(timer)
}

const delCurrentUser = () => {
    ElMessageBox.confirm('此操作将彻底删除当前用户!!', "警告", {
        confirmButtonText: '取消',
        cancelButtonText: '确认注销',
        type: 'warning',
        // 区分不同的操作
        distinguishCancelAndClose: true,
    }).then(() => {
        console.log("取消注销");
    }).catch(async () => {
        // console.log("取消");
        // 确认删除
        // 请求后端
        const re = await delUser('65')
        console.log(re);
    })
}
</script>

<style scoped lang='less'>
.home-header {
    display: flex;
    justify-content: end;
    padding: 5px 20px 5px 0;
    box-shadow: 0 1px 4px #00000010;
}

.arrow {
    display: inline-block;
    transform: translate(3px, -10px);
    width: 0;
    height: 0;
    border-left: 4px solid transparent;
    border-right: 4px solid transparent;
    border-top: 4px solid #000;
}
</style>