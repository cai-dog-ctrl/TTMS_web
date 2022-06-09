<template>
    <div>
        <!-- 面包屑导航区域 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>数据统计</el-breadcrumb-item>
            <el-breadcrumb-item>日销售额</el-breadcrumb-item>
        </el-breadcrumb>
        <!-- 卡片视图区 -->
        <el-card v-loading="loading" :element-loading-text="message_update">
            <!-- 为 ECharts 准备一个具备大小（宽高）的 DOM -->
            <div id="main" style="width: 600px;height:auto;">
                <div style="margin-left: 25px;"><span>总销售额：{{this.totalPrice}}元</span></div>
                <div class="sale_item" v-for="item in this.sale" :key="item.index">
                    <div class="img">
                        <img class="image" :src="'http://127.0.0.1:8080/api/getpicturebyfilename/' + item.move_img">
                    </div>
                    <div class="descirption">
                        <div class="movie_name">

                            <span>《{{ item.movie_name }}》</span>
                        </div>
                    </div>
                    <div class="price">
                        <span>￥{{ item.movie_total_price }}</span>

                    </div>
                </div>
            </div>
        </el-card>
    </div>
</template>

<script>
import * as echarts from 'echarts'
export default {
    data() {
        return {
            sale: {},
            totalPrice: '',

            message_updata: "正在查询",
            loading: false,
        }
    },
    created() {
        this.getAllSaleByDay()
    },
    methods: {
        async getAllSaleByDay() {
            this.loading = true
            var myDate = new Date((new Date).getTime());
            var time = myDate.toJSON().split('T').join(' ').substr(0, 10);
           const { data: res } = await this.$http.get('CountSalesByDay/' + time)
            if (res.code !== 1000) {
                this.$message.error("获取日销售额列表失败")
                return
            }
            this.$message.success("获取日销售额成功")
            this.sale = res.data.List
            this.totalPrice = res.data.all_total_price
            this.loading = false
         }
    }
}
</script>

<style lang="less" scoped>
.sale_item {
    width: 90%;
    margin-top: 10px;
    margin-left: auto;
    margin-right: auto;
    border: #E5E5E5 solid 1px;
    border-radius: 10px;
    height: 130px;
    border-radius: 3px;
    display: flex;
    background-color: rgb(126, 206, 193);
}

.img {
    height: 100%;
    margin-left: 20px;
}

.image {
    margin-top: 5px;
    height: 120px;
    width: 110px;
}

.descirption {
    padding-top: 30px;
    width: 250px;
    margin-left: 20px;
}

.movie_name {
    padding-left: 0;
    width: 200px;
}

.price {
    padding-top: 30px;
}
</style>