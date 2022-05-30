<template>
    <!-- <el-container>
        <el-header>
            <el-menu :default-active="active1" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item :index="active1" @click="topBorad">票房榜</el-menu-item>
                <el-menu-item :index="active2" @click="tosBroad">评分榜</el-menu-item>
            </el-menu>
        </el-header>

        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container> -->
    <div class="main">
        <el-tabs v-model="activeName" @tab-click="handleClick" :stretch="true">
            <el-tab-pane label="票房榜" name="first">用户管理</el-tab-pane>
            
            <el-tab-pane label="评分榜" name="second">配置管理</el-tab-pane>
            
        </el-tabs>
    </div>

</template>

<script>
export default {
    name: 'WorkspaceJsonBorad',
    data() {
        return {
            active1: 'pborad',
            active2: 'sborad'

        };
    },
    methods: {
        tosBroad() {
            this.$router.push('/sborad');
        },
        topBorad() {
            this.$router.push('/borad');
        }
    },
};
</script>



<style scoped>
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
.main{
    min-width: 900px;
    max-width: 900px;
    text-align: center;
    margin: auto;
    margin-top: 25px;
}
</style>