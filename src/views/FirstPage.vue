<template>
    <el-row>
        <el-carousel :interval="4000" type="card" height="260px">
            <el-carousel-item v-for="item in slideData" :key="item">
                <h3 class="medium"><img :src="item.img" alt=""></h3>
            </el-carousel-item>
        </el-carousel>

        <div class="Container">
            <div class="el-col1">
                <div class="main">
                    <span class="panel">
                        <div class="label">
                            <h1>正在热映</h1>
                            <span class="panel-more"><a href="">全部></a></span>
                        </div>
                    </span>
                    <div v-for="item in first_page_info.ShowingList.ShowingList" :key="item.id" class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="gotoinfo(item.id)">
                            <div style="padding: 14px;">
                                <span>{{ item.name }}</span>
                                <div class="bottom clearfix">

                                    <el-button type="text" class="button" @click="gotoinfo(item.id)">特惠购买</el-button>
                                </div>
                            </div>
                        </el-card>
                    </div>

                </div>
                <div class="main">
                    <span class="panel">
                        <div class="label">
                            <h1>即将上映</h1>
                            <span class="panel-more"><a href="">全部></a></span>

                        </div>
                    </span>
                    <div v-for="item in first_page_info.ComingList.ComingList" :key="item.id" class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="gotoinfo(item.id)">
                            <div style="padding: 14px;">
                                <span>{{ item.name }} |</span><span> 预售</span>
                            </div>
                        </el-card>
                    </div>

                </div>
            </div>
            <div class="el-col2">
                <div class="aside">
                    <div class="label">
                        <h1>票房最佳</h1>
                        <span class="panel-more"><a href="http://localhost:8080/#/pfborad">全部></a></span>
                    </div>
                    <div class="borde">
                        <div class="borde_top"
                            @click="gotoinfo(first_page_info.BoxOfficeRankingList.BoxOfficeRankingList[0].id)">
                            <div><img
                                    src="https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644"
                                    alt="" class="bored_top_img"></div>

                            <div style="padding-top:20px;margin-left:20px;width:60px">
                                <a class="borde_top_name">{{
                                        first_page_info.BoxOfficeRankingList.BoxOfficeRankingList[0].name
                                }}</a><br>

                            </div>
                            <div style="padding-top: 25px;margin-left: 144px;"><span class="borde_top_pf">{{
                                    first_page_info.BoxOfficeRankingList.BoxOfficeRankingList[0].box_office
                            }}万</span></div>
                        </div>
                        <div v-for="(item, index) in first_page_info.BoxOfficeRankingList.BoxOfficeRankingList.slice(1, first_page_info.BoxOfficeRankingList.BoxOfficeRankingList.length)"
                            :key="index" class="bored_item" @click="gotoinfo(item.id)">
                            <span class="bored_item_span">
                                <i>{{ index + 2 }}</i>&nbsp;&nbsp;&nbsp;
                                <span class="bored_item_name">{{ item.name }}</span>
                                <span class="bored_item_pf">{{ item.box_office }}万</span>
                            </span>
                        </div>
                    </div>
                </div>
                <div class="aside">
                    <div class="label">
                        <h1>TOP</h1>
                        <span class="panel-more"><a href="http://localhost:8080/#/sborad">全部></a></span>
                    </div>
                    <div class="bored" style="margin-top:15px">
                        <div class="borde_top"
                            @click="gotoinfo(first_page_info.ScoreRankingList.ScoreRankingList[0].id)">
                            <div><img :src="src_url" class="bored_top_img"
                                    @click="gotoinfo(first_page_info.ScoreRankingList.ScoreRankingList[0].id)"></div>

                            <div style="padding-top:20px;margin-left:20px;width:60px">
                                <a class="borde_top_name">{{
                                        first_page_info.ScoreRankingList.ScoreRankingList[0].name
                                }}</a>

                            </div>
                            <div style="margin-left:144px;padding-top: 25px;"><span class="borde_top_pf">{{
                                    first_page_info.ScoreRankingList.ScoreRankingList[0].score
                            }}分</span></div>
                        </div>
                    </div>
                    <div v-for="(item, index) in first_page_info.ScoreRankingList.ScoreRankingList.slice(1, first_page_info.ScoreRankingList.ScoreRankingList.length)"
                        :key="index" class="bored_item" @click="gotoinfo(item.id)">
                        <span class="bored_item_span">
                            <i>{{ index + 2 }}</i>&nbsp;&nbsp;&nbsp;
                            <span class="bored_item_name">{{ item.name }}</span>
                            <span class="bored_item_pf">{{ item.score }}分</span>
                        </span>
                    </div>
                </div>
            </div>
        </div>


       
    </el-row>
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
            movie_info_drawer: false,
            url: 'https://p0.pipi.cn/mmdb/25bfd69ae7a5bff0ee230fed0210646bfdde9.jpg?imageView2/1/w/160/h/220',
            borad_list: [{ "name": "我和我的祖国", "pf": 123 }, { "name": "喜羊羊与灰太狼", "pf": 122 }, { "name": "海绵宝宝", "pf": 121 }, { "name": "桃花侠大战菊花怪", "pf": 120 }],
            borad_top: {
                img: '',
                name: '坏蛋联盟',
                pf: '1271.16'
            },
            slideData: [
                {
                    id: 1,
                    img: 'https://img.alicdn.com/imgextra/i2/6000000001117/O1CN01Mn6ES81K7d5SAd6hU_!!6000000001117-0-octopus.jpg'
                },
                {
                    id: 2,
                    img: 'https://aecpm.alicdn.com/simba/img/TB1XotJXQfb_uJkSnhJSuvdDVXa.jpg'
                },
                {
                    id: 3,
                    img: 'https://aecpm.alicdn.com/simba/img/TB1JNHwKFXXXXafXVXXSutbFXXX.jpg'
                }
            ],
            currentIndex: 0
        }
    },
    created() {
        this.autoPlay()
        this.get_firstPage()
    },
    methods: {
        // 自动轮播，每隔 1 秒轮播一次
        autoPlay() {
            setInterval(() => {
                this._setIndex()
            }, 4000)
        },
        // 设置当前索引
        _setIndex() {
            this.currentIndex++
            if (this.currentIndex === this.slideData.length) this.currentIndex = 0
        },
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
        openRate(id){

        }

    }



}
</script>

