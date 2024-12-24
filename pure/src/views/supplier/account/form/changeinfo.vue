<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "../utils/rule";
import { FormProps } from "../utils/types";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    id: -1,
    name: "",
    nickname: "",
    dept: -1,
    phone: "",
    email: "",
    our_balance: 0,
    up_balance: 0,
    up_balance_update_time: "",
    up_template: -1,
    status: 1,
    status_info: "",
    remark: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

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
    label-width="82px"
  >
    <el-form-item label="供应商ID" prop="id">
      <el-input
        v-model="newFormInline.id"
        clearable
        placeholder="请输入供应商ID"
        disabled
      />
    </el-form-item>
    <el-form-item label="供应商名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入供应商名称"
        disabled
      />
    </el-form-item>
    <el-form-item label="产品名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入产品名称"
      />
    </el-form-item>
    <el-form-item label="备注" prop="remark">
      <el-input
        v-model="newFormInline.remark"
        clearable
        placeholder="请输入备注"
      />
    </el-form-item>
    <el-form-item label="状态" prop="status">
      <el-select v-model="newFormInline.status" clearable>
        <el-option key="0" label="维护" value="0" />
        <el-option key="1" label="上架" value="1" />
      </el-select>
    </el-form-item>
  </el-form>
</template>
