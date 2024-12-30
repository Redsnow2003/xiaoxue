<script setup lang="ts">
import { onMounted, ref } from "vue";
import { formRules } from "./utils/rule";
import { FormProps, CategoryProps } from "./utils/types";
import {
  BusinessTypeList,
  OperatorListTelecom,
  OperatorListTelecom2,
  ProvinceList,
  UnitList
} from "@/api/constdata";
import { getProductCategoryList } from "@/api/product";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    id: -1,
    type: -1,
    category: -1,
    name: "",
    operator: -1,
    scope: -1,
    unit: -1,
    price: 0,
    base_price: 0,
    sale_inventory: 0,
    sale_inventory_amount: 0,
    api_limit: 0,
    remark: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);
const categoryList = ref<CategoryProps>([]);

onMounted(async () => {
  // 动态获取产品类别列表
  var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
  const response = await getProductCategoryList(requestData);
  categoryList.value = response.data.list;
});

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="120px"
  >
    <el-form-item label="业务类型" prop="type">
      <el-select v-model="newFormInline.type" clearable disabled>
        <el-option
          v-for="item in BusinessTypeList"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="产品类别" prop="category">
      <el-select v-model="newFormInline.category" clearable>
        <el-option
          v-for="item in categoryList"
          :key="item.id"
          :label="item.category_name"
          :value="item.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="产品名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入产品名称"
      />
    </el-form-item>
    <el-form-item label="运营商" prop="operator">
      <el-select v-model="newFormInline.operator" clearable>
        <el-option
          v-for="item in OperatorListTelecom"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="单位" prop="unit">
      <el-select v-model="newFormInline.unit" clearable disabled>
        <el-option
          v-for="item in UnitList"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="基础价格(元)" prop="base_price">
      <el-input
        v-model="newFormInline.base_price"
        clearable
        placeholder="请输入产品价格"
        :type="'number'"
      />
    </el-form-item>
    <el-form-item label="API提单限制数量" prop="api_limit">
      <el-input
        v-model="newFormInline.api_limit"
        clearable
        placeholder="请输入限制数量"
        :type="'number'"
      />
    </el-form-item>
    <el-form-item label="备注" prop="remark">
      <el-input
        v-model="newFormInline.remark"
        clearable
        placeholder="请输入产品备注"
      />
    </el-form-item>
  </el-form>
</template>
