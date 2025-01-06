<script setup lang="ts">
import { formRules } from "../utils/rule";
import { NewProdtuctFormProps, CategoryProps } from "../utils/types";
import {
  getProductCategoryList,
  getProductInformationList
} from "@/api/product";
import { BusinessTypeList, OperatorListAll, UnitList } from "@/api/constdata";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import Refresh from "@iconify-icons/ep/refresh";
import { h, onMounted, reactive, ref } from "vue";
import type { PaginationProps } from "@pureadmin/table";
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

const newProductForm = reactive({
  id: "",
  //业务类型
  type: "",
  //产品类别
  category: "",
  //产品名称
  name: "",
  //运营商
  operator: "",
  //面额
  price: "",
  //适用范围
  scope: ""
});
const columns: TableColumnList = [
  {
    label: "勾选列", // 如果需要表格多选，此处label必须设置
    type: "selection",
    fixed: "left",
    reserveSelection: true // 数据刷新后保留选项
  },
  {
    label: "业务类型",
    prop: "type",
    cellRenderer: ({ row }) => {
      const label =
        BusinessTypeList.find(item => item.value === row.type)?.label || "";
      return h("span", label);
    }
  },
  {
    label: "产品ID",
    prop: "id"
  },
  {
    label: "产品名称",
    prop: "name"
  },
  {
    label: "产品类别",
    prop: "category",
    cellRenderer: ({ row }) => {
      const label =
        newProductCategoryList.value.find(item => item.id === row.category)
          ?.category_name || "";
      return h("span", label);
    }
  },
  {
    label: "运营商",
    prop: "operator",
    cellRenderer: ({ row }) => {
      const label =
        OperatorListAll.find(item => item.value === row.operator)?.label || "";
      return h("span", label);
    }
  },
  {
    label: "适用范围",
    prop: "scope",
    cellRenderer: ({ row }) => {
      const label =
        OperatorListAll.find(item => item.value === row.operator)?.label || "";
      return h("span", label);
    }
  },
  {
    label: "面额",
    prop: "price"
  },
  {
    label: "单位",
    prop: "unit",
    cellRenderer: ({ row }) => {
      const label = UnitList.find(item => item.value === row.unit)?.label || "";
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
        },
        type: "number"
      })
  }
];
const pagination = reactive<PaginationProps>({
  total: 0,
  pageSize: 10,
  currentPage: 1,
  background: true
});
const selectedNum = ref(0);
const dataList = ref([]);
const loading = ref(true);
const newProductFormRef = ref();
const newProductCategoryList = ref([] as CategoryProps);
const newFormInline = ref(props.formInline);
function getRef() {
  return newProductFormRef.value;
}

async function onSearch() {
  console.log("onSearch");
  var requestData = {
    ...newProductForm,
    ...pagination
  };
  await getProductInformationList(requestData).then(response => {
    dataList.value = response.data.list;
    dataList.value.forEach((item: any) => {
      item.discount = 1;
    });
    pagination.total = response.data.total;
    loading.value = false;
  });
}

function resetForm() {
  newProductForm.id = "";
  newProductForm.type = "";
  newProductForm.category = "";
  newProductForm.name = "";
  newProductForm.operator = "";
  newProductForm.price = "";
  newProductForm.scope = "";
  onSearch();
}

function handleSizeChange(val: number) {
  pagination.pageSize = val;
  onSearch();
}

function handleCurrentChange(val: number) {
  pagination.currentPage = val;
  onSearch();
}

/** 当CheckBox选择项发生变化时会触发该事件 */
function handleSelectionChange(val: Array<any>) {
  selectedNum.value = val.length;
  newFormInline.value.product_list = val;
}

onMounted(async () => {
  var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
  const response2 = await getProductCategoryList(requestData);
  newProductCategoryList.value = response2.data.list;
  onSearch();
});

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="newProductFormRef"
    :model="newProductForm"
    :rules="formRules"
    label-width="82px"
  >
    <el-row :gutter="24">
      <el-col :span="8">
        <el-form-item label="业务类型" prop="business_type">
          <el-select
            v-model="newProductForm.type"
            placeholder="请选择业务类型"
            clearable
            class="!w-[180px]"
          >
            <el-option
              v-for="item in BusinessTypeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="产品类别" prop="product_category">
          <el-select
            v-model="newProductForm.category"
            placeholder="请选择产品类别"
            clearable
            class="!w-[180px]"
          >
            <el-option
              v-for="item in newProductCategoryList"
              :key="item.id"
              :label="item.category_name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="产品运营商" prop="operator">
          <el-select
            v-model="newProductForm.operator"
            placeholder="请选择产品运营商"
            clearable
            class="!w-[180px]"
          >
            <el-option
              v-for="item in OperatorListAll"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="24">
      <el-col :span="8">
        <el-form-item label="产品ID" prop="product_id">
          <el-input
            v-model="newProductForm.id"
            placeholder="请输入产品ID"
            clearable
            class="!w-[180px]"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="产品名称" prop="product_name">
          <el-input
            v-model="newProductForm.name"
            placeholder="请输入产品名称"
            clearable
            class="!w-[180px]"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item>
          <el-button
            type="primary"
            :icon="useRenderIcon('ri:search-line')"
            @click="onSearch"
          >
            搜索
          </el-button>
          <el-button :icon="useRenderIcon(Refresh)" @click="resetForm()">
            重置
          </el-button>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
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
    :pagination="{ ...pagination }"
    :header-cell-style="{
      background: 'var(--el-fill-color-light)',
      color: 'var(--el-text-color-primary)'
    }"
    @selection-change="handleSelectionChange"
    @page-size-change="handleSizeChange"
    @page-current-change="handleCurrentChange"
  />
</template>
