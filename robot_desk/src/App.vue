<template>
  <div class="container">
      <suspense>
        <router-view/>
      </suspense>
  </div>
</template>
<script setup>
// import {onMounted} from "vue";
// import router from "./router/router.js";
// onMounted(to=>{
//   if(localStorage.getItem("token")!=="null"){
//     router.push({name: "home", replace: true});
//   }
// })
</script>