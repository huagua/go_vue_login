<template>
    <el-form class="form-css" ref="form" :model="form" label-width="80px">
        <el-form-item label="昵称">
            <el-input v-model="form.nickname" placeholder="请输入昵称"></el-input>
        </el-form-item>
        <el-form-item label="用户名">
            <el-input v-model="form.user_name" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码">
            <el-input v-model="form.password" placeholder="请输入密码" show-password></el-input>
        </el-form-item>
        <el-form-item label="确认密码">
            <el-input v-model="form.password_confirm" placeholder="再次输入密码" show-password></el-input>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="onSubmit" class="button-css">注册</el-button>
            <el-button class="button-css">取消</el-button>
        </el-form-item>
    </el-form>

</template>

<script>
    import * as API from '@/api/user/';

    export default {
        name: 'Register',
        data() {
            return {
                form: {
                    nickname: '',
                    user_name: '',
                    password: '',
                    password_confirm: '',
                },
            };
        },
        methods: {
            onSubmit() {
                API.register(this.form).then((res) =>{
                    if(res.code > 0){
                        this.$notify.error({
                            title:'注册失败',
                            message:res.msg,
                        });
                    }else{
                        this.$notify({
                            title: '注册成功',
                            message: '您可使用用户名与密码登录',
                            type:'success',
                        });
                        this.$router.replace({path: '/index'})
                    }
                }).catch((error) =>{
                    this.$notify.error({
                        title:'注册失败',
                        message:error,
                    });
                });
            },
        },
    }
</script>

<style>
    .form-css{
        margin-top: 70px;
        position:absolute;
        right: 32%;
        left: 32%;
    }

    .button-css{
        width: 178px;
    }
</style>
