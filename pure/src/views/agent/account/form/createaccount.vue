<script setup lang="ts">
import { onMounted, ref } from "vue";
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

onMounted(async () => {});
</script>

<template>
  <el-form label-width="230px" />
</template>
