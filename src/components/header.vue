<template>
  <Row class="house_header" type="felx" justify="space-between">
    <Col span="12" class="header_left">
      <icon name="wap-home-o" />
      <div class="text">长乐街16A</div>
    </Col>
    <Col span="12" class="header_right">
      <icon name="chat-o" />
      <!-- <icon name="plus"></icon> -->
    </Col>
  </Row>
</template>

<script>
import { Row, Col, Icon } from "vant";
export default {
  name: "AppHeader",
  components: { Row, Col, Icon },
  data() {
    return {
      active: false,
      langIcons: {
        en: require("../assets/lang_en.png"),
        zh: require("../assets/lang_zh.png"),
      },
    };
  },
  computed: {
    locale() {
      return this.$i18n.locale;
    },
  },
  methods: {
    onChangeLanguage() {
      this.$i18n.locale = this.locale === "zh" ? "en" : "zh";
    },
  },
};
</script>

<style lang="scss" scoped>
.house_header {
  padding: 8px 16px;
  background: white;
  .header_left {
    display: flex;
   .van-icon {
      padding-right: 5px;
      font-size: 1.5em;
    }
    .text {
      line-height: 1.5em;
    }
  }
  .header_right {
    text-align: right;
    .van-icon {
      font-size: 1.5em;
      padding-left: 10px;
    }
  }
}
</style>