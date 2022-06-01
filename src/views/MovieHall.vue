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
                    <el-input placeholder="请输入内容" v-model="queryInfo.key_word" clearable @clear="getCinemaList">
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
                <el-table-column label="最大行" prop="row"></el-table-column>
                <el-table-column label="最大列" prop="col"></el-table-column>
                <el-table-column label="类型" prop="tag"></el-table-column>
                <el-table-column label="操作" width="180px">
                    <template slot-scope="scope">
                        <!-- 修改影厅按钮 -->
                        <el-tooltip effect="dark" content="修改影厅" placement="top" :enterable="false">
                            <el-button type="primary" icon="el-icon-edit" size="mini"
                                @click="showEditDialog(scope.row.id)">
                            </el-button>
                        </el-tooltip>
                        <!-- 修改座位按钮 -->
                        <el-tooltip effect="dark" content="修改座位" placement="top" :enterable="false">
                            <el-button type="success" icon="el-icon-edit" size="mini"></el-button>
                        </el-tooltip>
                        <!-- 删除按钮 -->
                        <el-button type="danger" icon="el-icon-delete" size="mini"
                            @click="removeCinemaById(scope.row.id)"></el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.page_num" :page-sizes="[1, 2, 5, 10]" :page-size="queryInfo.num"
                layout="total,sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-card>

        <!-- 添加用户的对话框 -->
        <el-dialog title="添加影院" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
            <!--内容主题区-->
            <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="70px">
                <el-form-item label="名称" prop="name">
                    <el-input v-model="addForm.name"></el-input>
                </el-form-item>
                <el-form-item label="最大行" prop="row">
                    <el-input v-model="addForm.roww"></el-input>
                </el-form-item>
                <el-form-item label="最大列" prop="col">
                    <el-input v-model="addForm.coll"></el-input>
                </el-form-item>
                <el-form-item label="类型" prop="tag">
                    <el-input v-model="addForm.tag"></el-input>
                </el-form-item>
            </el-form>
            <!--底部区域-->
            <span slot="footer" class="dialog-footer">
                <el-button @click="addDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addHall">确 定</el-button>
            </span>
        </el-dialog>

        <!-- 修改影厅的对话框 -->
        <el-dialog title="修改影厅" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">

            <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px">
                <el-form-item label="id">
                    <el-input v-model="editForm.id" :disabled="true"></el-input>
                </el-form-item>
                <el-form-item label="名称" prop="name">
                    <el-input v-model="editForm.name"></el-input>
                </el-form-item>
                <el-form-item label="类型" prop="tag">
                    <el-input v-model="editForm.tag"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="editHallInfo">确 定</el-button>
            </span>
        </el-dialog>

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
                num: 2,
            },
            cinemalist: [],
            total: '',

            //控制添加影厅对话框的显示与隐藏
            addDialogVisible: false,
            //添加影厅的表单数据
            addForm: {
                name: '',
                row: '',
                col: '',
                tag: ''
            },
            
            //添加表单的验证规则对象
            addFormRules: {
                name: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 3, max: 10, message: '用户名的长度在3到10之间', trigger: 'blur' }
                ],

                row: [
                    { required: true, message: '请输入行数', trigger: 'blur' },
                    { min: 1, max: 3, message: '行数在1至100之间', trigger: 'blur' }
                ],

                col: [
                    { required: true, message: '请输入列数', trigger: 'blur' },
                    {  min: 1, max: 3, message: '列数在1至100之间', trigger: 'blur' }
                ],

                tag: [
                    { required: true, message: '请输入类型', trigger: 'blur' },
                    { min: 3, max: 10, message: '类型名的长度在3到10之间', trigger: 'blur' }
                ]

            },

            // 控制修改影厅对话框的显示与隐藏
            editDialogVisible: false,
            // 查询到的影厅信息对象
            editForm: {},
             //修改表单的验证规则
            editFormRules: {
                name: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 3, max: 10, message: '用户名的长度在3到10之间', trigger: 'blur' }
                ],

                row: [
                    { required: true, message: '请输入行数', trigger: 'blur' },
                    { min: 1, max: 3, message: '行数在1至100之间', trigger: 'blur' }
                ],

                col: [
                    { required: true, message: '请输入列数', trigger: 'blur' },
                    {  min: 1, max: 3, message: '列数在1至100之间', trigger: 'blur' }
                ],

                tag: [
                    { required: true, message: '请输入类型', trigger: 'blur' },
                    { min: 3, max: 10, message: '类型名的长度在3到10之间', trigger: 'blur' }
                ]
            },



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

        // 监听num
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
        async removeCinemaById(id) {
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
            const { data: res } = await this.$http.get('GetCinemaByID/' + id)
            if (res.code !== 1000) {
                return this.$message.error("删除影院信息失败！")
            }


            const { data: res2 } = await this.$http.put('DeleteCinemaByID/' + id, {
                id: res.data.id,
                // name: res.data.name,
                // row: res.data.row,
                // col: res.data.col,
                // tag: res.data.tag,
                // is_delete: 1

            })
            // console.log(res.data.identity);
            if (res2.code !== 1000) {
                return this.$message.error('删除用户失败！')
            }

            this.$message.success('删除用户成功！')
            this.getCinemaList()

        },

        //监听添加影厅对话框的关闭事件
        addDialogClosed() {
            this.$refs.addFormRef.resetFields()
        },

        //点击按钮，添加新影厅
        async addHall() {
            this.$refs.addFormRef.validate(async valid => {
                // console.log(valid)
                if (!valid) return
                //可以发起添加用户的网络请求
                // console.log(111111);
                const { data: res } = await this.$http.post('AddNewCinema', this.addForm)

                if (res.code !== 1000) {
                    this.$message.error('添加影厅失败！')
                }

                this.$message.success('添加影厅成功！')
                this.addDialogVisible = false
                this.getCinemaList()
            })
        },

         //展示编辑的对话框
        async showEditDialog(id) {
            // console.log(id);
            const { data: res } = await this.$http.get('GetCinemaByID/' + id)

            if (res.code !== 1000) {
                return this.$message.error('查询影厅信息失败！')
            }
        
            this.editForm = res.data
            console.log(this.editForm);
            this.editDialogVisible = true

        },

        //修改影厅表单的预验证
        async editHallInfo() {
            this.$refs.editFormRef.validate(async vaild => {
                
                if (!vaild) return
                //发起修改用户信息的数据请求
                const { data: res } = await this.$http.put('ModifyCinemaByID', {
                    id: this.editForm.id,
                    name: this.editForm.name,
                    row:this.editForm.row,
                    col:this.editForm.col,
                    tag: this.editForm.tag,
                    is_delete: this.editForm.is_delete

                })
                console.log(res.code)
                if (res.code !== 1000) {
                    return this.$message.error('更新用户信息失败！');
                }

                //关闭对话框
                this.editDialogVisible = false
                //刷新数据列表
                this.getCinemaList()
                //提示修改成功
                this.$message.success('更新用户信息成功！')

            })

        },

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