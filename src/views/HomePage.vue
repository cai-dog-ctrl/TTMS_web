<template>
    <el-container>

        <el-header>

            <el-menu :default-active="url1" class="el-menu-demo" mode="horizontal">
                <el-menu-item>
                    <div class="image"><img src="../assets/logo.png" alt="" class="imga"></div>
                </el-menu-item>

                <el-menu-item :index="url1" class="left1" @click="toFirstPage">首页</el-menu-item>
                <el-menu-item :index="url2" class="left1" @click="tomovie">电影</el-menu-item>
                <el-menu-item :index="url3" class="left1" @click="toBorad">榜单</el-menu-item>
                <el-menu-item :index="url4" class="left1" @click="User">订单</el-menu-item>
                <el-dropdown>
                    <img class="img" src="https://p0.meituan.net/movie/7dd82a16316ab32c8359debdb04396ef2897.png" alt="">
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item>
                            <p v-if="IsLogin === false"><a @click="login">登录</a></p>

                        </el-dropdown-item>
                        <el-dropdown-item>
                            <p><a @click="logout">退出</a></p>
                        </el-dropdown-item>
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
            url4: "userInfo",
            IsLogin: false,
            token: ''
        }
    },
    created() {
        this.token = window.sessionStorage.getItem('token');
        if (this.token === null) {
            this.IsLogin = false
        } else {
            this.IsLogin = true
        }
        //alert(this.token);
    },
    methods: {
        User(){
            this.$router.push('/order');
        },
        login() {
            this.$router.push('/login');
        },
        logout() {
            window.sessionStorage.removeItem('token');
            window.sessionStorage.removeItem('userid')
            this.IsLogin = false
        },
        toBorad() {
            this.$router.push('/borad');
        },
        toFirstPage() {
            this.$router.push('/first');
        },
        tomovie() {
            this.$router.push('/movie');
        }
    }

}
</script>

<style scoped>
.el-container {
    height: 100%;

}

.el-main {
    width: 100%;
    padding: 0;
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

.el-main {
    padding: 0;
}

.left1 {
    margin-left: 10%;
}

.image {
    height: 40px;
    width: 600px;
}

.imga {
    margin-left: 40%;
    height: 40px;
}

.el-menu {
    font-size: 40px;
}
</style>