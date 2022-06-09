<template>
    <el-container class="home-container">
        <!--头部区域---->
        <el-header>
            <div>
                <img src="../assets/logo.jpg" alt="">
                <span>有模有样影院后台管理系统</span>
            </div>
            <el-button type='info' @click="logout">退出</el-button>
        </el-header>
        <!--页面主体区-->
        <el-container>
            <!--侧边栏-->
            <el-aside :width="isCollapse ? '64px' : '200px'">
                <div class="toggle-button" @click="toggleCollapse">|||</div>
                <!--侧边栏菜单区域---->
                <el-menu background-color="#333744" text-color="#fff" active-text-color="#409fff" :unique-opened='true'
                    :collapse="isCollapse" :collapse-transition="false" router :default-active="activePath">
                    <!--一级菜单-->
                    <el-submenu :index="String(item.id)" v-for="item in menulist" :key="item.id">
                        <!--一级菜单模板区---->
                        <template slot="title">
                            <!--图标-->
                            <i :class="iconsObj[item.id]"></i>
                            <!--文本-->
                            <span>{{ item.authName }}</span>
                        </template>

                        <!--二级菜单---->
                        <el-menu-item :index="'/' + String(subItem.path)" v-for="subItem in item.children"
                            :key="subItem.id" @click="saveNavState('/' + String(subItem.path))">
                            <template slot="title">
                                <!--图标-->
                                <i class="el-icon-menu"></i>
                                <!--文本-->
                                <span>{{ subItem.authName }}</span>
                            </template>
                        </el-menu-item>
                    </el-submenu>
                </el-menu>
            </el-aside>
            <!--右侧内容主体-->
            <el-main>
                <!--路由占位符-->
                <router-view></router-view>
            </el-main>
        </el-container>

    </el-container>

</template>

<script>
export default {
    data() {
        return {
            //左侧菜单数据
            menulist: [
                {
                    "id": 125,
                    "authName": "用户管理",
                    "path": "users",
                    "children": [
                        {
                            "id": 110,
                            "authName": "用户列表",
                            "path": "mangeruser",
                   
                        }
                    ],
                  
                },
                {
                    "id": 103,
                    "authName": "影厅管理",
                    "path": "rights",
                    "children": [
                        {
                            "id": 111,
                            "authName": "影厅列表",
                            "path": "moviehall",
                          
                        },

                    ],
                   
                },
                {
                    "id": 101,
                    "authName": "剧目管理",
                    "path": "goods",
                    "children": [
                        {
                            "id": 104,
                            "authName": "剧目列表",
                            "path": "manger_movie",
                           
                        },
                        {
                            "id": 115,
                            "authName": "演出计划",
                            "path": "movieplan",
                           
                        },

                    ],
                   
                },
                {
                    "id": 102,
                    "authName": "数据统计",
                    "path": "data",
                    "children": [
                        {
                            "id": 111,
                            "authName": "总销售额",
                            "path": "datastatistic",
                            
                        },
                        {
                            "id":111,
                            "authName": "日销售额",
                            "path": "daydatastatistic",
                        },
                        {
                            "id":111,
                            "authName": "月销售额",
                            "path": "monthdatastatistic",
                        },
                        {
                            "id":111,
                            "authName": "年销售额",
                            "path": "yeardatastatistic",
                        }
                    ],
                    
                },
            ],
            iconsObj: {
                '125': 'el-icon-user-solid',
                '103': 'el-icon-s-management',
                '101': 'el-icon-video-camera-solid',
                '102': 'el-icon-s-order',
                '145': 'el-icon-s-marketing'
            },
            //是否折叠
            isCollapse: false,
            //被激活的链接地址
            activePath: ''
        }

    },
    created() {
        // this.getMenuList()
        this.activePath = window.sessionStorage.getItem('activePath')
    },
    methods: {
        logout() {
            window.sessionStorage.clear();
            this.$router.push("/login");
        },
        //获取所有的菜单
        async getMenuList() {
            const { data: res } = await this.$http.get('menus')
            if (res.list !== 1000) return this.$message.error(res.meta.msg)
            this.menulist = res.data
            console.log(res)
        },
        //点击按钮，切换菜单的折叠与展开
        toggleCollapse() {
            this.isCollapse = !this.isCollapse
        },
        //保存链接的激活状态
        saveNavState(activePath) {
            window.sessionStorage.setItem('activePath', activePath)
            this.activePath = activePath

        }
    },

};
</script>

<style lang="less" scoped>
.home-container {
    height: 100%;
}

.el-header {
    background-color: #373d41;
    display: flex;
    justify-content: space-between;
    padding-left: 0;
    align-items: center;
    color: #fff;
    font-size: 20px;

    >div {
        display: flex;
        align-items: center;

        span {
            margin-left: 15px;
        }
    }
}

.el-aside {
    background-color: #333744;

    .el-menu {
        border-right: none;
    }

}

.el-main {
    background-color: #eaedf1;
}

.toggle-button {
    background-color: #4a5064;
    font-size: 10px;
    line-height: 24px;
    color: #fff;
    text-align: center;
    letter-spacing: 0.2em;
    cursor: pointer; //悬空变小手

}
</style>