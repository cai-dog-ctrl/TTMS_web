<template>
    <div>
        <!-- 面包屑导航区域 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>数据统计</el-breadcrumb-item>
            <el-breadcrumb-item>数据报表</el-breadcrumb-item>
        </el-breadcrumb>
    </div>
</template>

<script>
export default {
  
}
</script>

<style lang="less" scoped>
.el-card {
    margin-top: 5px;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.15) !important;
}
  
</style>