<template>
    <div class="page">
        <div class="bg1">
            <!-- <div class="bg">
                <div class="bg1"></div>
                <div class="bg2"></div>
                <div class="bg3"></div>
            </div> -->
            <div class="page_header">
                <div class="logo"></div>
                F3D.one
            </div>
            <div class="superinfo">
                <div class="title">
                    当前奖超级大乐透池
                    <el-tooltip class="item" effect="dark" content="每次开奖为奖池的80%，按照最后五名分别可获得45%、25%、15%、10%、5%" placement="bottom-end">
                        <span class="tip"></span>
                    </el-tooltip>
                </div>
                <div class="number">{{Math.floor(balance*0.8*100)/100}} USDT</div>
                <div class="over_time_tip">倒计时</div>
                <div class="over_time">
                    <div class="time_block">08</div>
                    <div class="time_split">:</div>
                    <div class="time_block">59</div>
                    <div class="time_split">:</div>
                    <div class="time_block">00</div>
                </div>
            </div>
            <div class="solid"></div>
            <div class="superinfo">
                <div class="title">
                    当前时时乐透池
                    <el-tooltip class="item" effect="dark" content="每次开奖为超级大乐透奖池的5%，按照一小时内购买量前五名，按购买数量权重瓜分奖池。" placement="bottom-end">
                        <span class="tip"></span>
                    </el-tooltip>
                </div>
                <div class="number violet">{{Math.floor(balance*0.02*100)/100}} USDT</div>
                <div class="over_time_tip">倒计时</div>
                <div class="over_time">
                    <div class="time_block violet">08</div>
                    <div class="time_split">:</div>
                    <div class="time_block violet">59</div>
                    <div class="time_split">:</div>
                    <div class="time_block violet">00</div>
                </div>
            </div>

            <div class="now_button">立即抢购 F3D
                <span class="tip">Pancakeswap</span>
            </div>
        </div>
        <div class="list_view bg2">
            <div class="list_title">
                超级大乐透奖池名单
            </div>
            <div class="row header" style="background: rgba(255, 0, 142, 0.05);">
                <div>最新抢购地址</div>
                <div>数量(F3D)</div>
            </div>
            <div class="row" v-for="(row, index) in superList5" :key="index">
                <div>{{row.add?row.add.replace(row.add.substring(8,32),'.....'):''}}</div>
                <div>{{row.mount}}</div>
            </div>
            <div class="row header" style="background: rgba(255, 0, 142, 0.05);">
                <div>上期获奖名单</div>
                <div>中奖(USDT)</div>
            </div>
            <div class="row" v-for="(row, index1) in superList" :key="index1">
                <div>{{row.add?row.add.replace(row.add.substring(8,32),'.....'):''}}</div>
                <div>{{Math.floor(row.mount*balance*0.8*100)/100}}</div>
            </div>
        </div>
        <div class="list_view bg3">
            <div class="list_title" style="color:#D000FF;">
                时时乐透奖池名单
            </div>
            <div class="row header" style="background: rgba(0, 255, 85, 0.05);">
                <div style="color:#D000FF;">当前累计购买前五名</div>
                <div style="color:#D000FF;">数量(F3D)</div>
            </div>
            <div class="row" v-for="(row, index2) in timelist5" :key="index2">
                <div>{{row.add?row.add.replace(row.add.substring(8,32),'.....'):''}}</div>
                <div>{{row.mount}}</div>
            </div>
            <div class="row header" style="background: rgba(0, 255, 85, 0.05);">
                <div style="color:#D000FF;">上期获奖名单</div>
                <div style="color:#D000FF;">中奖(USDT)</div>
            </div>
            <div class="row" v-for="(row, index3) in timelist" :key="index3">
                <div>{{row.add?row.add.replace(row.add.substring(8,32),'.....'):''}}</div>
                <div>{{Math.floor(row.mount/timeamount*balance*0.02*100)/100}}</div>
            </div>
        </div>
    </div>
</template>

<script>
import {mapState} from 'vuex'
    export default {
        data() {
            return {
            }
        },
        created(){
            this.$store.dispatch("init")
        },
        computed:{
            ...mapState(['balance','superList5','timelist5','superList','timelist','timeamount'])
        }
    }
</script>

