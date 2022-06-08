<template>
    <div>


        <div v-if="orderInfo.status === -1">
            <div class="steps">
                <el-steps :active="1" finish-status="success" align-center>
                    <el-step title="选择座位"></el-step>
                    <el-step title="支付订单"></el-step>
                    <el-step title="支付完成"></el-step>
                </el-steps>
            </div>
            <div class="order_id">
                <div class="order_id_left"></div>
                <span>西邮订单号:{{ orderInfo.order_id }}</span>
                <div class="order_id_right"><span>（有订单问题可拨打西邮客服电话1010-5335，工作时间: 9:00-24:00）</span></div>

            </div>
            <div class="order">
                <div class="attribute">
                    <div class="movie_name">电影名</div>
                    <div class="movie_time">时间</div>
                    <div class="movie_ciname">影厅</div>
                    <div class="movie_seat">座位</div>
                </div>
                <div class="content">
                    <div class="name">《{{ orderInfo.movie_name }}》</div>
                    <div class="time">{{ orderInfo.schedule_date }} {{ orderInfo.schedule_time }}</div>
                    <div class="ciname">{{ orderInfo.cinema_name }}</div>
                    <div class="seat">
                        <span v-for="item in orderInfo.seat_list"
                            :key="item">|{{ item.row + 1 }}行{{ item.col + 1 }}列|&nbsp;</span>

                    </div>
                </div>
            </div>
            <div class="bottom">
                <div class="bottom_left">
                    <div class="bottom_left1">
                        <span>奥斯卡长安国际影城（中央广场店）</span>
                    </div>
                    <br>
                    <div class="bottom_left2">地址: 西长安街105号长安中央广场A座2楼（西北饭店斜对面）</div>
                    <div class="bottom_left3">电话: 029-84192558/029-84192559</div>
                </div>
                <div class="bottom_right">
                    总价：<span>98</span>
                    <el-button type="success">点我支付哟~</el-button>
                </div>
            </div>
        </div>
        <div v-if="orderInfo.status === 1">
            <div class="steps">
                <el-steps :active="3" finish-status="success" align-center>
                    <el-step title="选择座位"></el-step>
                    <el-step title="支付订单"></el-step>
                    <el-step title="支付完成"></el-step>
                </el-steps>
            </div>
            <div class="order_id">
                <div class="order_id_left"></div>
                <span>西邮订单号:{{ orderInfo.order_id }}</span>
                <div class="order_id_right"><span>（有订单问题可拨打西邮客服电话1010-5335，工作时间: 9:00-24:00）</span></div>

            </div>
            <div class="order">
                <div class="attribute">
                    <div class="movie_name">电影名</div>
                    <div class="movie_time">时间</div>
                    <div class="movie_ciname">影厅</div>
                    <div class="movie_seat">座位</div>
                </div>
                <div class="content">
                    <div class="name">《{{ orderInfo.movie_name }}》</div>
                    <div class="time">{{ orderInfo.schedule_date }} {{ orderInfo.schedule_time }}</div>
                    <div class="ciname">{{ orderInfo.cinema_name }}</div>
                    <div class="seat">
                        <span v-for="item in orderInfo.seat_list"
                            :key="item">|{{ item.row + 1 }}行{{ item.col + 1 }}列|&nbsp;</span>
                    </div>
                </div>
            </div>
            <div class="bottom">
                <div class="bottom_left">
                    <div class="bottom_left1">
                        <span>奥斯卡长安国际影城（中央广场店）</span>
                    </div>
                    <br>
                    <div class="bottom_left2">地址: 西长安街105号长安中央广场A座2楼（西北饭店斜对面）</div>
                    <div class="bottom_left3">电话: 029-84192558/029-84192559</div>
                </div>
                <div class="bottom_right">
                    总价：<span>98</span>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            orderInfo: {}
        }
    },
    created() {
        this.getOrderInfo()
    },
    methods: {
        async getOrderInfo() {
            var id = this.$route.params.id
            const { data: res } = await this.$http.get('GetOrderByID/' + id)
            if (res.code !== 1000) {
                this.$message.error("获取订单详情失败")

                return
            }
            this.orderInfo = res.data
        }
    }
}
</script>
<style lang="less" scoped>
.steps {
    margin-top: 30px;
    border: #F7F7F7 solid 1px;
    width: 1200px;
    height: 100px;
    margin-left: auto;
    margin-right: auto;
    padding-top: 20px;
    border-radius: 5px;
    background: #F7F7F7;
}

.order {
    margin-top: 10px;
    border: #e5e5e5 solid 1px;
    width: 1200px;
    height: 120px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 5px;
}

.bottom {
    display: flex;
    margin-top: 40px;
    width: 1200px;
    height: 120px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 5px;
}

.order_id {
    margin-top: 70px;
    margin-left: auto;
    margin-right: auto;
    width: 1200px;
    height: 15px;

}

.order_id_right {
    color: #B57E66;
    font-size: 14px;
    float: right;
}

.attribute {
    display: flex;
    font-size: 20px;
    background: #F7F7F7;
}

.movie_name,
.movie_time,
.movie_ciname,
.movie_seat {
    width: 300px;
    height: 40px;
    text-align: center;
    padding-top: 10px;
}

.content {
    display: flex;
}

.name,
.time,
.ciname,
.seat {
    padding-top: 30px;
    width: 300px;
    height: 40px;
    text-align: center;
    padding-top: 10px;
}

.ciname {
    font-weight: 800;
}

.time {
    color: #219FFF;
}

.bottom_left {
    width: 780px;
}

.bottom_left1 {

    font-size: 19px;
    font-weight: 800;
}

.bottom_right {
    width: 400px;
    margin-left: 700px;
    font-size: 18px;

    span {
        font-style: italic;
        font-weight: 900;
        font-size: 50px;
        color: red;
    }
}
</style>