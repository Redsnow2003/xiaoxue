<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "../utils/rule";
import { StatusItem } from "../utils/types";

const props = withDefaults(defineProps<StatusItem>(), {
  status: 1
});

const ruleFormRef = ref();
const newFormInline = ref(props.status);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form ref="ruleFormRef" :rules="formRules" label-width="82px">
    <el-form-item label="操作状态" prop="type">
      <el-select v-model="newFormInline" clearable disabled>
        <el-option key="0" label="维护" value="0" />
        <el-option key="1" label="上架" value="1" />
      </el-select>
    </el-form-item>
  </el-form>
</template>
