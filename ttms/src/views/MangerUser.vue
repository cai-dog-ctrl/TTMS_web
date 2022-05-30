<template>
    <div>
        <!--面包屑导航区-->
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
                    <el-button type="primary" @click="addDialogVisible = true">添加用户</el-button>
                </el-col>
            </el-row>

            <!-- 用户列表区 -->
            <el-table :data="userlist" :border='true' :stripe='true'>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="姓名" prop="username"></el-table-column>
                <el-table-column label="密码" prop="password"></el-table-column>
                <el-table-column label="邮箱" prop="email"></el-table-column>
                <el-table-column label="电话" prop="phone_number"></el-table-column>
                <el-table-column label="角色" prop="identity"></el-table-column>
                <el-table-column label="状态">
                    <template slot-scope="scope">
                        <el-switch v-model="scope.row.is_login" @change="userStateChanged(scope.row.id)" ></el-switch>
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="180px">
                    <template slot-scope="scope">
                        <!-- 修改按钮 -->
                        <el-button type="primary" icon="el-icon-edit" size="mini" @click="showEditDialog(scope.row.id)">
                        </el-button>
                        <!-- 删除按钮 -->
                        <el-button type="danger" icon="el-icon-delete" size="mini"
                            @click="removeUserById(scope.row.id)"></el-button>
                        <!-- 分配角色按钮 -->
                        <el-tooltip effect="dark" content="分配角色" placement="top" :enterable="false">
                            <el-button type="warning" icon="el-icon-setting" size="mini"></el-button>
                        </el-tooltip>

                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.page_num" :page-sizes="[1, 2, 5, 10]" :page-size="queryInfo.page_size"
                layout="total,sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-card>

        <!-- 添加用户的对话框 -->
        <el-dialog title="添加用户" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
            <!--内容主题区-->
            <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="70px">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="addForm.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="addForm.password"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="addForm.email"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="phone_number">
                    <el-input v-model="addForm.phone_number"></el-input>
                </el-form-item>
            </el-form>
            <!--底部区域-->
            <span slot="footer" class="dialog-footer">
                <el-button @click="addDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addUser">确 定</el-button>
            </span>
        </el-dialog>

        <!-- 修改用户的对话框 -->
        <el-dialog title="修改用户" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">

            <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px">
                <el-form-item label="id">
                    <el-input v-model="editForm.id" :disabled="true"></el-input>
                </el-form-item>
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="editForm.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="editForm.password"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="editForm.email"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="phone_number">
                    <el-input v-model="editForm.phone_number"></el-input>
                </el-form-item>
                <el-form-item label="角色" prop="identity">
                    <el-input v-model="editForm.identity"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="editUserInfo">确 定</el-button>
            </span>
        </el-dialog>
    </div>

</template>

