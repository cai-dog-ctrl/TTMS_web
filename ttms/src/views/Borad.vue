<template>
    <el-container>
        <el-header>
            <el-menu :default-active="active1" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item :index="active1">票房榜</el-menu-item>
                <el-menu-item :index="active2" @click="tosBroad">评分榜</el-menu-item>
            </el-menu>
        </el-header>

        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>

</template>

<script>
export default {
    name: 'WorkspaceJsonBorad',
    data() {
        return {
            active1:'pborad',
            active2:''
            
        };
    },
    methods: {
        tosBroad() {
            this.$router.push('/sborad');
        },
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
</style>