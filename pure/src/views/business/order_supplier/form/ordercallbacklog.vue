<script setup lang="ts">
import { onMounted, ref } from "vue";
import { supplierOrderCallbackLog } from "@/api/order";
const props = defineProps({
  supplier_order_id: {
    type: Number,
    default: 0
  }
});

const ruleFormRef = ref();
const supplier_order_id = ref(props.supplier_order_id);
const dataList = ref([]);
const columns: TableColumnList = [
  {
    label: "供货单号",
    prop: "supplier_order_id"
  },
  {
    label: "回调方IP",
    prop: "callback_ip"
  },
  {
    label: "收到回调信息",
    prop: "callback_context"
  },
  {
    label: "收到回调时间",
    prop: "callback_time"
  },
  {
    label: "我方返回",
    prop: "response_context"
  },
  {
    label: "备注",
    prop: "remark"
  }
];
onMounted(async () => {
  onSearch();
});

function getRef() {
  return ruleFormRef.value;
}

async function onSearch() {
  var requestData = {
    supplier_order_id: supplier_order_id.value
  };
  const { data } = await supplierOrderCallbackLog(requestData);
  console.log(data);
  dataList.value = data;
}

defineExpose({ getRef });
</script>

<template>
  <pure-table
    ref="tableRef"
    row-key="id"
    align-whole="center"
    showOverflowTooltip
    table-layout="auto"
    adaptive
    :adaptiveConfig="{ offsetBottom: 108 }"
    :data="dataList"
    :columns="columns"
    :header-cell-style="{
      background: 'var(--el-fill-color-light)',
      color: 'var(--el-text-color-primary)'
    }"
  />
</template>
