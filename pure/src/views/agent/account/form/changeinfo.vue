<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "../utils/rule";
import { FormProps } from "../utils/types";

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({
    /** 代理商ID */
    id: 0,
    /** 代理商名称 */
    name: "",
    /** 代理商简称 */
    nickname: "",
    /** 代理商部门 */
    dept: 0,
    /** 代理商电话 */
    phone: "",
    /** 代理商邮箱 */
    email: "",
    /** 密钥 */
    secret_key: "",
    /** 通知地址 */
    notification_address: "",
    /** 通知方式 */
    notification_method: 0,
    /** 客户 */
    customer: "",
    /** 状态 */
    status: 0,
    /** 资金余额 */
    fund_balance: 0,
    /** 授信余额 */
    credit_balance: 0,
    /** 冻结金额 */
    frozen_amount: 0,
    /** 缓存可用金额 */
    cache_amount: 0,
    /** 备注 */
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
    <el-form-item label="代理商ID" prop="id">
      <el-input
        v-model="newFormInline.id"
        clearable
        placeholder="请输入代理商ID"
        disabled
      />
    </el-form-item>
    <el-form-item label="代理商名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入代理商名称"
        disabled
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
        <el-option :key="0" label="维护" :value="0" />
        <el-option :key="1" label="上架" :value="1" />
      </el-select>
    </el-form-item>
  </el-form>
</template>
