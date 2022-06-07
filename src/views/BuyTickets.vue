<template>
    <div class="container">
        <div class="banner">
            <div class="flexa">
                <div class="imag"><img class="image"
                        src="https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644"
                        alt="">
                </div>
                <div class="name">
                    <h1>{{ movieInfo.name }}</h1>
                    <span>{{ movieInfo.tag }}&nbsp; &nbsp; &nbsp; 电影</span>
                    <br><br>
                    <span>{{ movieInfo.zone }}/{{ movieInfo.duration }}分钟</span>
                    <br><br>
                    <span>{{ movieInfo.up_Time }} 西安邮电大学上映</span>
                    <br><br>
                    <div class="btns">
                        <el-button type="info" size="small" icon="el-icon-bell"> 想看 &nbsp;&nbsp;&nbsp; </el-button>
                        <el-button type="info" size="small" icon="el-icon-star-off"> 评分&nbsp;&nbsp;&nbsp; </el-button>
                    </div>
                </div>
                <div style="margin-top:120px;margin-left: 80px;">
                    <div class="score" style="color:#fff">西邮评分<h1 style="color:#FFC600;font-size:30px;margin-top:10px">
                            {{ movieInfo.score }}</h1>
                    </div>
                    <div class="pf" style="color:#fff">累计票房<br>
                        <h1 style="color:#F3E7FF;font-size:30px;margin-top:8px">{{ movieInfo.box_office }}亿</h1>
                    </div>
                </div>
            </div>
            <div class="main">
                <div class="schedule">
                    <div class="schedule_btu">
                        <div class="schedule_btu_word"><span>观影日期</span></div>
                        <div class="schedule_btu_itme" v-for="item in dateList" :key="item">
                            <span>{{ item.data_day }}</span>
                        </div>
                    </div>
                    <div class="list">
                        <el-table :data="schedule" stripe style="width: 100%">
                            <el-table-column prop="start_time" label="开始时间" width="180">
                            </el-table-column>
                            <el-table-column prop="end_time" label="结束时间" width="180">
                            </el-table-column>
                            <el-table-column prop="cinema_name" label="放映厅">
                            </el-table-column>
                            <el-table-column prop="price" label="售价(元)">
                            </el-table-column>
                            <el-table-column label="选座购票">
                                <template slot-scope="scope">
                                    <el-button type="primary" round size="mini"
                                        @click="getticketOfSeatList(scope.row.id)">在线选票</el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>

                </div>
            </div>
            <el-dialog title="票信息" :visible.sync="dialogFormVisible" class="cinema_info">

                <el-card class="box-card" v-loading="loading" :element-loading-text="message_get">
                    <div class="cinema">
                        <div class="screen">荧幕</div>
                        <div class="seats">
                            <div v-for="(item, index) in map" :key="item" class="row">
                                <div class="aaa" style="margin-right: 5px">{{ index + 1 }}</div>

                                <div v-for="(ite, ind) in item" :key="ite" style="height: 25px; width:25px"
                                    class="aaa bbb" @click="point(ite.id)">
                                    <div v-if="ite.status === -1" class="ticofseatyes">{{ ind + 1 }}</div>
                                    <div v-if="ite.status === -2" class="ticofseatno">{{ ind + 1 }}</div>
                                    <div v-if="ite.status === -3" class="ticofseatblank"></div>
                                    <div v-if="ite.status === 1" class="ticofseatsaled">{{ ind + 1 }}</div>
                                    <div v-if="ite.status === 2" class="ticofseatselectbymyself">{{ ind + 1 }}</div>
                                </div>
                                <br>
                            </div>
                        </div>
                    </div>
                    <div>
                        <!-- <span>{{ this.tickets.id_list }}</span> -->
                        <div class="palceInfo" v-for="item in this.place" :key="item">
                        <span>{{ item.row }}排</span>
                        <span>{{item.col}}座</span>
                        </div>
                        
                    </div>
                </el-card>

                <div slot="footer" class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取 消</el-button>
                    <el-button type="primary" @click="updateTicket">确定</el-button>
                </div>
            </el-dialog>
        </div>

    </div>
</template>

<script>
import LoginVue from './Login.vue'

