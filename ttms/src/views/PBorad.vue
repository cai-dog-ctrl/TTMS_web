<template>
    <div>
        <div class="model" v-for="(item, index) in scoreRankingList" :key="item.id">
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
                        <span>{{ item.name }}</span>
                    </div>
                    <div class="sco">
                        <span>{{ item.score }}</span>
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
            form:{
                Num:'2',
                Page_num:'2'
            },
            scoreRankingList: [
                {
                    id: '',
                    name: '',
                    score: '',
                    cover_img_path: ''
                }
            ]
        }
    },
    created() {
        this.get_pBorad()
    },
    methods: {
        async get_pBorad() {
            const { data: res } = await this.$http.get('GetAllScoreRankingMovies', { params: this.form })
            console.log(res);
            if (res.code !== 1000) {
                this.$message.error("获取信息失败")
                return
            }
            this.scoreRankingList = res.data.ScoreRankingList
        }
    }
}

</script>

<style lang="less" scoped>
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
        margin-top: 140px;
    }
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
}

.sco {
    margin-left: 380px;
    margin-top: 140px;
}
</style>