<template>
    <div>
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>剧目管理</el-breadcrumb-item>
            <el-breadcrumb-item>剧目列表</el-breadcrumb-item>
        </el-breadcrumb>
        <!--卡片视图-->
        <el-card class="top_card">
            <el-row :gutter="20">
                
                <el-col :span="4">
                    <el-button type="primary" @click="addDialogVisible = true">上架电影</el-button>
                </el-col>
            </el-row>
             <el-divider content-position="left">少年包青天</el-divider>
            <!-- 搜索与添加区 -->
            <div v-for="item in movie_list" :key="item.id" class="movie">
                <el-card :body-style="{ padding: '0px' }" shadow="hover">
                    <div><img
                            src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                            class="image">
                        <div style="padding: 8px;display: flex;">
                            <div class="name"><span>{{ item.name }}</span></div>

                            <div class="score"><span>{{ item.score }}分</span></div>


                        </div>
                        <div class="bottom clearfix">
                            <el-button type="primary" size="mini" icon="el-icon-edit" @click="showEditDialog(item.id)">
                                修改
                            </el-button>
                            <el-button type="danger" size="mini" icon="el-icon-delete" @click="deleteMovie(item.id)">删除
                            </el-button>
                        </div>
                    </div>

                </el-card>
            </div>

            <el-dialog title="修改用户" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">

                <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px">
                    <el-form-item label="id">
                        <el-input v-model="editForm.id" :disabled="true"></el-input>
                    </el-form-item>
                    <el-form-item label="电影名" prop="name">
                        <el-input v-model="editForm.name"></el-input>
                    </el-form-item>
                    <el-form-item label="描述" prop="descrption">
                        <el-input v-model="editForm.description"></el-input>
                    </el-form-item>
                    <el-form-item label="类型" prop="tag">
                        <el-input v-model="editForm.tag"></el-input>
                    </el-form-item>
                    <el-form-item label="票房" prop="boxoffice">
                        <el-input-number v-model="editForm.box_office" :precision="2" :step="0.1" :min="1" :max="500">
                        </el-input-number>
                    </el-form-item>
                    <el-form-item label="评分" prop="score">
                        <el-input-number v-model="editForm.score" :precision="2" :step="0.1" :min="1" :max="10">
                        </el-input-number>
                    </el-form-item>

                </el-form>
                <span slot="footer" class="dialog-footer">
                    <el-button @click="editDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="editUserInfo">确 定</el-button>
                </span>
            </el-dialog>

        
        </el-card>
        <div class="page">
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.Page_num" :page-sizes="[5, 10, 12, 20]" :page-size="queryInfo.Num"
                layout="total,sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </div>


    </div>
</template>
<script>
export default {
    data() {
        const checkNumber = (rule, value, callback) => {

            value = parseInt(value)
            console.log(value);
            if (!Number.isInteger(value)) {
                callback(new Error("请输入数值"));
            }
            else if (value < 0 || value > 100) {
                callback(new Error("票房在0-99之间"));
            }
            else {
                callback();
            }
        };
        return {
             addForm: {
                name: '',
                description: '',
                tag: '',
                down_time: '',
                duration:1,
                up_time:'',
                cover_img_path:''
            },
            editDialogVisible: false,
            editFormRules: {
                name: [
                    { required: true, message: '请输入电影名', trigger: 'blur' },
                    { min: 3, max: 10, message: '电影名的长度在3到10之间', trigger: 'blur' }
                ],
                descrption: [
                    { required: true, message: '请输入电影描述', trigger: 'blur' },
                    { min: 3, max: 200, message: '电影描述的长度在3到200之间', trigger: 'blur' }
                ],
                boxoffice: [
                    { required: true, message: '请输入票房', trigger: 'blur' },
                    { validator: checkNumber, trigger: 'blur' }
                ],

                score: [
                    { required: true, message: '请输入评分', trigger: 'blur' },

                ],
                tag: [
                    { required: true, message: '请输入电影类型', trigger: 'blur' },
                    { min: 0, max: 10, message: '电影类型的长度在0到10之间', trigger: 'blur' }
                ]

            },

            // 查询到的用户信息对象
            editForm: {},
            form: {
                CarouselNum: 3,
                ShowingNum: 6,
                ComingNum: 6,
                ScoreRankingNum: 5,
                BoxOfficeRankingNum: 5
            },
            queryInfo: {

                //当前的页数
                Page_num: 1,
                //当前每页有多少条数据
                Num: 5
            },
            total: '',

            src_url: "https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644",
            movie_list: {},
        }
    },

    created() {
        this.get_firstPage()
    },

    methods: {
        async get_firstPage() {
            console.log(this.queryInfo);
            const { data: res } = await this.$http.get('GetAllMovies', { params: this.queryInfo })
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }
            console.log(res.data);
            this.movie_list = res.data.movieList
            this.total = res.data.total
        },
        // 监听pagesize
        handleSizeChange(newSize) {

            // console.log(newSize)
            this.queryInfo.Num = newSize
            this.get_firstPage()
        },
        //监听页码值改变的事件
        handleCurrentChange(newPage) {
            // console.log(newPage)
            this.queryInfo.Page_num = newPage
            this.get_firstPage()

        },
        async showEditDialog(id) {
            const { data: res } = await this.$http.get('GetMovieInfoByID/' + id)
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }
            this.editForm = res.data.movie
            this.editDialogVisible = true
        },
        async editUserInfo(){
            
            const { data: res } = await this.$http.put('ModifyMovie',this.editForm)
            if (res.code !== 1000) {
                this.$message.error("修改信息失败")
                return
            }
            this.get_firstPage()
            this.$message.success('修改信息成功')
            this.editDialogVisible = false
        },
        async deleteMovie(id){
            const confirmResult = await this.$confirm('此操作将永久删除该电影, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).catch(err => err)
            if (confirmResult !== 'confirm') {
                return this.$message.info('已取消删除')
            }
            const { data: res } = await this.$http.get('GetMovieInfoByID/' + id)
            if (res.code !== 1000) {
                this.$message.error("删除失败")
                return
            }
            res.data.movie["is_delete"]=1;
            console.log(res.data);
             const { data: res2 } = await this.$http.put('ModifyMovie',res.data.movie)
            if (res2.code !== 1000) {
                this.$message.error("修改信息失败")
                return
            }
            this.get_firstPage()
            this.$message.success('删除成功')
        }
    },
}
</script>
<style lang="less" scoped>
.name {
    width: 130px;
}

.score {
    margin-left: 12px;
}

.el-card {
    margin-top: 15px;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.15) !important;
}

.page {
    margin-top: 15px;
    text-align: center;
    margin: auto;
}

.movie {
    width: 200px;
    float: left;
    margin-left: 40px;
    margin-top: 10px;

    span {
        text-align: center;
    }

}

.image {

    cursor: pointer;
    position: relative;
    width: 200px;
}
</style>