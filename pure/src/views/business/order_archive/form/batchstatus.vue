<script setup lang="ts">
import { onMounted, ref } from "vue";
import { formRules } from "../utils/rule";
import { BatchStatusFormProps } from "../utils/types";
import { getDeptList } from "@/api/system";
import { handleTree } from "@/utils/tree";
import { OrderStatusList } from "@/api/constdata";

const props = withDefaults(defineProps<BatchStatusFormProps>(), {
  formInline: () => ({
    /** 状态 */
    status: 0,
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
    <el-form-item label="订单状态" prop="status">
      <el-select
        v-model="newFormInline.status"
        placeholder="请选择订单状态"
        clearable
        class="!w-[380px]"
      >
        <el-option
          v-for="item in OrderStatusList"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="备注" prop="remark">
      <el-input
        v-model="newFormInline.remark"
        clearable
        class="!w-[380px]"
        placeholder="请输入订单备注"
      />
    </el-form-item>
  </el-form>
</template>
