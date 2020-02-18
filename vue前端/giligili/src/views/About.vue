<template>
    <el-form ref="form" :model="form" label-width="100px" class="center">
        <el-form-item label="车次">
            <el-input v-model="form.TrainNum" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="始发站">
            <el-input v-model="form.From" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="终点站">
            <el-input v-model="form.To" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="发车日期">
            <el-input v-model="form.Date" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="开车时间">
            <el-input v-model="form.StartTime" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="抵达时间">
            <el-input v-model="form.EndTime" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="历时">
            <el-input v-model="form.Duration" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="商务座数量">
            <el-input type="number" v-model="form.Super" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="一等座数量">
            <el-input type="number" v-model="form.First" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item label="二等座数量">
            <el-input type="number" v-model="form.Second" style="width: 600px"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onSubmit" class="button-css">登记</el-button>
            <el-button class="button-css">取消</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
    import * as API from '@/api/user/';
    export default {
        data() {
            return {
                form: {
                    TrainNum: '',
                    From: '',
                    To: '',
                    StartTime: '',
                    EndTime: '',
                    Duration: '',
                    Date: '',
                    Super: 0,
                    First: 120,
                    Second: 300
                }
            }
        },
        methods: {
            onSubmit() {
                API.input(this.form).then((res) =>{
                    if(res.code > 0){
                        this.$notify.error({
                            title:'新增车次信息失败',
                            message:res.msg,
                        });
                    }else{
                        this.$notify({
                            title: '新增车次信息成功',
                            message: '用户可订购该车次车票',
                            type:'success',
                        });
                        this.$router.replace({path: '/order'})
                    }
                }).catch((error) =>{
                    this.$notify.error({
                        title:'注册失败',
                        message:error,
                    });
                });
            },
        }
    }
</script>

<style>
    .center{
        position:absolute;
        right: 20%;
        left: 20%;
    }
</style>
