<script setup lang="ts">
import { onMounted, ref } from "vue";
import { formRules } from "./utils/rule";
import { FormProps, AgentSimpleItem } from "./utils/types";
import { getAgentSimpleList } from "@/api/agent";
import { fa } from "element-plus/es/locale/index.mjs";
const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    // ID
    id: 0,
    //供应商id
    agent_id: -1,
    //产品列表
    agent_name: "",
    ip: "",
    ip_location: "",
    remark: "",
    disabled: false
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);
const agentItemLists = ref([] as AgentSimpleItem[]);
console.log("newFormInline", newFormInline);
function getRef() {
  return ruleFormRef.value;
}
onMounted(async () => {
  // 动态获取产品类别列表
  const response = await getAgentSimpleList();
  agentItemLists.value = response.data;
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
    <el-form-item label="IP地址" prop="ip">
      <el-input
        v-model="newFormInline.ip"
        clearable
        placeholder="请输入IP地址"
      />
    </el-form-item>
    <el-form-item label="备注" prop="remark">
      <el-input
        v-model="newFormInline.remark"
        clearable
        placeholder="请输入备注"
      />
    </el-form-item>
  </el-form>
</template>