export default {
    data() {
        return {
            movie_query: {
                movie_id: ""
            },
            schedule_query: {
                movie_id: "",
                date_day: ""
            },
            tickets_query: {
                id: ""
            },
            tableData: {},
            dateList: {},
            movieInfo: {},
            schedule: {},
            ticketOfSeatList: {},

            map: {},
            tickets: {
                id_list: [],
                user_id: ''
            },
            place: [],
            loading: false,
            dialogFormVisible: false,

        }
    },
    created() {
        this.getMovieInfo()
        this.tickets.user_id = window.sessionStorage.userId
    },
    methods: {
        point(id) {

            for (let i = 0; i < this.map.length; i++) {
                for (let j = 0; j < this.map[0].length; j++) {
                    if (id === this.map[i][j].id) {
                        if (this.map[i][j].status === -1) {
                            this.map[i][j].status = 2
                            this.tickets.id_list.push(this.map[i][j].id)
                            // this.place.row.push(i + 1)
                            // this.place.col.push(j + 1)
                            // this.place.price.push(this.map[i][j].status)
                            var tem={}
                            tem.row=i+1
                            tem.col=j+1
                            tem.price=this.map[i][j].status
                            this.place.push(tem)
                        }
                        else if (this.map[i][j].status === 2) {
                            this.map[i][j].status = -1
                            console.log(this.map[i][j].id);
                            for (let k = 0; k < this.tickets.id_list.length; k++) {
                                if (this.tickets.id_list[k] === this.map[i][j].id) {
                                    this.tickets.id_list.splice(k, 1)
                                    this.place.splice(k, 1)
                                }
                            }
                        }
                    }

                }
            }
        },
        async updateTicket() {
            const { data: res } = await this.$http.put('SaleTicket', this.tickets)
            console.log(res.code)
            if (res.code !== 1000) {
                this.$message.error("购票失败")
                this.$router.push("/home")
                return
            }

            this.$message.success("购票成功")
            this.dialogFormVisible = false
        },
        async getMovieInfo() {
            var id = this.$route.params.id
            const { data: res } = await this.$http.get('GetMovieInfoByID/' + id)
            if (res.code !== 1000) {
                this.$message.error("获取电影详情失败")
                this.$router.push("/home")
                return
            }
            this.movieInfo = res.data.movie
            this.getdateList()
        },
        async getticketOfSeatList(ID) {
            this.tickets_query.id = ID
            const { data: res } = await this.$http.get('GetTicketByScheduleId', { params: this.tickets_query })
            if (res.code !== 1000) {
                this.$message.error("获取票失败")
                this.$router.push("/home")
                return
            }
            this.map = res.data.list
            this.dialogFormVisible = true
        },
        async getdateList() {
            var id = this.$route.params.id
            this.movie_query.movie_id = id

            const { data: res } = await this.$http.get('getallscheduledaybymovieid', { params: this.movie_query })
            if (res.code !== 1000) {
                this.$message.error("获取电影上映日期失败")
                this.$router.push("/home")
                return
            }
            this.dateList = res.data.Time
            this.schedule_query.date_day = this.dateList[0].data_day
            this.schedule_query.movie_id = id
            this.getScheduleList()

        },
        async getScheduleList() {


            const { data: res } = await this.$http.get('getallschedulebymovieidandday', { params: this.schedule_query })


            if (res.code !== 1000) {
                this.$message.error("获取演出计划失败")
                this.$router.push("/home")
                return
            }
            this.schedule = res.data.list

        }
    },
}

</script>

<style lang="less" scoped>
.palceInfo {
    color: #409EFF;
    margin-left: 30px;
}

.container {
    height: 100%;
}

.banner {

    margin: 0;
    width: 100%;
    min-width: 1200px;
    background: #392f59 no-repeat 50%;
    height: 350px;
}

.flexa {
    display: flex;
    padding-top: 45px;
    margin: auto;
    // transform: translate(-50%);
    height: 98%;
    width: 1200px;

}

.buy_btn {
    padding-top: 5px;
    text-align: center;
    height: 30px;
    width: 185px;
    background: #A6B8FF;
    border-radius: 5px;
    cursor: pointer;
}

.buy_btn:hover {
    background: #A0CFFF;
}

.name {
    margin-left: 25px;
    color: #fff;

    h1 {
        margin-top: 0;
        font-size: 30px;
    }
}

.imag {
    height: 325px;
    padding-bottom: 0;
    width: 260px;
    background: #fff;
}

.image {
    margin-top: 3px;
    margin-left: 5px;
    width: 250px;
    height: 320px;
}

.main {
    display: flex;
    margin: auto;
    width: 1200px;
}

.schedule {
    width: 2400px;
}

.schedule_btu {
    color: #000;
    cursor: pointer;
    display: flex;

    .schedule_btu_word {
        padding-top: 2px;
    }

    .schedule_btu_itme {

        text-align: center;
        height: 25px;
        width: 100px;
        margin-left: 20px;
        border-radius: 13px;
        padding-top: 2px;
    }

    .schedule_btu_itme:hover {
        color: #fff;
        background: #409EFF;
    }
}


.cinema {
    border: 1px solid #000;
    height: 400px;
    width: 450px;
    background: blanchedalmond;
    left: 25%;
    top: 20%;
    border-radius: 8px;
    float: left;
}

.seats {
    margin-top: 50px;
    margin-left: 20px;
}

.screen {
    height: 35px;
    width: 150px;
    background: white;
    border-radius: 5px;
    text-align: center;
    margin: auto;
    margin-top: 30px;
}

.ticofseatyes,
.ticofseatno,
.ticofseatblack,
.ticofseatsaled,
.ticofseatselectbymyself {
    height: 25px;
    width: 25px;
    border-radius: 3px;
    margin-top: 3px;
    margin-left: 3px;
    text-align: center;
}

.ticofseatyes {
    background: #fff;
}

.ticofseatno {
    background: red;
}

.ticofseatblack {
    background: blanchedalmond;
}

.ticofseatselectbymyself {
    background: green;
}

.row {
    width: 450px;
    margin-top: 8px;
}

.aaa {
    float: left;
    margin-left: 5px;
}

.bbb {
    cursor: pointer;
}

.ticket {
    border: 1px solid #000;
    height: 500px;
    width: 350px;
    background: blanchedalmond;
    float: right;
    margin-right: 150px;
    margin-top: 80px;
}
</style>