<script setup>
import {userStore} from "../store/store.js";
const userinfo = userStore();
</script>

<template>
  <div class="header-container">
    <div class="avatar">
      {{userinfo.username}} <el-avatar
        class="jpg"
          src="/images/avatar.png"
      />
<!--    <el-popover-->
<!--        placement="bottom"-->
<!--        title="evil.zhang"-->
<!--        :width="60"-->
<!--        trigger="click"-->
<!--        content=""-->
<!--    >-->
<!--      <template #reference>-->
<!--          <el-avatar-->
<!--              src="/images/avatar.jpg"-->
<!--          />-->
<!--      </template>-->
<!--    </el-popover>-->
    </div>

</div>
</template>

<style scoped lang="scss">
.header-container {
  background-color: rgba(24,30,42,0.92);
  text-align: center;
  height: 60px;
  width: 100%;
  line-height: 60px;
  border-bottom-left-radius: 3px;
  border-bottom-right-radius: 3px;
}

.avatar {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-right: 3%;
  height: 40px;
  padding-top: 10px;
  font-size: 20px;
  color: rgba(239, 233, 233, 0.8);
  .jpg {
    margin-left: 20px;
  }
}
</style>