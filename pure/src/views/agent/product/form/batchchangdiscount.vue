<script setup lang="ts">
import {
  NewProdtuctFormProps,
  CategoryProps,
  AgentProcuctItem
} from "../utils/types";
import { getProductCategoryList } from "@/api/product";
import { getAllAgentProductList } from "@/api/agent";
import { h, onMounted, ref } from "vue";
import input from "element-plus/es/components/input/index.mjs";
const props = withDefaults(defineProps<NewProdtuctFormProps>(), {
  formInline: () => ({
    //代理商id
    agent_id: -1,
    agent_name: "",
    //产品列表
    product_list: []
  })
});

const columns: TableColumnList = [
  {
    label: "ID",
    prop: "id"
  },
  {
    label: "代理商ID",
    prop: "agent_id"
  },
  {
    label: "代理商名称",
    prop: "agent_name"
  },
  {
    label: "产品ID",
    prop: "product_id"
  },
  {
    label: "产品名称",
    prop: "product_name"
  },
  {
    label: "产品类别",
    prop: "category",
    cellRenderer: ({ row }) => {
      const label =
        newProductCategoryList.value.find(
          item => item.id === row.product_category
        )?.category_name || "";
      return h("span", label);
    }
  },
  {
    label: "基础价(元)",
    prop: "base_price"
  },
  {
    label: "折扣",
    prop: "discount",
    cellRenderer: ({ row }) =>
      h(input, {
        modelValue: row.discount,
        "onUpdate:modelValue": (val: string) => {
          row.discount = val;
          newFormInline.value.product_list.find(
            item => item.id === row.id
          ).discount = Number(val);
        },
        type: "number"
      })
  }
];
const dataList = ref([]);
const loading = ref(true);
const newFormInline = ref(props.formInline);
const newProductCategoryList = ref([] as CategoryProps);
async function onSearch() {
  var requestData = { agent_id: props.formInline.agent_id };
  await getAllAgentProductList(requestData).then(response => {
    dataList.value = response.data;
    newFormInline.value.product_list = [];
    dataList.value.forEach((item: any) => {
      newFormInline.value.product_list.push({
        id: item.id,
        discount: item.discount
      } as AgentProcuctItem);
    });
    loading.value = false;
  });
}

onMounted(async () => {
  var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
  const response2 = await getProductCategoryList(requestData);
  newProductCategoryList.value = response2.data.list;
  onSearch();
});
</script>

<template>
  <pure-table
    ref="tableRef"
    adaptive
    :adaptiveConfig="{ offsetBottom: 45 }"
    align-whole="center"
    row-key="id"
    showOverflowTooltip
    table-layout="auto"
    default-expand-all
    :loading="loading"
    :data="dataList"
    :columns="columns"
    :header-cell-style="{
      background: 'var(--el-fill-color-light)',
      color: 'var(--el-text-color-primary)'
    }"
  />
</template>
