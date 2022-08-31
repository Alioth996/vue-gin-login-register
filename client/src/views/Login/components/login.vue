<template>
    <!-- 登录 -->
    <div class="form-container sign-in-container">
        <form action="#">
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Password" />
            <br>
            <button>登录</button>
        </form>
    </div>
</template>

<style 
lang="less" scoped>
.form-container {
    position: absolute;
    top: 0;
    height: 100%;
    transition: all 0.6s ease-in-out;
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