<script>
export default {
    data() {
        //验证邮箱的规则
        var checkEmail = (rule, value, cb) => {
            //验证邮箱的正则表达式 
            const regEmail = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/
            if (regEmail.test(value)) {
                //合法邮箱
                return cb()
            }

            cb(new Error('请输入合法的邮箱'))

        }

        //验证手机号的规则
        var checkMobile = (rule, value, cb) => {
            //验证手机号的正则表达式
            const regMobile = /^((13[0-9])|(17[0-1,6-8])|(15[^4,\\D])|(18[0-9]))\d{8}$/
            if (regMobile.test(value)) {
                //合法手机号
                return cb()
            }

            cb(new Error('请输入合法的手机号'))
        }

        return {
            queryInfo: {
                key_word: '',
                //当前的页数
                page_num: 1,
                //当前每页有多少条数据
                page_size: 2
            },
            userlist: [],
            total: '',
            //控制添加用户对话框的显示与隐藏
            addDialogVisible: false,
            //添加用户的表单数据
            addForm: {
                username: '',
                password: '',
                email: '',
                phone_number: ''
            },
            //添加表单的验证规则对象
            addFormRules: {
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 3, max: 10, message: '用户名的长度在3到10之间', trigger: 'blur' }
                ],

                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    { min: 6, max: 15, message: '密码的长度在6到15之间', trigger: 'blur' }
                ],

                email: [
                    { required: true, message: '请输入邮箱', trigger: 'blur' },
                    { validator: checkEmail, trigger: 'blur' }
                ],

                phone_number: [
                    { required: true, message: '请输入手机号码', trigger: 'blur' },
                    { validator: checkMobile, trigger: 'blur' }
                ]

            },

            // 控制修改用户对话框的显示与隐藏
            editDialogVisible: false,
            // 查询到的用户信息对象
            editForm: {},
            //修改表单的验证规则
            editFormRules: {
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 3, max: 10, message: '用户名的长度在3到10之间', trigger: 'blur' }
                ],
                email: [
                    { required: true, message: '请输入用户的邮箱', trigger: 'blur' },
                    { validator: checkEmail, trigger: 'blur' }
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    { min: 6, max: 15, message: '密码的长度在6到15之间', trigger: 'blur' }
                ],

                phone_number: [
                    { required: true, message: '请输入用户的手机号码', trigger: 'blur' },
                    { validator: checkMobile, trigger: 'blur' }
                ]


            },
            



        }
    },
    created() {
        this.getUserList()
    },
    methods: {
        async getUserList() {
            const { data: res } = await this.$http.get('getallmsg', {
                params: this.queryInfo
            })
            if (res.code != 1000) {
                return this.$message.error('获取用户列表失败！'); ``

            }
            this.userlist = res.data.list

            this.total = res.data.Total
            //console.log(res)
        },
        // 监听pagesize
        handleSizeChange(newSize) {

            // console.log(newSize)
            this.queryInfo.page_size = newSize
            this.getUserList()
        },
        //监听页码值改变的事件
        handleCurrentChange(newPage) {
            // console.log(newPage)
            this.queryInfo.page_num = newPage
            this.getUserList()

        },
        //监听switch开关状态的改变
        async userStateChanged(id) {
            const{data: res} = await this.$http.get('getusermsgbyid/' + id)

            if(res.code !== 1000){
                return this.$message.error("更新用户信息失败！")
            }
            
            if(res.is_login === 1){

                res.is_login = -1
            }
            if(res.is_login === -1){

                res.is_login = 1
            }
            

        },

        //监听添加用户对话框的关闭事件
        addDialogClosed() {
            this.$refs.addFormRef.resetFields()

        },

        //点击按钮，添加新用户
        addUser() {
            this.$refs.addFormRef.validate(async valid => {
                // console.log(valid)
                if (!valid) return
                //可以发起添加用户的网络请求
                const { data: res } = await this.$http.post('addadmin/', this.addForm)

                if (res.code !== 1000) {
                    this.$message.error('添加用户失败！')
                }

                this.$message.success('添加用户成功！')
                this.addDialogVisible = false
                this.getUserList()
            })
        },

        //展示编辑的对话框
        async showEditDialog(id) {
            // console.log(id);
            const { data: res } = await this.$http.get('getusermsgbyid/' + id)

            if (res.code !== 1000) {
                return this.$message.error('查询用户信息失败！')
            }

            this.editForm = res.data
            console.log(this.editForm);
            this.editDialogVisible = true

        },

        //监听修改用户对话框的关闭事件
        editDialogClosed() {
            this.refs.editFormRef.resetFields()

        },
        //修改用户表单的预验证
        editUserInfo() {
            this.$refs.editFormRef.validate(async vaild => {

                if (!vaild) return
                //发起修改用户信息的数据请求
                const { data: res } = await this.$http.put('updatemsg/' +
                    this.editForm.id, {
                    id: this.editForm.id,
                    username: this.editForm.username,
                    password: this.editForm.password,
                    email: this.editForm.email,
                    phone_number: this.editForm.phone_number,
                    identity: this.editForm.identity

                })
                console.log(res.code)
                if (res.code !== 1000) {
                    return this.$message.error('更新用户信息失败！');
                }

                //关闭对话框
                this.editDialogVisible = false
                //刷新数据列表
                this.getUserList()
                //提示修改成功
                this.$message.success('更新用户信息成功！')

            })

        },

        //根据ID删除用户信息
        async removeUserById(id) {
            //弹框询问用户是否删除数据
            const confirmResult = await this.$confirm('此操作将永久删除该用户, 是否继续?', '提示', {
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

            const { data: res } = await this.$http.delete('updatemsg/' + id)

            if (res.code !== 1000) {
                return this.$message.error('删除用户失败！')
            }

            this.$message.success('删除用户成功！')
            this.getUserList()
        }
    }
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