<style lang="less" scoped>
.el-carousel {
    left: 50%;
    transform: translate(-50%);
    width: 1400px;
}

.el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
}

.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
}

.borde {
    margin-top: 20px;
}

.main {
    margin-bottom: 40px;
    background: #fff;
    width: 750px;
    float: left;
    height: 100%;
}

.borde_top {
    //margin-top: 20px;
    margin-left: 20px;
    border: #F7F7F7 solid 1px;
    height: 80px;
    display: flex;
    padding-right: 0;
    cursor: pointer;
}

.borde_top:hover {
    background: #F7F7F7;
}

.bored_top_img {
    height: 80px;
    width: 120px;
}

.borde_top_name {
    font-size: 20px;
    float: inline-start;
    margin-top: 20px;
}

.borde_top_pf {
    float: right;
    color: #EF4238;
    font-size: 10px;
    width: 45px;
}

.bored_item {
    margin-left: 20px;
    padding-top: 25px;
    width: 380px;
    height: 50px;
}

.bored_item:hover {
    cursor: pointer;
    background: #F7F7F7;
}

.bored_item_pf {
    float: right;
    color: #EF4238;
    font-size: 10px;
}


i {
    font-family: 'Courier New', Courier, monospace;
    font-size: 25px;
    margin-left: 0;
    color: #EF4238;
}

.bored_item_index_name {
    margin-left: 15px;
    width: 200px;
}

.aside {
    margin-bottom: 278px;
    height: 508px;
    background: #fff;
    width: 400px;
    float: left;
}

.el-col1 {
    width: 750px;
    height: 100%;

    margin-top: 35px;
}

.el-col2 {

    width: 400px;
    height: 100%;
    margin-top: 35px;
}

.el-row {

    margin: 0 auto;
}

.panel-more {
    margin-top: 18px;

    a {
        color: #219FFF;
    }
}

.label {
    display: flex;
    //align-items: center;
    justify-content: space-between;

}

h1 {

    font-family: 'Courier New', Courier, monospace;
    background-color: #fff;
    font-size: 35px;
    margin-top: 0;
    color: #219FFF;
}

.Container {
    //text-align: center;
    margin: auto;
    display: flex;
    justify-content: space-around;
    width: 1200px;

    h1 {
        margin-left: 18px;
    }
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

.panel {
    margin-bottom: 0;
}

* {
    margin: 0;
    padding: 0;
}

ul li {
    list-style: none;
    float: left;
    width: 30px;
    height: 40px;
    line-height: 40px;
    text-align: center;
    cursor: pointer;
    color: rgba(255, 255, 255, .8);
    font-size: 14px;
}
</style>