<script setup>
const removeToken = ()=>{
  if(localStorage.getItem("token")!==null){
    localStorage.setItem("token", null)
  }
}
// import {onBeforeUnmount, onMounted} from "vue";
//
// const sideBarStore = useSideBarStore();
//
// onMounted(()=>{
//   window.addEventListener('resize', resizeWindow);
// })
//
// onBeforeUnmount(()=>{
//   window.removeEventListener('resize', resizeWindow);
// })

// let resizeWindow = ()=>{
//   let collapse = sideBarStore.collapsed;
//   if(window.innerWidth<=1200){
//     if(collapse){
//       sideBarStore.toggle();
//     }else {
//       if(!collapse){
//         sideBarStore.toggle();
//       }
//     }
//   }
// }
</script>

<template>
  <div class="side-left">
    <perfect-scrollbar>
      <ul class="navigation-left">
        <router-link
            to="#"
            tag="li"
            class="nav-item"
        >
          <div class="nav-item-hold">
            <i class="i-Bar-Chart text-3xl"></i>
            <p>机器人信息</p>
          </div>
        </router-link>
        <!--      <router-link-->
        <!--          to="#"-->
        <!--          tag="li"-->
        <!--          class="nav-item"-->
        <!--      >-->
        <!--        <div class="nav-item-hold">-->
        <!--          <i class="i-Wallet text-3xl"></i>-->
        <!--          <p>Buttons</p>-->
        <!--        </div>-->
        <!--      </router-link>-->
        <!--      <router-link-->
        <!--          to="#"-->
        <!--          tag="li"-->
        <!--          class="nav-item"-->
        <!--      >-->
        <!--        <div class="nav-item-hold">-->
        <!--          <i class="i-Find-User text-3xl"></i>-->
        <!--          <p>Profile</p>-->
        <!--        </div>-->
        <!--      </router-link>-->
        <!--      <router-link to="#" tag="li" class="nav-item">-->
        <!--        <div class="nav-item-hold">-->
        <!--          <i class="i-Checked-User text-3xl"></i>-->
        <!--          <p>Sign In</p>-->
        <!--        </div>-->
        <!--      </router-link>-->
        <router-link to="/login" tag="li" class="nav-item" @click="removeToken">
          <div class="nav-item-hold">
            <i class="i-Checked-User text-3xl"></i>
            <p>切换用户</p>
          </div>
        </router-link>
      </ul>
    </perfect-scrollbar>
  </div>
</template>

<style scoped lang="scss">
.side-left {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  width: 120px;
  background: #ffffff;
  color: #2f2e2e;
}

.navigation-left {
  list-style: none;
  text-align: center;
  width: 120px;
  height: 100%;
  margin: 0;
  padding: 0;

  .nav-item,
  .nav-item-single {
    position: relative;
    display: block;
    width: 100%;
    overflow: hidden;
    cursor: pointer;

    &:hover {
      .nav-item-hold {
        @apply text-purple-500;
      }

      &:after {
        content: '';
        position: absolute;
        width: 30px;
        height: 30px;
        bottom: -15px;
        right: -15px;
        transform: rotate(45deg);
        @apply bg-purple-500;
      }
    }

    &.active {
      @apply text-purple-500;

      &:after {
        content: '';
        position: absolute;
        width: 30px;
        height: 30px;
        bottom: -15px;
        right: -15px;
        transform: rotate(45deg);
        @apply bg-purple-500;
      }
    }

    border-bottom: 1px solid #dee2e6;

    .nav-item-hold {
      display: block;
      width: 100%;
      padding: 26px 0;

      span.material-icons {
        font-size: 2rem;
      }
    }
  }
}
</style>