<script setup lang="ts">
import { ref } from "vue";
import { FundOperationTypeList2 } from "@/api/constdata";
import { formRules } from "../utils/rule";
import { ChangeFundFormProps } from "../utils/types";
import AddFill from "@iconify-icons/ri/add-circle-line";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
const props = withDefaults(defineProps<ChangeFundFormProps>(), {
  formInline: () => ({
    agent_id: "", // 供应商ID
    agent_name: "", // 供应商名称
    fund_action: 0, // 余额操作
    amount: 0, // 金额
    confirm_amount: 0, // 确认金额
    file: "", // 上传图片
    remark: "" // 备注
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);
const fileList = ref([]);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });

function beforeUpload(file) {
  const isJPG = file.type === "image/jpeg";
  if (!isJPG) {
    this.$message.error("上传图片只能是 JPG 格式!");
  }
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isLt2M) {
    this.$message.error("上传图片大小不能超过 2MB!");
  }
  return isJPG && isLt2M;
}

function handleAvatarSuccess(response: any, files: string | any[]) {
  console.log(response, files);
  if (files.length === 0) {
    return;
  }
  const file = files[0];
  const reader = new FileReader();
  reader.onload = (e: ProgressEvent<FileReader>) => {
    newFormInline.value.file = String(e.target.result);
  };
  reader.readAsDataURL(file.raw);
}
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-form-item label="供应商ID" prop="agent_id">
      <el-input
        v-model="newFormInline.agent_id"
        clearable
        placeholder="供应商ID"
        disabled
      />
    </el-form-item>
    <el-form-item label="供应商名称" prop="agent_name">
      <el-input
        v-model="newFormInline.agent_name"
        clearable
        placeholder="供应商名称"
        disabled
      />
    </el-form-item>
    <el-form-item label="余额操作" prop="fund_action">
      <el-select v-model="newFormInline.fund_action" clearable>
        <el-option
          v-for="item in FundOperationTypeList2"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="金额" prop="amount">
      <el-input
        v-model="newFormInline.amount"
        clearable
        placeholder="请输入金额"
        type="number"
      />
    </el-form-item>
    <el-form-item label="确认金额" prop="confirm_amount">
      <el-input
        v-model="newFormInline.confirm_amount"
        clearable
        placeholder="请确认金额"
        type="number"
      />
    </el-form-item>
    <el-form-item label="上传图片" prop="upload">
      <el-upload
        action="#"
        list-type="picture-card"
        :file-list="fileList"
        :show-file-list="true"
        :multiple="false"
        :auto-upload="false"
        :on-change="handleAvatarSuccess"
        :limit="1"
        :before-upload="beforeUpload"
      >
        <el-icon :icon="useRenderIcon(AddFill)" />
      </el-upload>
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
