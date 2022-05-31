<template>
    <div>
        <!--面包屑导航区-->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>影厅管理</el-breadcrumb-item>
            <el-breadcrumb-item>影厅列表</el-breadcrumb-item>
        </el-breadcrumb>
        <!--卡片视图-->
        <el-card>
            <!-- 搜索与添加区 -->
            <el-row :gutter="20">
                <el-col :span="8">
                    <el-input placeholder="请输入内容" v-model="queryInfo.key_word " clearable @clear="getCinemaList">
                        <el-button slot="append" icon="el-icon-search" @click="getCinemaList"></el-button>
                    </el-input>
                </el-col>
                <el-col :span="4">
                    <el-button type="primary" @click="addDialogVisible = true">添加影院</el-button>
                </el-col>
            </el-row>
            <!-- 影院列表区 -->
            <el-table :data="cinemalist" :border='true' :stripe='true'>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="id" prop="id"></el-table-column>
                <el-table-column label="名称" prop="name"></el-table-column>
                <el-table-column label="最大行" prop="roww"></el-table-column>
                <el-table-column label="最大列" prop="coll"></el-table-column>
                <el-table-column label="类型" prop="tag"></el-table-column>
                <el-table-column label="操作" width="180px">
                    <template slot-scope="scope">
                        <!-- 修改按钮 -->
                        <el-button type="primary" icon="el-icon-edit" size="mini" @click="showEditDialog(scope.row.id)">
                        </el-button>
                        <!-- 删除按钮 -->
                        <el-button type="danger" icon="el-icon-delete" size="mini"
                            @click="removeCinemaById(scope.row.id)"></el-button>

                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.page_num" :page-sizes="[1, 2, 5, 10]" :page-size="queryInfo.page_size"
                layout="total,sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-card>


    </div>
</template>

<script>
export default {
    data() {
        return {
            queryInfo: {
                key_word: '',
                //当前的页数
                page_num: 1,
                //当前每页有多少条数据
                num : 2,
            },
            cinemalist: [],
            total: '',

        }
    },

    created() {
        this.getCinemaList()

    },

    methods: {
        async getCinemaList() {
            const { data: res } = await this.$http.get('GetAllCinemas', {
                params: this.queryInfo
            })
            if (res.code != 1000) {
                return this.$message.error('获取影厅列表失败！'); ``

            }
            this.cinemalist = res.data.cinemaList
            this.total = res.data.cinema_num
        },

        // 监听pagesize
        handleSizeChange(newSize) {
            // console.log(newSize)
            this.queryInfo.num = newSize
            this.getCinemaList()
        },
        //监听页码值改变的事件
        handleCurrentChange(newPage) {
            // console.log(newPage)
            this.queryInfo.page_num = newPage
            this.getCinemaList()

        },
        //根据ID删除影厅
        async removeCinemaById(){
            //弹框询问用户是否删除数据
            const confirmResult = await this.$confirm('此操作将永久删除该影厅, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).catch(err => err)


            //如果用户确认删除，则返回为字符串confirm
            //如果用户取消了删除，则返回为字符串cancel
            // console.log(confirmResult);
            if (confirmResult !== 'confirm') {
                return this.$message.info('已取消删除')
            }
            const { data: res } = await this.$http.get('DeleteCinemaByID/' + id)
            if (res.code !== 1000) {
                return this.$message.error("删除用户信息失败！")
            }


            // const { data: res2 } = await this.$http.put('ModifyCinemaByID', {
            //     id: res.data.id,
            //     name: res.data.name,
            //     row: res.data.row,
            //     col: res.data.col,
            //     is_delete: 

            // })
            // console.log(res.data.identity);
            if (res2.code !== 1000) {
                return this.$message.error('删除用户失败！')
            }

            this.$message.success('删除用户成功！')
            this.getCinemaList()

        }

    },
}
</script>

<style lang="less" scoped>
.el-card {
    margin-top: 5px;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.15) !important;
}

.e-table {

    margin-top: 12px;
    font-size: 15px;

}
</style>