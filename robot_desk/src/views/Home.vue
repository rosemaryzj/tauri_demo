<script setup>
import {invoke} from "@tauri-apps/api/tauri";
import {onBeforeUnmount, onMounted, reactive, ref} from "vue";
import {Eleme, InfoFilled} from '@element-plus/icons-vue'
import BaseCard from "../components/Base/BaseCard.vue";
import {listen} from "@tauri-apps/api/event";
import {ElNotification} from "element-plus";
import Pop from "../components/Pop.vue"
const total_robots = ref(0);
const loading = ref(false);
const queryBtnLoading = ref(false);
const addBtnLoading = ref(false);
const isModalOpen = ref(false);
const page = ref(1);
const unListen = await listen('network-error', (event) => {
  ElNotification({
    title: '网络错误',
    message: event.payload.message,
    duration: 2000,
    type: "error",
    showClose: false
  })
})

onBeforeUnmount(() => {
  unListen();
})

function closeModal(e) {
  isModalOpen.value = e;
}

function addRobot() {
  isModalOpen.value = true;
}

const robots = ref([]);
const formInline = reactive({
  user: '',
  region: '',
  date: '',
})
// 分页查询机器人
const fetchRobots = async () => {
  loading.value = true;
  try {
    const response = await invoke("fetch_robots", {paginator: {page: page.value, page_size: 5}});
    if (response !== null && response !== undefined) {
      if (response.meta.code !== 200) {
        ElNotification({
          title: 'get robots',
          message: response.meta.message,
          duration: 2000,
          type: "error",
          showClose: false
        })
      } else {
        robots.value = response.data;
        total_robots.value = response.total;
      }
    }
  } catch (error) {
    console.error("An error occurred:", error);
  } finally {
    loading.value = false;
  }
};
// 在 setup 函数内注册生命周期钩子
onMounted(() => {
  fetchRobots();
});
const confirmDeleteRobot = ()=>{
  console.log("delete robot")
}

</script>

<template>
  <div class="home-container">
    <el-row :gutter="20">
      <el-col :span="5">
        <BaseCard>
          <div class="flex align-center">
            <i
                class="i-Robot text-6xl text-purple-200"
            ></i>
            <div class="m-auto">
              <p class="text-gray-400">机器人数</p>
              <p class="text-xl text-primary">{{ total_robots }}</p>
            </div>
          </div>
        </BaseCard>
      </el-col>
      <el-col :span="5" :offset="1">
        <BaseCard>
          <div class="flex align-center">
            <i class="i-People-on-Cloud text-6xl text-purple-200"></i>
            <div class="m-auto">
              <p class="text-gray-400">租用人数</p>
              <p class="text-xl text-primary">10</p>
            </div>
          </div>
        </BaseCard>
      </el-col>
      <el-col :span="5" :offset="1">
        <BaseCard>
          <div class="flex align-center">
            <i
                class="i-Time-Bomb text-6xl text-purple-200"
            ></i>
            <div class="m-auto">
              <p class="text-gray-400">到期数量</p>
              <p class="text-xl text-primary">1</p>
            </div>
          </div>
        </BaseCard>
      </el-col>
    </el-row>
    <div class="robot-query-div">
      <el-form :inline="true" :model="formInline" class="form">
        <el-form-item label="租借人">
          <el-input v-model="formInline.user" clearable autocomplete="false"/>
        </el-form-item>
        <el-form-item label="到期日期">
          <el-date-picker
              v-model="formInline.date"
              type="datetime"
              format="YYYY/MM/DD hh:mm:ss"
              value-format="YYYY-MM-DD h:m:s a"
              clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="queryBtnLoading" :loading-icon="Eleme">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="warning" :loading="addBtnLoading" :loading-icon="Eleme" @click="addRobot">添加</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="robot-table">
      <el-table
          :data="robots"
          row-key="name"
          v-loading="loading"
          fit
          border
          stripe
          max-height="500px"
          style="width: 70%;border-radius: 3px;"
          :header-cell-style="{ backgroundColor: 'rgba(24,30,42,0.95)', color: '#ffffff', fontSize: '14px', textAlign: 'center', borderLeft: '0.5px #154480 solid', borderBottom: '0.5px #154480 solid' }"
          :cell-style="{backgroundColor: 'rgba(24,30,42,0.95)' ,color: 'rgb(189,196,239)', fontSize: '14px', borderBottom: '0.5px #143275 solid', borderLeft: '0.5px #143275 solid' }"
          :row-style="{ color: '#fff', fontSize: '14px' }"
      >
        <el-table-column prop="name" label="名称" align="center"/>
        <el-table-column prop="internal_ip" label="内部IP" align="center"/>
        <el-table-column prop="external_ip" label="外部IP" align="center"/>
        <el-table-column prop="borrowed_user" label="租用人" align="center"/>
        <el-table-column prop="borrowed_used" label="状态" align="center"/>
        <el-table-column prop="borrowed_date" label="租用日期" align="center"/>
        <el-table-column prop="sn" label="SN" align="center"/>
        <el-table-column fixed="right" label="操作" width="120">
          <template #default>
            <el-button link type="warning" size="small">续期</el-button>
            <el-popconfirm width="220"
                           confirm-button-text="Yes"
                           cancel-button-text="No"
                           @confirm="confirmDeleteRobot"
                           :icon="InfoFilled"
                           icon-color="#626AEF" title="Are you sure to delete this?">
              <template #reference>
                <el-button link type="danger" size="small"
                >删除</el-button
                >
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="robot-table-pagination">
      <el-pagination
          :small="true"
          :background="true"
          v-model:current-page="page"
          background layout="prev, pager, next"
          @current-change="fetchRobots"
          :page-size="5" :total="total_robots"/>
    </div>
    <pop :is-modal-open="isModalOpen" @close-modal="closeModal"></pop>
  </div>
</template>

<style lang="scss" scoped>
.home-container {
  top: 5vh;
  position: relative;
  left: 18%;
  right: 18%;
}

.robot-query-div {
  padding-top: 20px;
  margin-top: 50px;
  width: 70%;
  border-top-left-radius: 3px;
  border-top-right-radius: 3px;
  background-color: rgba(24, 30, 42, 0.95);

  .form {
    padding-left: 15px;
    color: #faf9f9;
  }
}

.robot-table {
  overflow-x: scroll;
}

.robot-table-pagination {
  padding-top: 1.2rem;
  display: flex;
  flex-direction: row;
  align-content: center;
  justify-content: space-between;
}
</style>