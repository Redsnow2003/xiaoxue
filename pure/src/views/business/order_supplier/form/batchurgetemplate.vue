<script setup lang="ts">
import { ref } from "vue";
const props = defineProps({
  ids: {
    type: Array,
    default: () => []
  }
});

const ids = ref(props.ids[0]);
</script>

<template>
  <el-form label-width="82px">
    <el-form-item label="订单号" prop="ids"
  /></el-form>
</template>
