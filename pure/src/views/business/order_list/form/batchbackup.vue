<script setup lang="ts">
import { ref } from "vue";
import { BatchBackupFormProps } from "../utils/types";
const props = withDefaults(defineProps<BatchBackupFormProps>(), {
  formInline: () => ({
    /** 次数 */
    count: 1,
    /** 时间间隔 */
    interval: 10
  })
});

const newFormInline = ref(props.formInline);
</script>

<template>
  <el-form label-width="82px">
    <el-form-item label="重试次数" prop="count">
      <el-input
        v-model="newFormInline.count"
        clearable
        type="number"
        class="!w-[380px]"
        placeholder="请输入重试次数"
      />
    </el-form-item>
    <el-form-item label="时间间隔" prop="interval">
      <el-input
        v-model="newFormInline.interval"
        clearable
        class="!w-[380px]"
        type="number"
        placeholder="请输入时间间隔"
      />
    </el-form-item>
  </el-form>
</template>
