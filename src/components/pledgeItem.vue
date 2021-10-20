<template>
  <div class="flow_item">
    <div class="flow_item__status">NEW</div>
    <div class="flow_item__title">
      CAME{{ $t("收益池") }}
      <div class="icon"></div>
    </div>
    <div class="flow_item__content">
      <div class="content_statictis">
        <div class="sta_left sta_icon_left">
          <img src="../assets/pledge_item_bg.png" alt="" />
        </div>
        <div class="sta_right sta_icon_right">
          <div class="right_title">{{ $t("总质押") }}</div>
          <div class="right_value">1123123212 CAME</div>
          <div class="right_title">{{ $t("日产量") }}</div>
          <div class="right_value">960 CAME</div>
        </div>
      </div>

      <div class="content_statictis">
        <div class="sta_left">
          <div>{{ $t("年化收益：") }}</div>
          <div class="progcess">
            <!-- <van-circle
              :value="16.12"
              :size="112"
              color="#00FFF8"
              stroke-width="100"
              layer-color="#1E2029"
            > -->
              <div class="circle_text">16.12%</div>
            <!-- </van-circle> -->
          </div>
        </div>
        <div class="sta_right">
          <div class="right_title">{{ $t("我的质押") }}</div>
          <div class="right_value">0</div>
          <div class="right_title">{{ $t("我的质押") }}</div>
          <div class="right_value">0 CAME</div>
        </div>
      </div>
      <div class="button_view">
        <div class="button_flex_row">
          <div class="button button_blue">{{ $t("赎回") }}</div>
          <div class="button button_red">{{ $t("质押") }}</div>
        </div>
        <div class="button_flex_row">
          <div class="button button_grey">{{ $t("授权") }}HUSDT</div>
        </div>
        <div class="button_flex_row">
          <div class="button button_grey">{{ $t("领取收益") }}</div>
        </div>
      </div>
      <div class="to_get_lp">{{ $t("可用") }}：0.000000000HUSDT</div>
    </div>
  </div>
</template>

<script>
// import { Circle } from "vant";
export default {
  name: "FlowItem",
  // components: { vanCircle: Circle },
  data() {
    return {
      process: 10,
    };
  },
};
</script>

<style lang="scss" scoped>
.flow_item {
  margin-top: 20px;
  .flow_item__status {
    text-align: right;
    padding-right: 15px;
    line-height: 30px;
    background: url("../assets/homepage_title_bg_right.png") no-repeat;
    font-size: 14px;
    font-family: PingFangSC-Medium, PingFang SC;
    font-weight: 500;
    color: #ffffff;
  }
  .is_new {
    background: url("../assets/header_bar_red.png") no-repeat;
  }
  .flow_item__title {
    background: url("../assets/header_sub_bar_blue.png") no-repeat;
    background-size: 100% 100%;
    height: 60px;
    line-height: 60px;
    padding: 0 30px;
    position: relative;
    &.is_new {
      background: url("../assets/destory_came.png") no-repeat;
    }
    .icon {
      top: 10px;
      right: 5px;
      bottom: 2px;
      width: 50%;
      position: absolute;
      background: url("../assets/pledge_header_bg.png") no-repeat;
      background-size: 100% 100%;
    }
  }
  .flow_item__content {
    margin-top: 10px;
    background-color: #2d303e;
    .content_statictis {
      display: flex;
      margin-top: 30px;
      &:first-child {
        margin-top: 0;
      }
      .sta_left {
        padding-left: 15px;
        flex: 2;
        font-size: 16px;
        font-family: PingFangSC-Medium, PingFang SC;
        font-weight: 500;
        color: #d6dff2;
        line-height: 22px;
        .progcess {
          margin: 32px 0;
          display: flex;
          justify-content: center;
          align-items: center;
        }
      }
      .sta_icon_left {
        position: absolute;
        flex: 2;
        width: 50%;
        img {
          width: 100%;
          height: 100%;
        }
      }
      .sta_right {
        flex: 3;
        text-align: right;
        .right_title {
          padding-right: 15px;
          background: url("../assets/header_bar_black.png") no-repeat;
          background-size: 100% 100%;
          line-height: 30px;
          font-size: 14px;
          font-family: PingFangSC-Medium, PingFang SC;
          font-weight: 500;
        }
        .right_value {
          padding-right: 15px;
          line-height: 60px;
          font-size: 20px;
          font-family: PingFangSC-Medium, PingFang SC;
          font-weight: 500;
        }
      }
      .sta_icon_right {
        flex: 4;
        padding-left: 35%;
        .right_title {
          background: url("../assets/homepage_title_bg_right.png") no-repeat;
          background-size: 100% 100%;
        }
      }
    }
  }
}
.circle_text {
  width: 100px;
  height: 100px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 16px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #00fff8;
  line-height: 22px;
}
.button_view {
  margin: 0 27px;
  .button_flex_row {
    display: flex;
    .button {
      flex: 1;
      text-align: center;
      width: 100%;
      height: 44px;
      line-height: 44px;
      cursor: pointer;
      &.button_blue {
        background: url("../assets/exchange_button_view.png") no-repeat;
        background-size: 100% 100%;
      }
      &.button_red {
        background: url("../assets/button_red.png") no-repeat;
        background-size: 100% 100%;
      }
      &.button_grey {
        margin-top: 30px;
        background: url("../assets/button_black.png") no-repeat;
        background-size: 100% 100%;
        font-size: 18px;
        font-family: PingFangSC-Semibold, PingFang SC;
        font-weight: 600;
        color: #00fff8;
      }
    }
  }
}

.wallet_usable {
  margin-top: 31px;
  padding-right: 15px;
  text-align: right;
  font-size: 16px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #d6dff2;
  line-height: 22px;
}
.wallet_balance {
  margin-top: 6px;
  padding-right: 15px;
  text-align: right;
  font-size: 20px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #ffffff;
  line-height: 28px;
}
.to_get_lp {
  margin-top: 30px;
  padding-right: 15px;
  line-height: 30px;
  background: url("../assets/header_bar_black.png") no-repeat;
  font-size: 14px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #ffffff;
  text-align: right;
}
</style>