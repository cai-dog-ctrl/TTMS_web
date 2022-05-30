<template>
    <div>
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>用户管理</el-breadcrumb-item>
            <el-breadcrumb-item>用户列表</el-breadcrumb-item>
        </el-breadcrumb>
        <!--卡片视图-->
        <el-card>
            <!-- 搜索与添加区 -->
            <el-row :gutter="20">
                <el-col :span="8">
                    <el-input placeholder="请输入内容" v-model="queryInfo.key_word" clearable @clear="getUserList">
                        <el-button slot="append" icon="el-icon-search" @click="getUserList"></el-button>
                    </el-input>
                </el-col>
                <el-col :span="4">
                    <el-button type="primary" @click="addDialogVisible = true">上架剧目</el-button>
                </el-col>
            </el-row>

            <div class="page">
                <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                    :current-page="queryInfo.pagenum" :page-sizes="[1, 2, 5, 10]" :page-size="queryInfo.pagesize"
                    layout="total,sizes, prev, pager, next, jumper" :total="total">
                </el-pagination>
            </div>

        </el-card>

    </div>
</template>
<script>
export default {
    data() {
        return {
            addDialogVisible: false,
            queryInfo: {
                key_word: '',
                //当前的页数
                page_num: 1,
                //当前每页有多少条数据
                page_size: 2
            },
        }
    }
}
</script>
<style lang="less" scoped>
.el-card {
    margin-top: 15px;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.15) !important;
}
.page{
    margin-top: 15px;
    text-align: center;
    margin: auto;
}
</style>