<template>
    <div class="contain">
        <div class="classify">
            <div class="category">
                <div class="category_tr">
                    <span>类型:</span>
                </div>
                <div @click="getcategoryIndex(index), categoryForm(item)"
                    :class="btnActive.categoryIndex === index ? 'category_borde_focus' : 'category_borde'"
                    v-for="(item, index) in SortInfo.tag" :key="index">
                    <div>
                        <span>{{ item }}</span>
                    </div>
                </div>
            </div>
            <el-divider></el-divider>
            <div class="zone">
                <div class="zone_tr">
                    <span>区域:</span>
                </div>
                <div @click="getzoneIndex(index), zoneForm(item)"
                    :class="btnActive.zoneIndex === index ? 'zone_borde_focus' : 'zone_borde'"
                    v-for="(item, index) in SortInfo.zone" :key="index">
                    <div>
                        <span>{{ item }}</span>
                    </div>
                </div>
            </div>
            <el-divider></el-divider>
            <div class="sort">
                <el-radio-group v-model="radio" class="sort_item" @change="sortForm">
                    <el-radio :label="1">按热门排序</el-radio>
                    <el-radio :label="2">按时间排序</el-radio>
                    <el-radio :label="3">按评分排序</el-radio>
                </el-radio-group>
            </div>
        </div>
        <div class="card">
            <div v-for="item in MovieInfo.movieList" :key="item.id" class="movie">
                <el-card :body-style="{ padding: '0px' }" shadow="hover">
                    <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                        class="image" @click="gotoinfo(item.id)">
                    <div style="padding: 14px;">
                        <div class="text">
                            <div class="text_name"><span>{{ item.name }}</span></div>
                            <div class="text_score"><span>{{ item.score }}</span></div>
                        </div>
                        <div class="bottom clearfix">

                            <el-button type="text" class="button" @click="buyTickets(item.id)">特惠购买</el-button>
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
            radio: 1,
            btnActive: {
                categoryIndex: 0,
                zoneIndex: 0
            },
            form: {
                Num: 8,
                Page_num: 1,
                FlagOfType: '',
                FlagOfZone: '',
                FlagOfOrder: 'boxoffice'
            }
        };
    },

    created() {
        this.getSortInfo();
        this.getMovieInfo();
    },
    methods: {
        getcategoryIndex(index) {
            this.btnActive.categoryIndex = index
        },
        getzoneIndex(index) {
            this.btnActive.zoneIndex = index
        },
        categoryForm(item) {
            this.form.FlagOfType = item
            this.getMovieInfo();
        },
        zoneForm(item) {
            this.form.FlagOfZone = item
            this.getMovieInfo();
        },
        sortForm() {
            if (this.radio === 2) {
                this.form.FlagOfOrder = 'date';
            }
            if (this.radio === 3) {
                this.form.FlagOfOrder = 'score';
            }
            this.getMovieInfo();
        },
        async getSortInfo() {
            const { data: res } = await this.$http.get('GetMovieSort')
            if (res.code !== 1000) {
                this.$message.error("获取排序信息失败");
                return
            }
            this.SortInfo = res.data;
        },
        async getMovieInfo() {
            const { data: res2 } = await this.$http.get('GetAllMovies', { params: this.form })
            if (res2.code !== 1000) {
                this.$message.error("获取电影失败");
                return
            }
            this.MovieInfo = res2.data;
        },
        gotoinfo(id) {
            this.$router.push('/movie/' + id);
        },
        sort(index) {
            if (index === 0) {
                this.form.FlagOfOrder = 'boxoffice'
            } else if (index === 1) {
                this.form.FlagOfOrder = 'date'
            } else {
                this.form.FlagOfOrder = 'score'
            }
            this.getMovieInfo();
        },
        buyTickets(id) {
            this.$router.push('/buytickets/' + id);
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
    border: #dedcdc solid 1px;
    margin: auto;
    border-radius: 5px;
}


.category {
    height: 36px;
    width: 1300px;
    margin-left: 10px;
}

.zone {
    height: 3px;
    width: 1300px;
    margin-left: 10px;
}

.category_tr,
.zone_tr {
    float: left;
    margin-top: 5px;
}

.category_borde {
    padding-left: 5px;
    padding-right: 5px;
    margin-top: 5px;
    margin-bottom: 5px;
    float: left;
    margin-left: 14px;
    cursor: pointer;
    // border: 1px solid #000;
}

.category_borde:hover {
    color: blue;
}

.category_borde_focus {
    color: #fff;
    background-color: blue;
    border-radius: 10px;
    padding-left: 5px;
    padding-right: 5px;
    margin-top: 5px;
    margin-bottom: 5px;
    float: left;
    margin-left: 14px;
    // border: 1px solid #4FA9FD;
}


.zone_borde {
    padding-left: 5px;
    padding-right: 5px;
    margin-top: 5px;
    margin-bottom: 5px;
    float: left;
    margin-left: 14px;
    cursor: pointer;
    // border: 1px solid #000;
}

.zone_borde:hover {
    color: blue;
}

.zone_borde_focus {
    color: #fff;
    background-color: blue;
    border-radius: 10px;
    padding-left: 5px;
    padding-right: 5px;
    margin-top: 5px;
    margin-bottom: 5px;
    float: left;
    margin-left: 14px;
    // border: 1px solid #4FA9FD;
}

.el-divider--horizontal {
    display: block;
    width: 100%;
    height: 1px;
    margin: 23px 0 6px 0;
}


.sort_item {
    margin-left: 15px;
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

.text {
    display: flex;
    justify-content: space-between;
}

.text_score {
    font-style: italic;
    color: rgb(241, 220, 137);
}

.image {

    cursor: pointer;
    position: relative;
}
</style>