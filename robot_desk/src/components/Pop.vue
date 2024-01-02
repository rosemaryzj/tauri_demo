<script setup>
import {Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot,} from '@headlessui/vue'
import {ref} from "vue";

const props = defineProps({
  isModalOpen: {
    type: Boolean,
    required: true,
    default: true
  }
})


const robotInfo = ref({
  sn: "",
  name: "",
  internal_ip: "",
  external_ip: "",
  borrowed_date: ""
})
const emits = defineEmits(["closeModal"])

function closeModal() {
  console.log(robotInfo.value);
  robotInfo.value = robotInfo.value = {
    sn: "",
    name: "",
    internal_ip: "",
    external_ip: "",
    borrowed_date: ""
  };
  emits("closeModal", false);
}
</script>

<template>
  <TransitionRoot appear :show="props.isModalOpen" as="template">
    <Dialog as="div" @close="closeModal" class="relative z-10">
      <TransitionChild
          as="template"
          enter="duration-300 ease-out"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="duration-200 ease-in"
          leave-from="opacity-100"
          leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-black/25"/>
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div
            class="flex min-h-full items-center justify-center p-4 text-center"
        >
          <TransitionChild
              as="template"
              enter="duration-300 ease-out"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
          >
            <DialogPanel
                class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all"
            >
              <DialogTitle
                  as="h3"
                  class="text-lg font-medium leading-6 text-gray-900"
              >
                填写机器人信息
              </DialogTitle>
              <div class="mt-2">
                <el-form v-model="robotInfo" :label-width="100">
                  <el-form-item
                      prop="name"
                      label="名称"
                      :rules="{
                      required: true,
                      message: '名称不能为空',
                      trigger: 'blur',
      }"
                  >
                    <el-input v-model="robotInfo.name" clearable/>
                  </el-form-item>
                  <el-form-item
                      prop="internal_ip"
                      label="内部地址"
                      :rules="{
                      required: true,
                      message: '内部地址不能为空',
                      trigger: 'blur'}"
                  >
                    <el-input v-model="robotInfo.internal_ip" clearable/>
                  </el-form-item>
                  <el-form-item
                      prop="external_ip"
                      label="外部地址"
                      :rules="{
                      required: true,
                      message: '外部地址不能为空',
                      trigger: 'blur'}"
                  >
                    <el-input v-model="robotInfo.external_ip"/>
                  </el-form-item>
                  <el-form-item
                      label="SN"
                      prop="sn"
                      :rules="{
                      required: true,
                      message: 'SN不能为空',
                      trigger: 'blur'}"
                  >
                    <el-input v-model="robotInfo.sn"/>
                  </el-form-item>
                  <el-form-item
                      prop="borrowed_data"
                      label="借用日期"
                      :rules="{
                      required: true,
                      message: '借用日期不能为空',
                      trigger: 'blur'}"
                  >
                    <el-input v-model="robotInfo.borrowed_date"/>
                  </el-form-item>
                </el-form>
              </div>

              <div class="mt-4">
                <button
                    type="button"
                    class="inline-flex justify-center rounded-md border border-transparent bg-orange-100 px-4 py-2 text-sm text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="closeModal"
                >
                  确认
                </button>
                <button
                    type="button"
                    class="ml-4 inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                    @click="closeModal"
                >
                  关闭
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>

<style scoped lang="scss">

</style>