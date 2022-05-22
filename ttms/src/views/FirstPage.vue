<template>
    <el-row>
        <div class="banner">
            <div class="item">
                <img :src="dataList[currentIndex]">
            </div>
            <div class="page" v-if="this.dataList.length > 1">
                <ul>
                    <li @click="gotoPage(prevIndex)">&lt;</li>
                    <li v-for="(item, index) in dataList" @click="gotoPage(index)"
                        :class="{ 'current': currentIndex == index }"> {{ index + 1 }}</li>
                    <li @click="gotoPage(nextIndex)">&gt;</li>
                </ul>
            </div>
        </div>
        <div class="Container">
            <div class="el-col1">
                <div class="main">
                    <span class="panel">
                        <div class="label">
                            <h1>正在热映</h1>
                            <span class="panel-more"><a href="">全部></a></span>

                        </div>
                    </span>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>好吃的汉堡</span>
                                <div class="bottom clearfix">

                                    <el-button type="text" class="button">特惠购买</el-button>
                                </div>
                            </div>
                        </el-card>
                    </div>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>好吃的汉堡</span>
                                <div class="bottom clearfix">

                                    <el-button type="text" class="button">特惠购买</el-button>
                                </div>
                            </div>
                        </el-card>
                    </div>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>好吃的汉堡</span>
                                <div class="bottom clearfix">

                                    <el-button type="text" class="button">特惠购买</el-button>
                                </div>
                            </div>
                        </el-card>
                    </div>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>好吃的汉堡</span>
                                <div class="bottom clearfix">

                                    <el-button type="text" class="button">特惠购买</el-button>
                                </div>
                            </div>
                        </el-card>
                    </div>
                </div>
            </div>
            <div class="el-col2">
                <div class="aside">
                    <div class="label">
                        <h1>票房最佳</h1>
                        <span class="panel-more"><a href="">全部></a></span>
                    </div>
                    <div>
                        <div class="borde_top">
                            <img src="https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644" alt="" class="bored_top_img">
                            <div>
                                <span class="borde_top_name">{{borad_top.name}}</span><br>
                                <span class="borde_top_pf">{{borad_top.pf}}万</span>
                            </div>
                        </div>
                        <div v-for="(item, index) in borad_list" :key="index" class="bored_item">
                            <span class="bored_item_span">
                                <i>{{index+2}}</i>&nbsp;&nbsp;&nbsp;
                                <span class="bored_item_name">{{item.name}}</span>
                                <span class="bored_item_pf">{{item.pf}}万</span>
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <el-drawer title="我是标题" :visible.sync="movie_info_drawer" :with-header="false" size="80%">
            <div>
                <div class="drawer_picture">
                    <el-image style="width: 432px; height: 600px" :src="url" :fit="fill"></el-image>
                </div>
            </div>
        </el-drawer>
    </el-row>
</template>
<script>
export default {
    data() {
        return {
            movie_info_drawer: false,
            url: 'https://p0.pipi.cn/mmdb/25bfd69ae7a5bff0ee230fed0210646bfdde9.jpg?imageView2/1/w/160/h/220',
            dataList: ["https://i1.mifile.cn/a4/xmad_15535933141925_ulkYv.jpg", "https://i1.mifile.cn/a4/xmad_15532384207972_iJXSx.jpg", "https://i1.mifile.cn/a4/xmad_15517939170939_oiXCK.jpg"],
            currentIndex: 0,   //默认显示图片
            timer: null,   //定时器,
            borad_list: [{ "name": "我和我的祖国", "pf": 123 }, { "name": "喜羊羊与灰太狼", "pf": 122 }, { "name": "海绵宝宝", "pf": 121 }, { "name": "桃花侠大战菊花怪", "pf": 120 }],
            borad_top:{
                img:'',
                name:'坏蛋联盟',
                pf:'1271.16'
            }
        }
    },
    methods: {
        gotoPage(index) {
            this.currentIndex = index;
        },
        //定时器
        runInv() {
            this.timer = setInterval(() => {
                this.gotoPage(this.nextIndex)
            }, 1000)
        }
    },
    computed: {
        //上一张
        prevIndex() {
            if (this.currentIndex == 0) {
                return this.dataList.length - 1;
            } else {
                return this.currentIndex - 1;
            }
        },
        //下一张
        nextIndex() {
            if (this.currentIndex == this.dataList.length - 1) {
                return 0;
            } else {
                return this.currentIndex + 1;
            }
        }
    }
}
</script>

<style lang="less" scoped>
.main {
    background: #fff;
    width: 100%;
    float: left;
    height: 100%;
}

.borde_top {
    margin-top: 20px;
    margin-left: 20px;
    border: #F7F7F7 solid 3px;
    height: 80px;
    display: flex;
    cursor: pointer;
}
.borde_top:hover{
    background: #F7F7F7;
}
.bored_top_img{
    height: 80px;
    width: 120px;
}

.borde_top_name {
    font-size: 24px;
}
.borde_top_pf {
    color: #EF4238;
    font-size: 10px;
}
.bored_item {
    width: 100%;
    height: 55px;
}
.bored_item:hover{
    cursor: pointer;
    background: #F7F7F7;
}

.bored_item_pf{
    float: right;
    color: #EF4238;
    font-size: 10px;
}


i {
    font-family: 'Courier New', Courier, monospace;
    font-size: 25px;
    margin-left: 20px;
    color: #EF4238;
}

.bored_item_index_name {
    margin-left: 15px;
    width: 200px;
}

.aside {
    height: 100%;
    background: #fff;
    width: 100%;
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

.banner {
    height: 500px;
    width: 1200px;
    margin: 0 auto;
    position: relative;
    margin-top: 0;
}

.banner img {
    margin: auto;
    width: 100%;
    display: block;
    height: 500px;
}

.banner .page {
    margin: auto;
    background: rgba(0, 0, 0, .5);
    position: absolute;
    right: 0;
    bottom: 0;
    width: 100%;
}

.banner .page ul {
    float: right;
}

.current {
    color: #ff6700;
}
</style>