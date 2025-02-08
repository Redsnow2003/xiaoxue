<script setup lang="ts">
import { onMounted, ref, h } from "vue";
import { getAgentFundLog } from "@/api/agent";
import { NotifyStatusList } from "@/api/constdata";
import { agentOrderNoticeLog } from "@/api/order";
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
    label: "发送日志",
    prop: "send_log"
  },
  {
    label: "通知地址",
    prop: "send_address"
  },
  {
    label: "代理商返回",
    prop: "agent_return"
  },
  {
    label: "通知状态",
    prop: "notify_status",
    formatter: (row: any) => {
      return NotifyStatusList.find(item => item.value === row.notify_status)
        ?.label;
    }
  },
  {
    label: "通知时间",
    prop: "notify_time"
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
  const { data } = await agentOrderNoticeLog(requestData);
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
