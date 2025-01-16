<script setup lang="ts">
import { onMounted, ref } from "vue";
import { formRules } from "./utils/rule";
import { FormProps, AgentSimpleItem, SupplierSimpleItem } from "./utils/types";
import { getAgentSimpleList } from "@/api/agent";
import { fa } from "element-plus/es/locale/index.mjs";
import { getSupplierSimpleList } from "@/api/supplier";
const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    // ID
    id: 0,
    // 代理商ID
    agent_id: -1,
    // 代理商名称
    agent_name: "",
    // 供应商ID
    supplier_id: -1,
    // 供应商名称
    supplier_name: "",
    // 是否禁用
    disabled: false
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);
const agentItemLists = ref([] as AgentSimpleItem[]);
const supplierItemLists = ref([] as SupplierSimpleItem[]);
function getRef() {
  return ruleFormRef.value;
}
onMounted(async () => {
  // 动态获取产品类别列表
  const response = await getAgentSimpleList();
  agentItemLists.value = response.data;
  const supplierResponse = await getSupplierSimpleList();
  supplierItemLists.value = supplierResponse.data;
});
defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-form-item label="代理商" prop="agent_id">
      <el-select
        v-model="newFormInline.agent_id"
        placeholder="请选择代理商"
        clearable
        :disabled="newFormInline.disabled"
      >
        <el-option
          v-for="item in agentItemLists"
          :key="item.id"
          :label="item.id + ' ' + item.name"
          :value="item.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="供应商" prop="supplier_id">
      <el-select
        v-model="newFormInline.supplier_id"
        placeholder="请选择供应商"
        clearable
      >
        <el-option
          v-for="item in agentItemLists"
          :key="item.id"
          :label="item.id + ' ' + item.name"
          :value="item.id"
        />
      </el-select>
    </el-form-item>
  </el-form>
</template>
