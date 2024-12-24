<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";
import type { PaginationProps } from "@pureadmin/table";
const props = defineProps({
  supplier_id: {
    type: Number,
    default: 0
  }
});

const ruleFormRef = ref();
const supplier_id = ref(props.supplier_id);
const dataList = ref([]);
const pagination = reactive<PaginationProps>({
  total: 0,
  pageSize: 10,
  currentPage: 1,
  background: true
});
const columns: TableColumnList = [
  {
    label: "勾选列", // 如果需要表格多选，此处label必须设置
    type: "selection",
    fixed: "left",
    reserveSelection: true // 数据刷新后保留选项
  },
  {
    label: "供应商ID",
    prop: "id"
  },
  {
    label: "供应商名称",
    prop: "name"
  },
  {
    label: "我方平台账户余额(元)",
    prop: "our_balance"
  },
  {
    label: "上游平台账户余额(元)",
    prop: "up_balance"
  },
  {
    label: "上游平台账户余额更新时间",
    prop: "up_balance_update_time"
  },
  {
    label: "上游模板配置",
    prop: "up_template"
  },
  {
    label: "备注",
    prop: "remark"
  },
  {
    label: "状态",
    prop: "status"
  },
  {
    label: "状态信息",
    prop: "status_info"
  },
  {
    label: "操作",
    fixed: "right",
    width: 210,
    slot: "operation"
  }
];
onMounted(async () => {});

function getRef() {
  return ruleFormRef.value;
}

function handleSizeChange(val: number) {
  pagination.pageSize = val;
  onSearch();
}
function handleCurrentChange(val: number) {
  pagination.currentPage = val;
  onSearch();
}

function onSearch() {
  console.log("onSearch");
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
    :pagination="pagination"
    :header-cell-style="{
      background: 'var(--el-fill-color-light)',
      color: 'var(--el-text-color-primary)'
    }"
    @page-size-change="handleSizeChange"
    @page-current-change="handleCurrentChange"
  />
</template>
