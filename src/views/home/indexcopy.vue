<template>
    <div class="page">
        <div class="bg">
            <div class="bg1"></div>
            <div class="bg2"></div>
            <div class="bg3"></div>
        </div>
        <div class="page_header">
            <div class="logo"></div>
            FM3D.one
        </div>
        <div class="superinfo">
            <div class="title">
                当前奖超级大乐透池
                <span class="tip"></span>
            </div>
            <div class="number">18650,658 FM3D</div>
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
                当前奖超级大乐透池
                <span class="tip"></span>
            </div>
            <div class="number green">18650,658 FM3D</div>
            <div class="over_time_tip">倒计时</div>
            <div class="over_time">
                <div class="time_block">08</div>
                <div class="time_split">:</div>
                <div class="time_block">59</div>
                <div class="time_split">:</div>
                <div class="time_block">00</div>
            </div>
        </div>

        <div class="now_button">立即抢购 FM3D
            <span class="tip">Pancakeswap</span>
        </div>

        <div class="list_view">
            <div class="list_title">
                超级大乐透奖池名单
            </div>
            <div class="row header" style="background: rgba(255, 0, 142, 0.05);">
                <div>最新抢购地址</div>
                <div>数量(FM3D)</div>
                <div>中奖(FM3D)</div>
            </div>
            <div class="row" v-for="(row, index) in superList" :key="index">
                <div>{{row.address}}</div>
                <div>{{row.amount}}</div>
                <div>{{row.status}}</div>
            </div>
        </div>

        <div class="list_view">
            <div class="list_title" style="color:#00FF55;">
                时间乐透奖池名单
            </div>
            <div class="row header" style="background: rgba(0, 255, 85, 0.05);">
                <div style="color:#00FF55;">当前累计购买前五名</div>
                <div style="color:#00FF55;">数量(FM3D)</div>
                <div style="color:#00FF55;">中奖(FM3D)</div>
            </div>
            <div class="row" v-for="(row, index) in timeList" :key="index">
                <div>{{row.address}}</div>
                <div>{{row.amount}}</div>
                <div>{{row.status}}</div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                superList: [
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                ],
                timeList: [
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                    { address: '0x12312321312312321321', amount: 20000, status: '待开奖' },
                ]
            }
        },
    }
</script>

<style lang="scss" scoped>
.page {
    @media screen and (max-width: 768px) {
        width: 100%;
        padding: 15px 0 40px 0;
    }
    @media screen and (min-width: 768px) {
        margin: 40px auto 60px;
        width: 1291px;
        .bg {
            .bg3 {
                width: 536px;
                margin: 0 auto;;
            }
        }
    }
    position: relative;
    .bg {
        z-index: 1000;
        position: absolute;
        top: 0;
        width: 100%;
        .bg1 {
            background: url('../../assets/bg1.png') no-repeat;
            background-size: 100% 100%;
            height: 768px;
        }
        .bg2 {
            background: url('../../assets/bg2.png') no-repeat;
            background-size: 100% 100%;
            height: 340px;
            opacity: 0.5;
        }
        .bg3 {
            background: url('../../assets/bg3.png') no-repeat;
            background-size: 100% 100%;
            height: 325px;
            opacity: 0.35;
        }
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
            color: #FF008E;
            line-height: 42px;
            &.green {
                color: #00FF55;
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
                background: rgba(255, 0, 142, 0.15);
                border-radius: 4px;
                font-size: 28px;
                font-family: PingFangSC-Semibold, PingFang SC;
                font-weight: 600;
                color: #FFFFFF;
                line-height: 39px;
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
            margin: 78px auto;
        }
        height: 54px;
        background: linear-gradient(180deg, #FF008E 0%, #A8005E 100%);
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
        margin-top: 80px;
        .list_title {
            text-align: center;
            font-size: 16px;
            font-family: PingFangSC-Medium, PingFang SC;
            font-weight: 500;
            color: #FF008E;
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
                    color: #FF008E;
                }
            }
            div:nth-child(1) {
                text-align: left;
                flex: 2;
            } 
            div:nth-child(2) {
                text-align: center;
                flex: 1;
            }
            div:nth-child(3) {
                text-align: right;
                flex: 1;
            }
        }
    }
}

</style>