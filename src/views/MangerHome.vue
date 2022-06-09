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
<<<<<<< HEAD
>>>>>>> 1145525110d732ca675081c628ce6ff150df6801
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
