<template>
    <div style="text-align: center">
        <el-form :inline="true" :model="formInline" class="demo-form-inline">
            <el-form-item label="日期">
                <el-date-picker type="date" placeholder="选择日期" value-format="MM-dd" v-model="form.Date" style="width: 200px"></el-date-picker>
            </el-form-item>
            <el-form-item>
                <el-select v-model="form.From" filterable placeholder="起始城市">
                    <el-option
                            v-for="item in options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-select v-model="form.To" filterable placeholder="到达城市">
                    <el-option
                            v-for="item in options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">查询</el-button>
            </el-form-item>
        </el-form>

        <el-table
                v-loading="loading"
                :data="tableData"
                border
                style="width: 100%;margin-bottom: 15px;"
                row-key="id"
                lazy
                :load="load"
                :tree-props="{children: 'children', hasChildren: 'hasChildren'}">
            <el-table-column
                    fixed
                    prop="train_num"
                    label="车次"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    prop="from"
                    label="起始城市"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    prop="to"
                    label="到达城市"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    prop="start_time"
                    label="开车时间"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    prop="duration"
                    label="历时"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    prop="super"
                    label="商务座"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    prop="first"
                    label="一等座"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    prop="second"
                    label="二等座"
                    style="width: 14%">
            </el-table-column>
            <el-table-column
                    label="备注"
                    style="width: 16%">
                <template slot-scope="scope">
                    <el-button
                            @click.native.prevent="orderThis(scope.$index, tableData)"
                            type="primary"
                            size="small">
                        预定
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    import * as API from '@/api/user/';
    export default {
        name:'search',
        methods: {
            orderThis(index, rows) {
                this.$router.push({path:'/order', name:'order',
                    params:{
                        nums:rows[index].train_num,
                        date: this.form.Date,
                        from: this.form.From,
                        to: this.form.To
                }});
                // eslint-disable-next-line no-console
                console.log(index);
                // eslint-disable-next-line no-console
                console.log(rows[index])
                //this.$router.replace({path: '/order'})
            },
            onSubmit() {
                API.search(this.form).then((res) =>{
                    if(res.code > 0){
                        this.$notify.error({
                            title:'查询失败',
                            message:res.msg,
                        });
                    }else{
                        this.tableData = res.data
                    }
                }).catch((error) =>{
                    this.$notify.error({
                        title:'查询失败',
                        message:error,
                    });
                });
            },
            load(tree, treeNode, resolve) {
                resolve([
                    {
                        id: 1000,
                        super: '$390',
                        first: '$200',
                        second: '$120'
                    },
                ])
            }
        },
        data() {
            return {
                formInline: {
                    user: '',
                    region: ''
                },
                tableData:[{
                    train_num : ''
                }],
                options: [{
                    value: '北京',
                    label: '北京'
                }, {
                    value: '上海',
                    label: '上海'
                }, {
                    value: '郑州',
                    label: '郑州'
                }, {
                    value: '厦门',
                    label: '厦门'
                }, {
                    value: '武汉',
                    label: '武汉'
                }],
                form:{
                    From: '',
                    To: '',
                    Date:''
                },
            }
        }
    }
</script>

<style>
    .center1{
        position:absolute;
        right: 30%;
        left: 30%;
    }
</style>
