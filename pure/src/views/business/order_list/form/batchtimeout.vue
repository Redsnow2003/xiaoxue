<script setup lang="ts">
import { ref } from "vue";
import { BatchTimeoutFormProps } from "../utils/types";
import { BatchTimeoutType } from "@/api/constdata";
const props = withDefaults(defineProps<BatchTimeoutFormProps>(), {
  formInline: () => ({
    type: 0,
    seconds: 0,
    time: "",
    create_after: 0
  })
});

const newFormInline = ref(props.formInline);
</script>

<template>
  <el-form label-width="82px">
    <el-form-item label="方式" prop="type">
      <el-select
        v-model="newFormInline.type"
        placeholder="请选择订单状态"
        clearable
        class="!w-[380px]"
      >
        <el-option
          v-for="item in BatchTimeoutType"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item v-if="newFormInline.type === 1" label="秒数" prop="seconds">
      <el-input
        v-model="newFormInline.seconds"
        clearable
        type="number"
        class="!w-[380px]"
        placeholder="请输入秒数"
      />
    </el-form-item>
    <el-form-item v-if="newFormInline.type === 2" label="时间点" prop="time">
      <el-date-picker
        v-model="newFormInline.time"
        class="!w-[380px]"
        type="datetime"
      />
    </el-form-item>
    <el-form-item
      v-if="newFormInline.type === 3"
      label="秒数"
      prop="create_after"
    >
      <el-input
        v-model="newFormInline.create_after"
        clearable
        type="number"
        class="!w-[380px]"
        placeholder="请输入秒数"
      />
    </el-form-item>
  </el-form>
</template>
