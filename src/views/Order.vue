<template>
    <div class="main_order">
        <div class="aside">
            <div class="Head">
                <div>
                    <h1>个人信息</h1>
                </div>

            </div>
            <div class="BODY">
                <div class="Info_user" @click="goto_user()">
                    <div><span>基本信息</span></div>
                </div>
                <div class="Info_order">
                    <div><span>我的订单</span></div>
                </div>
            </div>
        </div>
        <div class="order">
            <div class="order_head">
                <h1>我的订单</h1>
            </div>
            <el-divider><i class="el-icon-mobile-phone"></i></el-divider>
            <div class="order_items" v-for="item in this.orderInfo.OrderFrontList" :key="item.index">
                <div class="order_items_head">
                    <div class="order_items_time">
                        <span>{{ item.date }}</span>

                    </div>
                    <div class="order_items_id">
                        <span>西邮订单号:{{ item.order_id }}</span>
                    </div>
                    <div class="order_status">
                        <div class="no_pay">(未支付)</div>
                    </div>
                </div>
                <div class="order_items_body">
                    <div class="imag">
                        <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                            alt="" class="image">
                    </div>
                    <div class="descirption">
                        <div class="movie_name">
                            《{{ item.movie_name }}》
                        </div>
                        <div class="cinema">
                            &nbsp; 3号ALPD激光{{ item.cinema_name }}
                        </div>

                    </div>
                    <div class="price">
                        ￥{{ item.price }}
                    </div>
                    <div class="detail_information">
                        <a href="#" @click="goto_order_info(id)">查看详情</a>
                    </div>
                </div>
            </div>
            <div class="page">
                <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                    :current-page="form.Page_num" :page-sizes="[2, 1]" :page-size="form.Num"
                    layout="total,sizes, prev, pager, next, jumper" :total="orderInfo.Total">
                </el-pagination>
            </div>

        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            form: {
                ID: '',
                Num: '2',
                Page_num: '1'
            },
            orderInfo: {}
        }
    },
    created() {
        this.getOrderByUserId()
    },
    methods: {
        handleSizeChange(val) {
            console.log(`每页 ${val} 条`);
            this.form.Num = val
            this.getOrderByUserId()
        },
        handleCurrentChange(val) {
            console.log(`当前页: ${val}`);
            this.form.Page_num = val
            this.getOrderByUserId()
        },
        goto_user() {
            this.$router.push('/user')
        },
        goto_order_info(id) {
            this.$router.push('/orderInfo/' + id)
        },
        async getOrderByUserId() {
            this.form.ID = window.sessionStorage.userid
            if (this.form.ID === '') {
                this.$message.error("请登录后使用")
                return
            }
            const { data: res } = await this.$http.get('GetOrderByUserID/', { params: this.form })
            if (res.code !== 1000) {
                this.$message.error("获取订单信息失败")
                this.$router.push("/home")
                return
            }
            this.orderInfo = res.data

        }
    }
}
</script>

<style lang="less" scoped>
.main_order {
    height: 600px;
    width: 1100px;
    border: #E5E5E5 solid 1px;

    margin-top: 50px;
    margin-left: auto;
    margin-right: auto;
    display: flex;
    border-radius: 5px;
}

.aside {
    text-align: center;
    background: #F4F3F4;
    width: 250px;
    height: 100%;
    border: #E5E5E5 solid 1px;
    margin-left: 0;
}

.Head {

    width: 100%;
    height: 60px;


    div {

        color: #219FFF;
        font-size: 20px;
    }
}

.Info_user,
.Info_order {
    width: 100%;
    height: 30px;

    font-size: 19px;
    padding-top: 10px;
    cursor: pointer;
}

.Info_order {
    background: #219FFF;
    color: #fff;
}

.Info_user {
    color: #000;
}

.order {
    width: 850px;

    .order_head {
        color: #219FFF;
    }
}

.order_head {
    text-align: center;
}

.order_items {
    width: 90%;
    margin-left: auto;
    margin-right: auto;
    border: #E5E5E5 solid 1px;
    height: 180px;
    border-radius: 3px;
}

.order_items_head {
    font-size: 14px;
    display: flex;
    height: 35px;
    background: #F7F7F7;
    padding-top: 10px;
}

.order_items_time {
    width: 100px;
    margin-left: 25px;
}

.order_items_id {
    color: #999;
    width: 250px;
    margin-left: 20px;
}

.imag {
    height: 100%;
    margin-left: 20px;
}

.image {
    margin-top: 5px;
    height: 120px;
    width: 90px;
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

.order_items_body {
    display: flex;
}

.cinema {
    font-size: 10px;
    color: #999;
}

.order_status {
    margin-left: 250px;
    width: 80px;
    font-size: 14px;
}

.no_pay {
    color: #E42D23;
}

.price {
    padding-top: 50px;
    width: 50px;
}

.detail_information {
    padding-top: 50px;
    margin-left: 210px;

    a {
        color: #000;

    }
}

.page {
    margin-top: 20px;
    text-align: center;
    margin: auto;
}
</style>