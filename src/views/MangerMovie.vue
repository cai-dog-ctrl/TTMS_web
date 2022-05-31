<template>
    <div>
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>剧目管理</el-breadcrumb-item>
            <el-breadcrumb-item>剧目列表</el-breadcrumb-item>
        </el-breadcrumb>
        <!--卡片视图-->
        <el-card>
            <!-- 搜索与添加区 -->
            <div v-for="item in first_page_info.ShowingList.ShowingList" :key="item.id" class="movie">
                <el-card :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                        class="image" @click="gotoinfo(item.id)">
                    <div style="padding: 14px;">
                        <span>{{ item.name }}</span>
                        <div class="bottom clearfix">

                            <el-button type="text" class="button" @click="gotoinfo(item.id)">特惠购买
                            </el-button>
                        </div>
                    </div>
                </el-card>
            </div>

            
        </el-card>
        <div class="page">
                <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                    :current-page="queryInfo.pagenum" :page-sizes="[1, 2, 5, 10]" :page-size="queryInfo.pagesize"
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
            addDialogVisible: false,
            queryInfo: {
                key_word: '',
                //当前的页数
                page_num: 1,
                //当前每页有多少条数据
                page_size: 2
            },
            total: '',

            src_url: "https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644",
            first_page_info: {},
        }
    },

    created() {
        this.get_firstPage()
    },

    methods: {
        gotoinfo(id) {
            this.$router.push('/movie/' + id);
        },
        async get_firstPage() {
            const { data: res } = await this.$http.get('GetFrontPage', { params: this.form })
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }
            this.first_page_info = res.data

        },
    },
}
</script>
<style lang="less" scoped>
.el-card {
    margin-top: 15px;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.15) !important;
}

.page {
    margin-top: 15px;
    text-align: center;
    margin: auto;
}

// .main {
//     margin-bottom: 40px;
//     background: #fff;
//     width: 750px;
//     float: left;
//     height: 100%;
// }
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
    width: 200px;
}
</style>