<template>
    <div>
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>剧目管理</el-breadcrumb-item>
            <el-breadcrumb-item>演出计划</el-breadcrumb-item>
        </el-breadcrumb>

        <div v-for="item in first_page_info.ShowingList.ShowingList" :key="item.id" class="movie">
            <el-card :body-style="{ padding: '0px' }" shadow="hover">
                <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                    class="image" @click="gotoinfo(item.id)">
                <div style="padding: 14px;">
                    <span>{{ item.name }}</span>
                    <div class="bottom clearfix">

                        <el-button type="text" class="button" @click="gotoinfo(item.id)">特惠购买</el-button>
                    </div>
                </div>
            </el-card>
        </div>


    </div>
</template>

<script>
export default {
    data() {
        return {

        };
    },

    created() {

    },

    methods: {

    },
}
</script>

<style lang="less" scoped>
</style>