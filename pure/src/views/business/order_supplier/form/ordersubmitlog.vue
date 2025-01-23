<script setup lang="ts">
import { onMounted, ref, h } from "vue";
import { agentOrderSubmit } from "@/api/order";
const props = defineProps({
  order_id: {
    type: Number,
    default: 0
  }
});

const ruleFormRef = ref();
const order_id = ref(props.order_id);
const dataList = ref([]);
const columns: TableColumnList = [
  {
    label: "订单号",
    prop: "order_id"
  },
  {
    label: "请求头",
    prop: "request_header"
  },
  {
    label: "请求参数",
    prop: "request_params"
  },
  {
    label: "请求IP",
    prop: "request_ip"
  },
  {
    label: "请求时间",
    prop: "request_time"
  },
  {
    label: "我方返回",
    prop: "response_context"
  },
  {
    label: "返回时间",
    prop: "response_time"
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
    order_id: order_id.value
  };
  const { data } = await agentOrderSubmit(requestData);
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
