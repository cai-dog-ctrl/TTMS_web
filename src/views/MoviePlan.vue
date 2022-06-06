<template>
    <div>
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>剧目管理</el-breadcrumb-item>
            <el-breadcrumb-item>演出计划</el-breadcrumb-item>
        </el-breadcrumb>
        <!--卡片视图-->
        <el-card class="top_card">
            <!-- 搜索与添加区 -->
            <div v-for="item in movie_list" :key="item.id" class="movie">
                <el-card :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                        class="image">
                    <div style="padding: 10px; display: flex;">
                        <div class="name"><span>{{ item.name }}</span></div>
                        <div class="score"><span>{{ item.score }}分</span></div>
                    </div>

                    <div class="bottom clearfix">
                        <el-button type="success" size="mini" icon="el-icon-edit" @click="aaa(item.id)">
                            添加
                        </el-button>
                        <el-button type="primary" size="mini" icon="el-icon-edit" @click="showEditDialog">
                            修改
                        </el-button>
                        <el-button type="primary" size="mini" icon="el-icon-edit" @click="showEditDialog">
                            查看
                        </el-button>
                        <!-- <el-button type="danger" size="mini" icon="el-icon-delete"
                            @click="deleteMoviePlan">
                            删除
                        </el-button> -->
                    </div>
                </el-card>
                <!-- 添加演出计划的对话框 -->
                <el-dialog title="添加演出计划" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
                    <!--内容主题区-->
                    <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="70px">
                        <el-form-item label="影厅">
                            <el-select v-model="addForm.cinema_id" placeholder="影厅名称">
                                <el-option v-for="item in cinemalist" :key="item.name" :label="item.name"
                                    :value="item.id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="演出日期">
                            <div class="block">
                                <!-- <span class="demonstration">默认</span> -->
                                <el-date-picker v-model="addForm.date_day" type="date" placeholder="选择日期"
                                    value-format="yyyy-MM-dd">
                                </el-date-picker>
                            </div>
                        </el-form-item>
                        <el-form-item label="开始时间">
                            <el-time-picker v-model="addForm.start_time" :picker-options="{
                                selectableRange: '00:00:00 - 23:59:59'
                            }" placeholder="任意时间点" value-format="hh:mm">
                            </el-time-picker>
                        </el-form-item>
                        <el-form-item label="结束时间">
                            <el-time-picker v-model="addForm.end_time" :picker-options="{
                                selectableRange: '00:00:00 - 23:59:59'
                            }" placeholder="任意时间点" value-format="hh:mm">
                            </el-time-picker>
                        </el-form-item>
                        <el-form-item label="票价">
                            <el-input-number v-model="addForm.price" @change="handleChange" :min="10" :max="100">
                            </el-input-number>
                        </el-form-item>
                    </el-form>
                    <!--底部区域-->
                    <span slot="footer" class="dialog-footer">
                        <el-button @click="addDialogVisible = false">取 消</el-button>
                        <el-button type="primary" @click="addPlan">确 定</el-button>
                    </span>
                </el-dialog>

                <!-- 修改演出计划的对话框 -->
                <el-dialog title="修改演出计划" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">
                    <!-- 内容主体区 -->
                    <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px">

                        <el-form-item label="影厅">
                            <el-select v-model="editForm.cinema_id" placeholder="影厅名称">
                                <el-option v-for="item in cinemalist" :key="item.name" :label="item.name"
                                    :value="item.id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="演出日期">
                            <div class="block">
                                <!-- <span class="demonstration">默认</span> -->
                                <el-date-picker v-model="editForm.date_day" type="date" placeholder="选择日期"
                                    value-format="yyyy-MM-dd">
                                </el-date-picker>
                            </div>
                        </el-form-item>
                        <el-form-item label="开始时间">
                            <el-time-picker v-model="editForm.start_time" :picker-options="{
                                selectableRange: '00:00:00 - 23:59:59'
                            }" placeholder="任意时间点" value-format="HH:mm">
                            </el-time-picker>
                        </el-form-item>
                        <el-form-item label="结束时间">
                            <el-time-picker v-model="editForm.end_time" :picker-options="{
                                selectableRange: '00:00:00 - 23:59:59'
                            }" placeholder="任意时间点" value-format="HH:mm">
                            </el-time-picker>
                        </el-form-item>
                        <el-form-item label="票价">
                            <el-input-number v-model="editForm.price" @change="handleChange" :min="10" :max="100">
                            </el-input-number>
                        </el-form-item>

                    </el-form>
                    <span slot="footer" class="dialog-footer">
                        <el-button @click="editDialogVisible = false">取 消</el-button>
                        <el-button type="primary" @click="editUserInfo">确 定</el-button>
                    </span>
                </el-dialog>

                <!-- 查看演出计划的对话框 -->
                 <el-dialog title="查看出计划" :visible.sync="editDialogVisible" width="50%" @close="editDialogClosed">
                    <!-- 内容主体区 -->
                    <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="70px">

                        <el-form-item label="影厅">
                            <el-select v-model="editForm.cinema_id" placeholder="影厅名称">
                                <el-option v-for="item in cinemalist" :key="item.name" :label="item.name"
                                    :value="item.id">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="演出日期">
                            <div class="block">
                                <!-- <span class="demonstration">默认</span> -->
                                <el-date-picker v-model="editForm.date_day" type="date" placeholder="选择日期"
                                    value-format="yyyy-MM-dd">
                                </el-date-picker>
                            </div>
                        </el-form-item>
                        <el-form-item label="开始时间">
                            <el-time-picker v-model="editForm.start_time" :picker-options="{
                                selectableRange: '00:00:00 - 23:59:59'
                            }" placeholder="任意时间点" value-format="HH:mm">
                            </el-time-picker>
                        </el-form-item>
                        <el-form-item label="结束时间">
                            <el-time-picker v-model="editForm.end_time" :picker-options="{
                                selectableRange: '00:00:00 - 23:59:59'
                            }" placeholder="任意时间点" value-format="HH:mm">
                            </el-time-picker>
                        </el-form-item>
                        <el-form-item label="票价">
                            <el-input-number v-model="editForm.price" @change="handleChange" :min="10" :max="100">
                            </el-input-number>
                        </el-form-item>

                    </el-form>
                    <span slot="footer" class="dialog-footer">
                        <el-button @click="editDialogVisible = false">取 消</el-button>
                        <el-button type="primary" @click="editUserInfo">确 定</el-button>
                    </span>
                </el-dialog>
            </div>
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
        return {
            form: {
                CarouselNum: 3,
                ShowingNum: 6,
                ComingNum: 6,
                ScoreRankingNum: 5,
                BoxOfficeRankingNum: 5
            },

            src_url: "https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644",
            first_page_info: {},
            movie_list: {},
            queryInfo: {
                //当前的页数
                Page_num: 1,
                //当前每页有多少条数据
                Num: 5,
            },
            cinemalist: [],
            total: '',

            //控制添加演出计划对话框的显示与隐藏
            addDialogVisible: false,
            //添加演出计划的表单数据
            addForm: {
                cinema_id: '',
                movie_id: '',
                name: '',
                date_day: '',
                start_time: '',
                end_time: '',
                price: '',
            },
            // 控制修改用户对话框的显示与隐藏
            editDialogVisible: false,
            // 查询到的演出计划信息对象
            editForm: {},
    }
},

    created() {
        this.get_firstPage()
        this.getCinemaList()
    },

    methods: {
        aaa(id) {
            this.addDialogVisible = true
            this.addForm.movie_id = id

        },

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

        async getCinemaList() {
            const { data: res } = await this.$http.get('GetAllCinemas', {
                params: this.queryInfo
            })
            if (res.code != 1000) {
                return this.$message.error('获取影厅列表失败！');

            }

            this.cinemalist = res.data.cinemaList
            this.total = res.data.cinema_num
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


        //监听添加用户对话框的关闭事件
        addDialogClosed() {
            this.$refs.addFormRef.resetFields()

        },
        //点击按钮，添加新演出计划
        async addPlan() {
            // this.$refs.addFormRef.validate(async valid => {
            //                 console.log(111);
            //  console.log(valid)
            //     if (!valid) return
            //可以发起添加用户的网络请求
            console.log(this.addForm.date_day);
            const { data: res } = await this.$http.post('addschedule', this.addForm)

            if (res.code !== 1000) {
                this.$message.error('添加演出计划失败！')
            }
            this.$message.success('添加演出计划成功！')
            this.addDialogVisible = false


        },

        async showEditDialog() {
            // console.log(movie_id);
            console.log(this.movie_id);
            const { data: res } = await this.$http.get('getallschedulemsgbymovieid', {
                params: this.movie_id
            })
             console.log(res.code);
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }

            this.editForm = res.data
            console.log(this.editForm);
            this.editDialogVisible = true
        },

        //监听修改用户对话框的关闭事件
        editDialogClosed() {
            this.refs.editFormRef.resetFields()

        },
        async editUserInfo() {

            const { data: res } = await this.$http.put('updateSchedule', this.editForm)
            if (res.code !== 1000) {
                this.$message.error("修改信息失败")
                return
            }
            // this.get_firstPage()
            this.$message.success('修改信息成功')
            this.editDialogVisible = false
        },

        async deleteMoviePlan() {
            const confirmResult = await this.$confirm('此操作将永久删除该计划, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).catch(err => err)
            if (confirmResult !== 'confirm') {
                return this.$message.info('已取消删除')
            }
            const { data: res } = await this.$http.get('getallschedulemsgbymovieid', {
                params: this.movie_id
            })
            if (res.code !== 1000) {
                this.$message.error("删除失败")
                return
            }
            res.data.movie["is_delete"] = 1;
            console.log(res.data);
            const { data: res2 } = await this.$http.put('deleteschedule', res.data.id)
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
    width: 150px;
}

.score {
    margin-left: 20px;
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
    width: 256px;
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
}
</style>