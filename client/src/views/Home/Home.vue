<template>
    <div id="home">
        <fieldset>
            <legend>用户信息</legend>
            <p>欢迎VIP:
                <span style="color: #f10;font-size: 20px;font-weight: 600;">{{ route.query.username }}</span>
                登录云熬夜系统,祝您辞藻药丸~~~(doge)
            </p>

        </fieldset>
        <fieldset>
            <legend>系统操作</legend>
            <p>
                <el-button @click="logout">下线</el-button>
                <el-button @click="delAccount">注销</el-button>
            </p>
        </fieldset>
    </div>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()

// 下线
const logout = () => {
    localStorage.removeItem('token')
    router.replace('/login')
}
const delAccount = () => {
    // delete请求后端删除当前用户
    

    localStorage.removeItem('token')
    router.replace('/login')
}

</script>

<style scoped lang='less'>
#home {
    width: 100%;
    height: 100%;
    display: flex;
    gap: 20px;
    flex-direction: column;
    align-items: center;
    background: url('../../assets/images/background.5825f033.svg') no-repeat center center/cover;
}

fieldset {
    &:first-child {
        margin-top: 200px;
    }

    width: 700px;
}
</style>