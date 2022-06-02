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
            <div v-for="item in first_page_info.ShowingList.ShowingList" :key="item.id" class="movie">
                <el-card :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                        class="image">
                    <div style="padding: 14px;">
                        <span>{{ item.name }}</span>
                        <div class="bottom clearfix">
                            <el-button type="success" size="mini" icon="el-icon-edit" @click="aaa(item.movie_id)" >
                                添加
                            </el-button>
                        </div>
                    </div>
                </el-card>
            </div>
        </el-card>

        <!-- 添加演出计划的对话框 -->
        <el-dialog title="添加演出计划" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
            <!--内容主题区-->
            <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="70px">
                <el-form-item label="影厅">
                    <el-select v-model="addForm.cinema_id" placeholder="影厅名称">
                        <el-option v-for="item in cinemalist" :key="item.name" :label="item.name" :value="item.id">
                        </el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="演出日期">
                    <div class="block">
                        <!-- <span class="demonstration">默认</span> -->
                        <el-date-picker v-model="addForm.date_day" type="date" placeholder="选择日期"
                            value-format="yyyy-mm-dd">
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
                    <el-input-number v-model="addForm.price" @change="handleChange" :min="1" :max="10">
                    </el-input-number>
                </el-form-item>
            </el-form>
            <!--底部区域-->
            <span slot="footer" class="dialog-footer">
                <el-button @click="addDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addPlan">确 定</el-button>
            </span>
        </el-dialog>

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
                key_word: '',
                //当前的页数
                Page_num: 1,
                //当前每页有多少条数据
                Num: 2
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
        }
    },

    created() {
        this.get_firstPage()
        this.getCinemaList()
        this.getMovieList()

    },

    methods: {
        aaa(id){
            this.addDialogVisible = true
        },

        async get_firstPage() {
            const { data: res } = await this.$http.get('GetFrontPage', { params: this.form })
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }
            this.first_page_info = res.data

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

        async getMovieList() {
            console.log(this.queryInfo);
            const { data: res } = await this.$http.get('GetAllMovies', {
                params: this.queryInfo
            })
            if (res.code !== 1000) {
                return this.$message.error("获取信息失败")

            }
            console.log(res.data);
            this.movie_list = res.data.movieList
            this.total = res.data.total
        },


        //监听添加用户对话框的关闭事件
        addDialogClosed() {
            this.$refs.addFormRef.resetFields()

        },
        //点击按钮，添加新演出计划
        async addPlan() {
            this.$refs.addFormRef.validate(async valid => {
                // console.log(valid)
                if (!valid) return
                //可以发起添加用户的网络请求
                console.log(this.addForm);

                const { data: res } = await this.$http.post('addschedule', this.addForm)

                if (res.code !== 1000) {
                    this.$message.error('添加演出计划失败！')
                }

                this.$message.success('添加演出计划成功！')
                this.addDialogVisible = false
                this.getCinemaList()
                this.getMovieList()

            })
        },
    }
}
</script>

<style lang="less" scoped>
.el-card {
    margin-top: 15px;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.15) !important;
}

.movie {
    width: 235px;
    float: left;
    margin-left: 15px;
    margin-top: 20px;

    span {
        text-align: center;
    }

}

.image {

    cursor: pointer;
    position: relative;
}
</style>