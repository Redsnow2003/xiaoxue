<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";
import type { PaginationProps } from "@pureadmin/table";
import { getSupplierFundLog } from "@/api/supplier";
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
    label: "ID",
    prop: "id"
  },
  {
    label: "供应商ID",
    prop: "supplier_id"
  },
  {
    label: "供应商名称",
    prop: "supplier_name"
  },
  {
    label: "操作",
    prop: "action",
    cellRenderer: ({ row }) => {
      if (row.action === 0) {
        return "余额加款";
      } else if (row.action === 1) {
        return "余额减款";
      } else {
        return "余额校正";
      }
    }
  },
  {
    label: "操作时间",
    prop: "time"
  },
  {
    label: "操作金额(元)",
    prop: "amount"
  },
  {
    label: "操作前余额(元)",
    prop: "before_amount"
  },
  {
    label: "操作后余额(元)",
    prop: "after_amount"
  },
  {
    label: "凭证图片",
    prop: "cert_pic",
    slot: "image"
  },
  {
    label: "备注",
    prop: "remark"
  },
  {
    label: "操作",
    fixed: "right",
    width: 210,
    slot: "operation"
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

async function onSearch() {
  var requestData = {
    supplier_id: supplier_id.value,
    currentPage: pagination.currentPage,
    pageSize: pagination.pageSize
  };
  const { data } = await getSupplierFundLog(requestData);
  console.log(data);
  dataList.value = data.list;
  pagination.total = data.total;
  pagination.pageSize = data.pageSize;
  pagination.currentPage = data.currentPage;
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
  >
    <template #image="{ row, index }">
      <el-image
        preview-teleported
        loading="lazy"
        :src="row.cert_pic"
        :initial-index="index"
        :preview-src-list="[row.cert_pic]"
        fit="cover"
        class="w-[100px] h-[100px]"
      />
    </template>
  </pure-table>
</template>
