<template>
    <el-form ref="form" :model="form" label-width="80px" class="center">
        <el-form-item label="车次">
            <el-input v-model="form.trainNum" style="width: 400px"></el-input>
        </el-form-item>
        <el-form-item label="日期">
            <el-date-picker type="date" placeholder="选择日期" value-format="MM-dd" v-model="form.date1" style="width: 400px"></el-date-picker>
        </el-form-item>
        <el-form-item label="起始">
            <el-input v-model="form.from" style="width: 400px"></el-input>
        </el-form-item>
        <el-form-item label="终点">
            <el-input v-model="form.to" style="width: 400px"></el-input>
        </el-form-item>
        <el-form-item label="座位等级">
            <el-radio-group v-model="form.rank">
                <el-radio label="super">商务座</el-radio>
                <el-radio label="first">一等座</el-radio>
                <el-radio label="second">二等座</el-radio>
            </el-radio-group>
        </el-form-item>
        <el-form-item label="用户名">
            <el-input v-model="form.name" style="width: 400px"></el-input>
        </el-form-item>
        <el-form-item label="身份证号">
            <el-input v-model="form.idNum" style="width: 400px"></el-input>
        </el-form-item>
        <el-form-item label="学生票">
            <el-switch v-model="form.student"></el-switch>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onSubmit" :disabled="isDisable">  预定  </el-button>
            <el-button>取消</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
    import * as API from '@/api/user/';

    export default {
        name: "order",
        data() {
            return {
                form: {
                    trainNum: '',
                    name: '',
                    idNum: '',
                    date1: '',
                    student: false,
                    rank: '',
                    from:'',
                    to:''
                },
                isDisable:false
            }
        },
        watch: {
            '$route': 'getParams'
        },
        created(){
            this.getParams()
        },
        methods: {
            getParams () {
                // 取到路由带过来的参数
                this.form.trainNum = this.$route.params.nums;
                this.form.date1 = this.$route.params.date;
                this.form.from = this.$route.params.from;
                this.form.to = this.$route.params.to
            },
            onSubmit(){
                this.isDisable=true;
                setTimeout(()=>{
                    this.isDisable=false   //点击一次时隔两秒后才能再次点击
                },2000);
                API.order(this.form).then((res) =>{
                    if(res.code > 0){
                        this.$notify.error({
                            title:'预定失败',
                            message:res.msg,
                        });
                    }else{
                        this.$notify({
                            title: '预定成功',
                            message: '请注意乘车时间',
                            type:'success',
                        });
                    }
                }).catch((error) =>{
                    this.$notify.error({
                        title:'预定失败',
                        message:error,
                    });
                });
            }
        }
    }
</script>

<style scoped>
    .center{
        position:absolute;
        right: 30%;
        left: 30%;
    }
</style>
