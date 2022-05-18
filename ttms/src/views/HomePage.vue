<template>
    <el-container>
        
            <el-header>

                <el-menu :default-active="url1" class="el-menu-demo" mode="horizontal">
                    <el-menu-item :index="url1" class="left1">首页</el-menu-item>
                    <el-menu-item :index="url2" class="left1">电影</el-menu-item>
                    <el-menu-item :index="url3" class="left1">榜单</el-menu-item>

                    <el-dropdown>
                        <img class="img" src="https://p0.meituan.net/movie/7dd82a16316ab32c8359debdb04396ef2897.png" alt="">
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item>登录</el-dropdown-item>

                        </el-dropdown-menu>
                    </el-dropdown>

                </el-menu>
            </el-header>
        
        <el-main>
            <router-view></router-view>
        </el-main>
        
    </el-container>
</template>

<script>


export default {

    data() {
        return {
            url1: "FisrtPage",
            url2: "borad",
            url3: "movie",
        }
    },
    methods: {

    }

}
</script>

<style>
.el-container {
    height: 100%;

}


.el-footer {
    bottom: 0;
}

.img {
    height: 45px;
    width: 45px;
    border-radius: 50%;
    cursor: pointer;
}

.people_logo {
    height: 60px;
}

.el-dropdown {
    margin-right: 25px;
    margin-top: 10px;
    float: right;
}

</style>