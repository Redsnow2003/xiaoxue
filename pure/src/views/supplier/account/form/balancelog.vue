<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";
import type { PaginationProps } from "@pureadmin/table";
import { getSupplierUpBalanceLog } from "@/api/supplier";
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
    label: "供货单号",
    prop: "supplier_id"
  },
  {
    label: "请求头",
    prop: "request_header"
  },
  {
    label: "请求地址",
    prop: "request_address"
  },
  {
    label: "请求参数",
    prop: "request_params"
  },
  {
    label: "请求时间",
    prop: "request_time"
  },
  {
    label: "上游返回",
    prop: "response_content"
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

function handleSizeChange(val: number) {
  pagination.pageSize = val;
  onSearch();
}
function handleCurrentChange(val: number) {
  pagination.currentPage = val;
  onSearch();
}

function onSearch() {
  const requestData = {
    supplier_id: supplier_id.value,
    currentPage: pagination.currentPage,
    pageSize: pagination.pageSize
  };
  getSupplierUpBalanceLog(requestData).then(res => {
    dataList.value = res.data.list;
    pagination.total = res.data.total;
  });
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
