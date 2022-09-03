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
                    <el-dropdown-item>logout</el-dropdown-item>
                    <el-dropdown-item>注销</el-dropdown-item>
                </el-dropdown-menu>
            </template>
        </el-dropdown>
    </div>
    <!--  -->
</template>

<script setup>

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