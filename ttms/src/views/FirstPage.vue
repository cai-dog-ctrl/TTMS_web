<template>
    <el-row>
        <el-col class="el-col1">
            <div class="main">
                <span class="panel">
                    <h1>正在热映</h1>
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
        </el-col>
        <el-col class="el-col2">
            <div class="aside">
                <div class="panel-title">
                    <h1>票房最佳</h1>
                </div>
            </div>
        </el-col>
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
            url:'https://p0.pipi.cn/mmdb/25bfd69ae7a5bff0ee230fed0210646bfdde9.jpg?imageView2/1/w/160/h/220'
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

.aside {
    height: 100%;
    background: #fff;
    width: 100%;
    float: left;
}

.el-col1 {
    width: 750px;
    height: 100%;
    margin-left: 10%;
}

.el-col2 {
    margin-left: 85px;
    width: 400px;
    height: 100%;
}

.el-row {
    height: 2000px;
    margin-top: 50px
}

h1 {
    margin-left: 18px;
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    font-size: 30px;
    margin-top: 0;
    color: #219FFF;
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
</style>