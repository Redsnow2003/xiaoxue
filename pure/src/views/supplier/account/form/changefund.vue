<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "../utils/rule";
import { ChangeFundFormProps } from "../utils/types";

const props = withDefaults(defineProps<ChangeFundFormProps>(), {
  formInline: () => ({
    supplierId: "", // 供应商ID
    supplierName: "", // 供应商名称
    fundAction: "", // 余额操作
    amount: 0, // 金额
    confirmAmount: 0, // 确认金额
    fileList: [], // 上传图片
    remark: "" // 备注
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

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
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-form-item label="供应商ID" prop="supplierId">
      <el-input
        v-model="newFormInline.supplierId"
        clearable
        placeholder="供应商ID"
        readonly
      />
    </el-form-item>
    <el-form-item label="供应商名称" prop="supplierName">
      <el-input
        v-model="newFormInline.supplierName"
        clearable
        placeholder="供应商名称"
        readonly
      />
    </el-form-item>
    <el-form-item label="余额操作" prop="fundAction">
      <el-select v-model="newFormInline.fundAction" clearable>
        <el-option label="余额加款" value="add" />
        <el-option label="余额减款" value="subtract" />
        <el-option label="余额校正" value="adjust" />
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
    <el-form-item label="确认金额" prop="confirmAmount">
      <el-input
        v-model="newFormInline.confirmAmount"
        clearable
        placeholder="请确认金额"
        type="number"
      />
    </el-form-item>
    <el-form-item label="上传图片" prop="upload">
      <el-upload
        action="#"
        list-type="picture-card"
        :file-list="newFormInline.fileList"
        :before-upload="beforeUpload"
      >
        <i class="el-icon-plus" />
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
