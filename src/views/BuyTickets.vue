<template>
    <div class="container">
        <div class="banner">
            <div class="flexa">
                <div class="imag"><img class="image"
                        src="https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644"
                        alt="">
                </div>
                <div class="name">
                    <h1>坏蛋联盟</h1>
                    <span>喜剧&nbsp; &nbsp; &nbsp; 电影</span>
                    <br><br>
                    <span>中国/100分钟</span>
                    <br><br>
                    <span>2022-4-29 西安邮电大学上映</span>
                    <br><br>
                    <div class="btns">
                        <el-button type="info" size="small" icon="el-icon-bell"> 想看 &nbsp;&nbsp;&nbsp; </el-button>
                        <el-button type="info" size="small" icon="el-icon-star-off"> 评分&nbsp;&nbsp;&nbsp; </el-button>
                    </div>
                </div>
                <div style="margin-top:120px;margin-left: 80px;">
                    <div class="score" style="color:#fff">西邮评分<h1 style="color:#FFC600;font-size:30px;margin-top:10px">
                            9.1</h1>
                    </div>
                    <div class="pf" style="color:#fff">累计票房<br>
                        <h1 style="color:#F3E7FF;font-size:30px;margin-top:8px">26.3亿</h1>
                    </div>
                </div>
            </div>
            <div class="main">
                <div class="schedule">
                    <div class="schedule_btu">
                        <div class="schedule_btu_word"><span>观影日期</span></div>
                        <div class="schedule_btu_itme" v-for="item in dateList" :key="item">
                            <span>{{ item }}</span>
                        </div>
                    </div>
                    <div class="list">
                        <el-table :data="tableData" stripe style="width: 100%">
                            <el-table-column prop="date" label="时间" width="180">
                            </el-table-column>
                            <el-table-column prop="type" label="类型" width="180">
                            </el-table-column>
                            <el-table-column prop="address" label="放映厅">
                            </el-table-column>
                            <el-table-column prop="price" label="售价(元)">
                            </el-table-column>
                            <el-table-column label="选座购票">
                                <el-button type="primary" round size="mini">在线选票</el-button>
                            </el-table-column>
                        </el-table>
                    </div>

                </div>
            </div>
        </div>

    </div>
</template>

<script>
export default {
    data() {
        return {
            tableData: [{
                date: '2016-05-02',
                type: '3D',
                address: '1号放映厅',
                price: '30'
            }, {
                date: '2016-05-04',
                type: '3D',
                address: '2号放映厅',
                price: '35'
            }, {
                date: '2016-05-01',
                type: '3D',
                address: '3号放映厅',
                price: '50'
            }, {
                date: '2016-05-03',
                type: '3D',
                address: '4号放映厅',
                price: '40'
            }],
            dateList: ['2022-10-09', '2022-10-23']
        }
    },
    methods: {
        async GetDateList() {
            var id = this.$route.params.id
            const { data: res } = await this.$http.get('getalldcheduledaybymovieid/' + id)
            if (res.code !== 1000) {
                this.$message.error("获取演出计划失败")
                return
            }
            
        }
    }
}
</script>

<style lang="less" scoped>
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
</style>