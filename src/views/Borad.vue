<template>
    <div class="main">
        <el-tabs v-model="url" @tab-click="handleClick" :stretch="true">
            <el-tab-pane label="票房榜" name="first">
                <div class="model" v-for="(item, index) in pbroadlist.BoxOfficeRankingList" :key="item.id">
                    <div class="model_max">
                        <div class="idx">
                            <img src="'../assets/'+index+'.png'" alt="">
                        </div>
                        <div class="model_mix">
                            <div class="pic">
                                <img src="https://p0.pipi.cn/mmdb/25bfd6020307e150c87e127cdd2e03061fbf6.jpg?imageView2/1/w/464/h/644"
                                    alt="">
                            </div>
                            <div class="tex">
                                <span class="name">{{ item.name }}</span>
                            </div>
                            <div class="sco">
                                <span>{{ item.box_office }}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </el-tab-pane>

            <el-tab-pane label="评分榜" name="second">
                <div class="model" v-for="(item, index) in sbroadlist.ScoreRankingList" :key="item.id">
                    <div class="model_max">
                        <div class="idx">
                            {{ index + 1 }}
                        </div>
                        <div class="model_mix">
                            <div class="pic">
                                <img src="https://p0.pipi.cn/mmdb/25bfd6020307e150c87e127cdd2e03061fbf6.jpg?imageView2/1/w/464/h/644"
                                    alt="">
                            </div>
                            <div class="tex">
                                <span class="name">{{ item.name }}</span>
                            </div>
                            <div class="sco">
                                <span class="sc1">{{ parseInt(9.4)}}{{'.'}}</span>
                                <span class="sc2">{{parseInt(9.4%1*10)}}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </el-tab-pane>

        </el-tabs>
    </div>

</template>

<script>

export default {
    name: 'WorkspaceJsonBorad',
    data() {
        return {
            url: 'first',
            form: {
                Num: 3,
                Page_num: 1
            },
            active1: 'pborad',
            active2: 'sborad',
            pbroadlist: {},
            sbroadlist: {}
        };

    },
    created() {
        this.get_pbroadlist();
        this.get_sbroadlist();
    },

    methods: {
        async get_pbroadlist() {
            const { data: res } = await this.$http.get('GetAllBoxOfficeRankingMovies', { params: this.form })
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }
            this.pbroadlist = res.data

        },
        async get_sbroadlist() {
            const { data: res } = await this.$http.get('GetAllScoreRankingMovies', { params: this.form })
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }
            this.sbroadlist = res.data
        }
    },
};
</script>



<style lang="less" scoped>
.el-container {
    height: 100%;

}

.el-menu {
    width: 100%;
}

.el-main {
    width: 100%;
    padding: 0;
    margin-top: 35px;
}

.main {
    min-width: 1200px;
    max-width: 1200px;
    text-align: center;
    margin: auto;
    margin-top: 25px;
}

.model {
    height: 300px;
    width: 1200px;
}

.model_max {
    margin-left: 300px;
    margin-top: 15px;
    height: 300px;
    width: 1200px;
    display: flex;

    .idx {
        // margin-top: 130px;
        width: 50px;
        padding-top: 100px;

        .inde {
            width: 62px;
            height: 61px;
            background-color: yellow;
            margin-top: 40px;
            padding-bottom: 20px;
        }
    }
}

.ide {
    //margin-top: 20px;
    
    //margin-top: 30px;
}

.model_mix {
    margin-left: 30px;
    height: 300px;
    width: 900px;
    display: flex;
    border: 1px solid rgb(222, 222, 226);
}

.pic {
    height: 300px;

    img {
        height: 300px;
    }
}

.tex {
    margin-left: 30px;
    margin-top: 140px;

    .name {
        font-size: 25px;
    }
}

.sco {
    margin-left: 380px;
    margin-top: 140px;

    .name {
        font-size: 25px;
    }
    .sc1{
        font-style: italic;
        font-weight: 900;
        font-size: 50px;
        color: rgb(241, 220, 137);
    }
    .sc2 {
        font-style: italic;
        font-weight: 700;
        font-size: 25px;
        color: rgb(241, 220, 137);
    }
}
</style>