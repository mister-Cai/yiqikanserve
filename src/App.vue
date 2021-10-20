<template>
  <div id="app">
      <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: "App",
  data() {
    return {
      superList: [
      ],
      timeList: []
    };
  },
  methods: {
  },
};
</script>

<style lang="scss">
:root {
  // 主要背景颜色
  --bg-color: #030827;
  // 主要颜色
  --primary-color: white;

  --size14: 14px;
  --size16: 16px;
  --size18: 18px;
  *::-webkit-scrollbar {
    display: none;
  }
  * {
    color: var(--primary-color);
    font-size: 13px;
  }
}
#app {
  background-color: var(--bg-color);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: -1;
  overflow: auto;
  
}
</style>
