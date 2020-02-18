<template>
    <el-form  class="form-css" ref="form" :model="form" label-width="80px">
    <el-form-item label="用户名">
            <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码">
            <el-input v-model="form.password" placeholder="请输入密码" show-password></el-input>
        </el-form-item>

        <el-form-item>
            <el-button class="button1-css" type="primary" @click="onSubmit">登录</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
    import * as API from '@/api/user/';

    export default {
        name: 'Login',
        data() {
            return {
                form: {
                    username: '',
                    password: '',
                },
            };
        },
        methods: {
            onSubmit() {
                API.login(this.form).then((res) =>{
                    if(res.code > 0){
                        this.$notify.error({
                            title:'登录失败',
                            message:res.msg,
                        });
                    }else{
                        this.$notify({
                            title: '登陆成功',
                            message: '您可使用用户名与密码登录',
                            type:'success',
                        });
                        this.$router.replace({path: '/'})
                    }
                }).catch((error) =>{
                    this.$notify.error({
                        title:'登录失败',
                        message:error,
                    });
                });
            },
        },
    }
</script>

<style>

    .button1-css{
        width: 360px;
    }

</style>
