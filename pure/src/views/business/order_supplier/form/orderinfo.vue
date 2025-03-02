<script setup lang="ts">
import { onMounted, ref, h } from "vue";
import { OrderStatusList } from "@/api/constdata";
import { getOrderList } from "@/api/order";
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
    label: "订单号|下游单号",
    prop: "id",
    cellRenderer: ({ row }) => {
      return h("span", [
        h("div", row.id),
        h("hr", { style: "border-color: lightgray;" }),
        h("div", row.down_id)
      ]);
    }
  },
  {
    label: "代理商",
    prop: "agent_id",
    cellRenderer: ({ row }) => {
      return h("span", [
        h("div", row.agent_id),
        h("hr", { style: "border-color: lightgray;" }),
        h("div", row.agent_name)
      ]);
    }
  },
  {
    label: "代理折扣|代理总价",
    prop: "agent_discount",
    cellRenderer: ({ row }) => {
      return h("span", [
        h("div", row.agent_discount),
        h("hr", { style: "border-color: lightgray;" }),
        h("div", (row.agent_discount * row.base_price * row.count).toFixed(2))
      ]);
    }
  },
  {
    label: "订单状态",
    prop: "status",
    cellRenderer: ({ row }) => {
      const label =
        OrderStatusList.find(item => item.value === row.status)?.label || "";
      return h("span", label);
    }
  },
  {
    label: "订单状态",
    prop: "status",
    cellRenderer: ({ row }) => {
      const label = OrderStatusList.find(
        item => item.value === row.status
      )?.label;
      return h("span", label);
    }
  },
  {
    label: "创建时间|完成时间",
    prop: "create_time",
    cellRenderer: ({ row }) => {
      return h("span", [
        h("div", row.create_time),
        h("hr", { style: "border-color: lightgray;" }),
        h("div", row.finish_time)
      ]);
    }
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
    id: String(order_id.value),
    currentPage: 1,
    pageSize: 100
  };
  const { data } = await getOrderList(requestData);
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
  />
</template>
