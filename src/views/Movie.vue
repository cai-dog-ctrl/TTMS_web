<template>
    <div class="contain">
        <div class="classify">
            <div class="category">
                <div class="category_tr">
                    <span>类型</span>
                    <div class="borde">
                        <div class="category_item" v-for="item in SortInfo.tag" :key="item">
                            <div class="category_item_info"><span>{{ item }}</span></div>
                        </div>
                    </div>
                </div>

            </div>
            <hr>
            <div class="zone">
                <div class="zone_tr">
                    <span>区域</span>
                    <div class="borde">
                        <div class="zone_item" v-for="item in SortInfo.zone" :key="item">
                            <div class="zone_item_info"><span>{{ item }}</span></div>
                        </div>
                    </div>
                </div>
            </div>
            <hr>
            <div class="sort">
                <div class="sort_tr">
                    <div class="borde">
                        <el-radio-group v-model="radio" class="sort_item" v-for="(item, index) in SortInfo.order"
                            :key="item">
                            <el-radio :label="index">{{ item }}</el-radio>
                        </el-radio-group>
                    </div>
                </div>
            </div>
        </div>
        <div class="card">
            <div v-for="item in MovieInfo.movieList" :key="item.id" class="movie">
                <el-card :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                        class="image" @click="gotoinfo(item.id)">
                    <div style="padding: 14px;">
                        <span>{{ item.name }}</span>
                        <span>{{ item.score }}</span>
                        <span>{{item.id}}</span>
                        <div class="bottom clearfix">

                            <el-button type="text" class="button" @click="gotoinfo(item.id)">特惠购买</el-button>
                        </div>
                    </div>
                </el-card>
            </div>
        </div>


    </div>
</template>

<script>

export default {
    name: 'TTMSWebMovie',
    data() {
        return {
            SortInfo: {},
            MovieInfo: {},
            radio: 3,
            form:{
                Num: 8,
                Page_num: 1
            }
        };
    },

    created() {
        this.getSortInfo();
        this.getMovieInfo();
    },

    methods: {
        async getSortInfo() {
            const { data: res } = await this.$http.get('GetMovieSort')
            if (res.code !== 1000) {
                this.$message.error("获取排序信息失败");
                return
            }
            this.SortInfo = res.data;
        },
        async getMovieInfo() {
            const { data: res2 } = await this.$http.get('GetAllMovies',{ params: this.form })
            if (res2.code !== 1000) {
                this.$message.error("获取电影失败");
                return
            }
            this.MovieInfo = res2.data;
        },
        gotoinfo(id) {
            this.$router.push('/movie/' + id);
        }
    },
};
</script>

<style lang="less" scoped>
.contain {
    margin-top: 30px;
}

.classify {
    width: 1300px;
    border: 1px solid #000;
    margin: auto;
}

.category {
    height: 100px;
}

.category_tr {
    width: 40px;
}

.borde {
    display: flex;

    width: 100%;
    align-items: center;
    /*垂直居中*/
    justify-content: center;
    /*水平居中*/
    // // justify-content: flex-start;
    // margin-left: 20px;
    // width: 100%;
    // // flex-direction: row;
    // flex-wrap: wrap;
    // flex-direction: column;
    // border: 1px solid #cccc;
    // display: flex;
    // flex-direction: row;
    // justify-content: space-evenly;
    // align-items: center;
    // word-wrap:break-word;
    // word-break:break-all;
}

.category_item_info {
    margin-left: 10px;
    width: 64px;

}

.card {
    width: 1400px;
    margin: auto;
    padding-left: 125px;
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
</style>