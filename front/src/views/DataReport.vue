<template>
    <div>
        <!-- 面包屑导航区域 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>数据统计</el-breadcrumb-item>
            <el-breadcrumb-item>统计图</el-breadcrumb-item>
        </el-breadcrumb>
        <!-- 卡片视图区 -->
        <el-card>
            <div id="main" style="width: 900px;height:600px;"></div>
        </el-card>
    </div>
</template>

<script>
import * as echarts from 'echarts'
export default {
    data() {
        return {
            
        }
    },

    created() {

    },

    mounted() {
    // 基于准备好的dom，初始化echarts实例
    var myChart = echarts.init(document.getElementById('main'))
    // 指定图表的配置项和数据
    var option = {
      title: {
        text: '票房统计'
      },
      tooltip: {},
      xAxis: {
        type: 'category',
        data: ['奇异','奇异2','花束','奇异3','奇异4','我和我的祖国','wangda1','坏蛋联盟','星际穿越','刿心剑','长津湖之水门桥','到达挚爱',]
      },
      yAxis: {type: 'value'},
      series: [
        {
          name: '销量',
          type: 'bar',
          data: [100, 120, 150, 144, 155, 160, 177, 200, 90, 100, 110, 130]
        }
      ]
    }
    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option)
  },

    methods: {

    },
}
</script>

<style lang="less" scoped>
.el-card {
    margin-top: 5px;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.15) !important;
}
</style>