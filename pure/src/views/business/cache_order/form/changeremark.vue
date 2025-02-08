<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "../utils/rule";
import { ChangeRemarkFormProps } from "../utils/types";
const props = withDefaults(defineProps<ChangeRemarkFormProps>(), {
  formInline: () => ({
    /** 订单ID */
    order_id: 0,
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
    <el-form-item label="供应商ID" prop="order_id">
      <el-input
        v-model="newFormInline.order_id"
        clearable
        placeholder="供应商ID"
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
  </el-form>
</template>
