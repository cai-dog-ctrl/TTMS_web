<template>
  <div class="container">
    <div class="banner">
      <div class="flexa">
        <div class="imag"><img class="image"
            src="https://p0.pipi.cn/mmdb/25bfd69a030c7eaf330e13fb0b08a6695f6f7.jpg?imageView2/1/w/464/h/644" alt="">
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
          <br>
          <div class="buy_btn">特惠购票</div>
        </div>
        <div style="margin-top:120px;margin-left: 80px;">
          <div class="score" style="color:#fff">西邮评分<h1 style="color:#FFC600;font-size:30px;margin-top:10px">9.1</h1>
          </div>
          <div class="pf" style="color:#fff">累计票房<br>
            <h1 style="color:#F3E7FF;font-size:30px;margin-top:8px">26.3亿</h1>
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


              马丽魏翔携手演绎高口碑爆笑喜剧！《夏洛特烦恼》《西虹市首富》导演闫非、彭大魔监制，邢文雄执导，马丽、魏翔等喜剧人主演，携手打造爆笑“杀手奇遇”！小人物龙套演员魏成功（魏翔 饰）得到当红女星米兰（马丽
              饰）“赏识”，被邀请出演男一号神秘杀手，却没想到这一切竟是一场危机四伏的“戏中戏”、“局中局”。被蒙在鼓里的魏成功数次凭借自己的演技和好运化险为夷，而这场阴谋却逐渐失控，残酷的真相也离他越来越近……
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
          <div v-for="item in recommend" :key="item" class="item">
            <img :src="item.img" alt="">
            <div style="height:30px;margin-top: 20px;margin-left: 10px;"><a class="recommend_a" href="#"
                style="color:#219FFF;text-decoration:none">《 {{ item.name }} 》</a></div>
            <div style="margin-left:80px">
              <h2 style="margin-top:12px;color: ;">{{ item.socre }}</h2>
            </div>
          </div>
        </el-card>
      </div>
    </div>
  </div>

</template>
<script>
export default {
  data() {
    return {
      recommend: [{ "name": "这个杀手不太冷", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.5 },
      { "name": "这个杀手冷不冷", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.4 },
      { "name": "这个杀手有点冷", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.3 },
      { "name": "这个杀手冷死了", "img": "https://p0.pipi.cn/mmdb/25bfd6877e15bfc7ed9257a2a0ba131a9d2fb.jpg?imageView2/1/w/464/h/644", "socre": 4.1 },]
    }
  },
  created() {

  },
  methods: {

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

  a {}

  a:hover {}
}

.item:hover {
  background: #F7F7F7;
}

.recommend_a:hover {
  color: #ff6700;
}
</style>