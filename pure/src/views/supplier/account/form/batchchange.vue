<script setup lang="ts">
import { ref } from "vue";
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
    up_template: "",
    template_json: "",
    status: 1,
    status_info: "",
    remark: ""
  })
});

const newFormInline = ref(props.formInline);
</script>

<template>
  <el-form label-width="82px">
    <el-form-item label="操作状态" prop="type">
      <el-select v-model="newFormInline.status" clearable>
        <el-option :key="0" label="维护" :value="0" />
        <el-option :key="1" label="上架" :value="1" />
      </el-select>
    </el-form-item>
  </el-form>
</template>
