<template>
    <el-container>
        <el-header>
            <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item index="2-1">票房榜</el-menu-item>
                <el-menu-item index="2-2">评分榜</el-menu-item>
            </el-menu>
        </el-header>

        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>

</template>



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