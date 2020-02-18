<template>
    <el-menu class="el-menu-demo el-menu-css" mode="horizontal" router >
        <el-menu-item index="/">首页</el-menu-item>
        <el-menu-item index="/about">关于我们</el-menu-item>
        <el-menu-item index="/search" class="search-css">
            <i class="el-icon-search"></i>
        </el-menu-item>
        <el-menu-item index="/register" class="register-css">注册</el-menu-item>
        <el-menu-item index="/index" class="login-css">登录</el-menu-item>
    </el-menu>
</template>

<script>
    export default {
        name:'NavBar',
    };
</script>

<style scoped>
    .el-menu-css{
        background-color:#000000;
        margin-bottom: 10px;
        border-radius: 5px;
    }

    .login-css{
        position: absolute;
        right: 0px;
    }

    .register-css{
        position: absolute;
        right: 70px;
    }

    .search-css{
        position: absolute;
        right: 140px;
    }
</style>
