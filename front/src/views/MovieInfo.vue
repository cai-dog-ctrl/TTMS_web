<template>
  <div class="container">
    <div class="banner">
      <div class="flexa">
        <div class="imag"><img class="image"
            :src="'http://127.0.0.1:8080/api/getpicturebyfilename/' + movieInfo.cover_img_path" alt="">
        </div>
        <div class="name">
          <h1>{{ movieInfo.name }}</h1>
          <span>{{ movieInfo.tag }}&nbsp; &nbsp; &nbsp; 电影</span>
          <br><br>
          <span>中国/{{ movieInfo.duration }}分钟</span>
          <br><br>
          <span>{{ movieInfo.up_Time }} 西安邮电大学上映</span>
          <br><br>
          <div class="btns">
            <el-button type="info" size="small" icon="el-icon-bell"> 想看 &nbsp;&nbsp;&nbsp; </el-button>
            <el-button type="info" size="small" icon="el-icon-star-off" @click="openRate(movieInfo.id)">
              评分&nbsp;&nbsp;&nbsp; </el-button>
          </div>
          <br>
          <div class="buy_btn" @click="buyTickets">特惠购票</div>
        </div>
        <div style="margin-top:120px;margin-left: 80px;">
          <div class="score" style="color:#fff">西邮评分
          </div>
          <br>
          <el-rate v-model="movieInfo.score" :colors="colors" disabled show-score text-color="#ff9900">
          </el-rate>
          <br>
          <div class="pf" style="color:#fff">累计票房<br>
            <h1 style="color:#F3E7FF;font-size:30px;margin-top:8px">{{ movieInfo.box_office }}亿</h1>
          </div>
        </div>



      </div>

    </div>
    <div class="main">
      <div class="left">
        <div class="introduction">
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span style="font-size:20px;color:#8FDCFE">电影介绍</span>

            </div>
            <div class="item222">
              <div>
                <h3>剧情简介</h3>
              </div>


              {{ movieInfo.description }}
            </div>
            <br><br><br><br><br><br>
          </el-card>
        </div>
      </div>
      <div class="right">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span style="font-size:20px;color:#8FDCFE">相似推荐</span>

          </div>
          <div v-for="item in recommend.ShowingList" :key="item" class="item">
            <img :src="'http://127.0.0.1:8080/api/getpicturebyfilename/' + item.cover_img_path" alt="">
            <div style="height:30px;margin-top: 20px;margin-left: 10px;width: 180px;"><a class="recommend_a" href="#"
                style="color:#219FFF;text-decoration:none">《 {{ item.name }} 》</a></div>
            <div style="margin-left:40px">
              <h2 style="margin-top:12px;color: #FFC600;">{{ item.score }}</h2>
            </div>
          </div>
        </el-card>
      </div>
    </div>
    <el-dialog title="提示" :visible.sync="rateDialogVisible" width="20%" center>
      <div class="block">
        <span class="demonstration">请打分</span>
        <el-rate v-model="value2" :colors="colors">
        </el-rate>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="rateDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="rateScore">确 定</el-button>
      </span>
    </el-dialog>
  </div>

</template>
<script>
import { removeDotSegments } from 'uri-js'

export default {
  data() {

    return {
      recommend: [{ "name": "这个杀手不太冷", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.5 },
      { "name": "这个杀手冷不冷", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.4 },
      { "name": "这个杀手有点冷", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.3 },
      { "name": "这个杀手冷死了", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.1 },],
      movieInfo: {},
      rateDialogVisible: false,
      value1: null,
      value2: null,
      colors: ['#99A9BF', '#F7BA2A', '#FF9900']
    }
  },
  created() {
    this.getMovieInfo()
  },
  methods: {
    openRate(id) {
      this.rateDialogVisible = true
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
      this.recommend = res.data.relevantMovies
    },

    buyTickets() {
      this.$router.push('/buytickets/' + this.movieInfo.id);
    },
    rateScore(){
      this.$message.success('感谢您的评价')
      this.rateDialogVisible=false
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

.left {
  margin-top: 50px;
  width: 700px;
  height: 500px;

}

.introduction {
  width: 700px;
  height: 400px
}

.right {
  margin-left: 50px;
  margin-top: 50px;
  width: 400px;
  height: 300px;

}

.item {
  margin-top: 10px;
  display: flex;

  img {
    height: 65px;
    width: 70px;
  }


}

.item:hover {
  background: #F7F7F7;
}

.recommend_a:hover {
  color: #ff6700;
}
</style>