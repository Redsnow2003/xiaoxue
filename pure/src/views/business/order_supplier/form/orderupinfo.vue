<script setup lang="ts">
import { onMounted, ref, h } from "vue";
import { queryUpOrderInfo } from "@/api/order";
const props = defineProps({
  up_id: {
    type: Number,
    default: 0
  },
  supplier_id: {
    type: Number,
    default: 0
  }
});

const up_id = ref(props.up_id);
const supplier_id = ref(props.supplier_id);
const up_status = ref("");
const up_info = ref("");
onMounted(async () => {
  onSearch();
});

async function onSearch() {
  var requestData = {
    up_id: up_id.value,
    supplier_id: supplier_id.value
  };
  const { data } = await queryUpOrderInfo(requestData);
}
</script>

<template>
  <el-form label-width="82px" min-height="500px">
    <el-form-item label="状态" prop="up_status">
      <el-input v-model="up_status" clearable disabled />
    </el-form-item>
    <el-form-item label="信息" prop="up_info">
      <el-input v-model="up_info" clearable disabled />
    </el-form-item>
  </el-form>
</template>
