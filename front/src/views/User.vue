<template>
    <div class="main_order">
        <div class="aside">
            <div class="Head">
                <div>
                    <h1>个人信息</h1>
                </div>

            </div>
            <div class="BODY">
                <div class="Info_user">
                    <div><span>基本信息</span></div>
                </div>
                <div class="Info_order" @click="goto_order()">
                    <div><span>我的订单</span></div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
export default{
    data(){
        return{

        }
    },
    methods:{
        goto_order(){
            this.$router.push('/order')
        }
    }
}
</script>

<style lang="less" scoped>
.main_order {
    height: 800px;
    width: 1100px;
    border: #000 solid 1px;
    text-align: center;
    margin-top: 50px;
    margin-left: auto;
    margin-right: auto;
    display: flex;
}

.aside {
    width: 250px;
    height: 100%;
    border: #000 solid 1px;
    margin-left: 0;
}

.Head {
    width: 100%;
    height: 60px;
    

    div {
        
        color: #219FFF;
        font-size: 20px;
    }
}


.Info_user,
.Info_order {
    width: 100%;
    height: 30px;

    font-size: 19px;
    padding-top: 10px;
    cursor: pointer;
}

.Info_user {
    background: #219FFF;
    color: #fff;
}

.Info_order {
    color: #000;
}
</style>