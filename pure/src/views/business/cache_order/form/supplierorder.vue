<script setup lang="ts">
import { onMounted, ref, h } from "vue";
import { SupplierOrderStatusList } from "@/api/constdata";
import { getSupplierOrderList } from "@/api/order";
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
    label: "供货单号|上游单号",
    prop: "id",
    cellRenderer: ({ row }) => {
      const label = row.id + h("br") + h("hr") + row.up_id;
      return h("span", label);
    }
  },
  {
    label: "供货商ID",
    prop: "supplier_id"
  },
  {
    label: "供货商名称",
    prop: "supplier_name"
  },
  {
    label: "供货折扣|供货总价",
    prop: "supplier_discount",
    cellRenderer: ({ row }) => {
      const label =
        row.supplier_discount +
        h("br") +
        h("hr") +
        (row.supplier_discount * row.base_price * row.count).toFixed(2);
      return h("span", label);
    }
  },
  {
    label: "供货单状态",
    prop: "status",
    cellRenderer: ({ row }) => {
      const label =
        SupplierOrderStatusList.find(item => item.value === row.status)
          ?.label || "";
      return h("span", label);
    }
  },
  {
    label: "创建时间|完成时间",
    prop: "create_time",
    cellRenderer: ({ row }) => {
      const label = row.create_time + h("br") + h("hr") + row.finish_time;
      return h("span", label);
    }
  },
  {
    label: "上游信息",
    prop: "up_information"
  },
  {
    label: "备注",
    prop: "remark"
  },
  {
    label: "备用通道重试",
    prop: "is_backup",
    cellRenderer: ({ row }) => {
      const label = row.is_backup ? "是" : "否";
      return h("span", label);
    }
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
  const { data } = await getSupplierOrderList(requestData);
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
