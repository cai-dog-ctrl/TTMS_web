<template>
    <el-row>
        <div id="appaaa">
            <ul>
                <li v-for="(item, index) in slideData" :key="index">
                    <transition>
                        <img class="imga" :src="item.img" v-if="index === currentIndex" />
                    </transition>
                </li>
            </ul>
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
                <div class="main">
                    <span class="panel">
                        <div class="label">
                            <h1>即将上映</h1>
                            <span class="panel-more"><a href="">全部></a></span>

                        </div>
                    </span>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>预告片 |</span><span> 预售</span>
                            </div>
                        </el-card>
                    </div>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>预告片 |</span><span> 预售</span>
                            </div>
                        </el-card>
                    </div>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>预告片 |</span><span> 预售</span>
                            </div>
                        </el-card>
                    </div>
                    <div class="movie">
                        <el-card :body-style="{ padding: '0px' }" shadow="hover">
                            <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                                class="image" @click="movie_info_drawer = true">
                            <div style="padding: 14px;">
                                <span>预告片 |</span><span> 预售</span>
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
                            <img src="https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644"
                                alt="" class="bored_top_img">
                            <div>
                                <span class="borde_top_name">{{ borad_top.name }}</span><br>
                                <span class="borde_top_pf">{{ borad_top.pf }}万</span>
                            </div>
                        </div>
                        <div v-for="(item, index) in borad_list" :key="index" class="bored_item">
                            <span class="bored_item_span">
                                <i>{{ index + 2 }}</i>&nbsp;&nbsp;&nbsp;
                                <span class="bored_item_name">{{ item.name }}</span>
                                <span class="bored_item_pf">{{ item.pf }}万</span>
                            </span>
                        </div>
                    </div>
                </div>
                 <div class="aside">
                    <div class="label">
                        <h1>票房最佳</h1>
                        <span class="panel-more"><a href="">全部></a></span>
                    </div>
                    <div>
                        <div class="borde_top">
                            <img src="https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644"
                                alt="" class="bored_top_img">
                            <div>
                                <span class="borde_top_name">{{ borad_top.name }}</span><br>
                                <span class="borde_top_pf">{{ borad_top.pf }}万</span>
                            </div>
                        </div>
                        <div v-for="(item, index) in borad_list" :key="index" class="bored_item">
                            <span class="bored_item_span">
                                <i>{{ index + 2 }}</i>&nbsp;&nbsp;&nbsp;
                                <span class="bored_item_name">{{ item.name }}</span>
                                <span class="bored_item_pf">{{ item.pf }}万</span>
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
        }
    }



}
</script>

<style lang="less" scoped>
.imga {
    width: 1200px;
}

#appaaa {

    >ul {
        left: 50%;
        transform: translate(-50%);
        position: relative;
        text-align: center;
        width: 1200px;
        height: 284px;
        overflow: hidden;

        img {
            position: absolute;
            top: 0;
            left: 0;
        }

        .v-enter-active,
        .v-leave-active {
            transition: transform 0.5s;
        }

        .v-enter {
            transform: translateX(100%);
        }

        .v-enter-to {
            transform: translateX(0);
        }

        .v-leave {
            transform: translateX(0);
        }

        .v-leave-to {
            transform: translateX(-100%);
        }
    }
}

.main {
    margin-bottom: 40px;
    background: #fff;
    width: 750px;
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

.borde_top:hover {
    background: #F7F7F7;
}

.bored_top_img {
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
    margin-top: 40px;
    width: 100%;
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
    margin-left: 20px;
    color: #EF4238;
}

.bored_item_index_name {
    margin-left: 15px;
    width: 200px;
}

.aside {
    margin-bottom: 278px;
    height: 100%;
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