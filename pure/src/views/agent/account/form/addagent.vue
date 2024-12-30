<script setup lang="ts">
import { onMounted, ref } from "vue";
import { formRules } from "../utils/rule";
import { FormProps } from "../utils/types";
import { getDeptList } from "@/api/system";
import { handleTree } from "@/utils/tree";

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
const deptTree = ref([]);
function getRef() {
  return ruleFormRef.value;
}

onMounted(async () => {
  // 动态获取部门列表
  const response = await getDeptList();
  deptTree.value = handleTree(response.data);
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
    <el-form-item label="代理商名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入代理商名称"
      />
    </el-form-item>
    <el-form-item label="归属部门">
      <el-cascader
        v-model="newFormInline.dept"
        class="w-full"
        :options="deptTree"
        :props="{
          value: 'id',
          label: 'name',
          emitPath: false,
          checkStrictly: true
        }"
        clearable
        filterable
        placeholder="请选择归属部门"
      >
        <template #default="{ node, data }">
          <span>{{ data.name }}</span>
          <span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
        </template>
      </el-cascader>
    </el-form-item>
    <el-form-item label="手机号码" prop="phone">
      <el-input
        v-model="newFormInline.phone"
        clearable
        placeholder="请输入手机号码"
      />
    </el-form-item>
    <el-form-item label="电子邮件" prop="email">
      <el-input
        v-model="newFormInline.email"
        clearable
        placeholder="请输入电子邮件"
      />
    </el-form-item>
    <el-form-item label="用户昵称" prop="nickname">
      <el-input
        v-model="newFormInline.nickname"
        clearable
        placeholder="请输入用户昵称"
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