<style lang="scss" scoped>
.page {
    @media screen and (max-width: 768px) {
        width: 100%;
    }
    @media screen and (min-width: 768px) {
        margin: 0 auto;
        width: 1291px;
        .list_view{
            padding:60px 40px 30px !important;
        }
        .bg {
            .bg3 {
                width: 536px;
                margin: 0 auto;;
            }
        }
    }
    .bg1 {
        background: url('../../assets/bg1.png') repeat;
        background-size: 1291px 100%;
        background-position-x: center;
        padding-top:20px;
    }
    .bg2 {
        background: url('../../assets/bg2.png') no-repeat;
        background-size: 100% 100%;
        // height: 340px;
        // opacity: 0.5;
        padding-bottom: 20px;
    }
    .bg3 {
        background: url('../../assets/bg3.png') no-repeat;
        background-size: 100% 100%;
        // height: 325px;
        // opacity: 0.35;
        }
    
    
    .page_header {
        display: flex;
        padding-left: 18px;
        .logo {
            width: 30px;
            height: 30px;
            background: url('../../assets/logo.png') no-repeat;
            background-size: 100% 100%;
            margin-right: 10px;
        }
        line-height: 30px;
        margin-bottom: 145px;
    }

    .superinfo {
        .title {
            text-align: center;
            font-size: 30px;
            font-family: PingFangSC-Semibold, PingFang SC;
            font-weight: 600;
            color: #FFFFFF;
            .tip {
                margin-top: 8px;
                margin-left: 8px;
                position: absolute;
                width: 26px;
                height: 26px;
                background: url('../../assets/question.png') no-repeat;
                background-size: 100% 100%;
            }
        }   
        .number {
            margin-top: 10px;
            text-align: center;
            font-size: 30px;
            font-family: PingFangSC-Medium, PingFang SC;
            font-weight: 500;
            color: #00FFCB;
            line-height: 42px;
            &.violet {
                color: #D000FF;
            }
        }
        .over_time_tip {
            text-align: center;
            margin-top: 8px;
            font-size: 20px;
            font-family: PingFangSC-Medium, PingFang SC;
            font-weight: 500;
            line-height: 28px;
        }
        .over_time {
            margin-top: 8px;
            display: flex;
            justify-content: center;
            .time_block {
                text-align: center;
                width: 44px;
                height: 44px;
                background: #00FFCB;
                border-radius: 4px;
                font-size: 28px;
                font-family: PingFangSC-Semibold, PingFang SC;
                font-weight: 600;
                color: #192337;
                line-height: 39px;
                &.violet{
                    background: #D000FF;
                }
            }
            
            .time_split {
                margin: 0 8px;
                font-size: 28px;
                font-family: PingFangSC-Medium, PingFang SC;
                font-weight: 500;
                color: #FFFFFF;
                line-height: 39px;
            }
        }
    }

    .solid {
        margin: 60px 0;
        height: 1px;
        background-image: linear-gradient(to right, rgba(255, 255, 255, 0.15) 0%, rgba(255, 255, 255, 0.15) 50%, transparent 50%);
        background-size: 12px 1px;
        background-repeat: repeat-x;
    }

    .now_button {
        @media screen and (max-width: 768px) {
            margin: 78px 30px 0;
        }
        @media screen and (min-width: 768px) {
            width: 400px;
            margin: 78px auto 0;
        }
        height: 54px;
        background: linear-gradient(180deg, #00FFCB 0%, #009173 100%);
        border-radius: 8px;
        text-align: center;
        font-size: 20px;
        font-family: PingFangSC-Semibold, PingFang SC;
        font-weight: 600;
        color: #FFFFFF;
        display: flex;
        justify-content: center;
        align-items: center;
        position: relative;
        .tip {
            position: absolute;
            right: 4px;
            bottom: 4px;
            left: 0;
            text-align: right;
            height: 17px;
            font-size: 12px;
            font-family: PingFangSC-Medium, PingFang SC;
            font-weight: 500;
            color: #FFFFFF;
            line-height: 17px;
        }
    }

    .list_view {
        padding: 60px 0 30px;
        .list_title {
            text-align: center;
            font-size: 16px;
            font-family: PingFangSC-Medium, PingFang SC;
            font-weight: 500;
            color: #00FFCB;
            line-height: 38px;
        }

        .row {
            padding: 0 19px;
            display: flex;
            line-height: 34px;
            &.header {
                font-size: 13px;
                font-family: PingFangSC-Medium, PingFang SC;
                font-weight: 500;
                line-height: 50px;
                div {
                    color: #00FFCB;
                }
            }
            div:nth-child(1) {
                text-align: left;
                flex: 2;
            } 
            div:nth-child(2) {
                text-align: right;
                flex: 1;
            }
        }
    }
}

</style>