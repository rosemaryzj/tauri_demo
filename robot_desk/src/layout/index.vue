<script setup>
import Sider from "./Sider.vue";
import Header from "./Header.vue";
</script>

<template>
  <div>
    <Header/>
    <Sider/>
    <router-view/>
  </div>
</template>