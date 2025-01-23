<script setup lang="ts">
import { onMounted, ref, h } from "vue";
import { getAgentFundLog } from "@/api/agent";
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
    label: "ID",
    prop: "id"
  },
  {
    label: "订单ID",
    prop: "order_id"
  },
  {
    label: "重试次数",
    prop: "retry_count"
  },
  {
    label: "已执行次数",
    prop: "execute_count"
  },
  {
    label: "失败重试间隔",
    prop: "interavl"
  },
  {
    label: "下次重试时间",
    prop: "retry_time"
  },
  {
    label: "状态",
    prop: "status"
  },
  {
    label: "备注",
    prop: "remark"
  },
  {
    label: "操作",
    fixed: "right",
    width: 100,
    slot: "operation"
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
  const { data } = await getAgentFundLog(requestData);
  console.log(data);
  dataList.value = data.list;
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
  >
    <template #operation="{ row }">
      <el-button class="reset-margin" link type="primary" @click="row">
        删除
      </el-button>
    </template>
  </pure-table>
</template>
