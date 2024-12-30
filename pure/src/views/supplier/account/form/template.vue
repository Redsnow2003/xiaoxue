<script setup lang="ts">
import { onMounted, ref } from "vue";
import { FormProps, TemplateInfoItem } from "../utils/types";
import {
  getSupplierTemplateNameList,
  getSupplierTemplateJson
} from "@/api/supplier";
import "plus-pro-components/es/components/form/style/css";
import {
  type PlusColumn,
  type FieldValues,
  PlusForm,
  FormItemValueType
} from "plus-pro-components";

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
const templateNameLists = ref([]);
const newFormInline = ref(props.formInline);
const valueList = ref<FieldValues>({});
const columns = ref<PlusColumn[]>([]);
const templateInfoList = ref<TemplateInfoItem[]>([]);
const remarkString = ref("");
const defaultTemplate = ref("");
const defaultTemplateJson = ref("");
const defaultValueList = ref<FieldValues>({});
onMounted(async () => {
  // 动态获取产品类别列表
  const response2 = await getSupplierTemplateNameList();
  templateNameLists.value = response2.data;
  defaultTemplate.value = newFormInline.value.up_template;
  defaultTemplateJson.value = newFormInline.value.template_json;
  changeTemplate();
});

async function changeTemplate() {
  var data = {
    template_name: newFormInline.value.up_template
  };
  valueList.value = {};
  await getSupplierTemplateJson(data).then(res => {
    templateInfoList.value = res.data.list as TemplateInfoItem[];
    columns.value = [];
    defaultValueList.value = {};
    // 以templateInfoList的index作为key进行排序
    templateInfoList.value.sort((a, b) => a.index - b.index);

    templateInfoList.value.forEach(item => {
      if (item.prop === "remark") {
        remarkString.value = item.value;
        return;
      }
      defaultValueList.value[item.prop] = item.value;
      columns.value.push({
        prop: item.prop,
        label: item.label,
        valueType: item.type as FormItemValueType,
        options: item.options
      });
    });
  });
  console.log(newFormInline.value.up_template);
  console.log(defaultTemplateJson.value);
  if (newFormInline.value.up_template == defaultTemplate.value) {
    if (defaultTemplateJson.value === "") {
      return;
    }
    valueList.value = JSON.parse(defaultTemplateJson.value);
    newFormInline.value.template_json = defaultTemplateJson.value;
  }
}
function handleColumnChange() {
  // 将valueList转成字符串
  var str = JSON.stringify(valueList.value);
  newFormInline.value.template_json = str;
}

function loadDefualtValue() {
  valueList.value = { ...defaultValueList.value };
}
</script>

<template>
  <el-button
    type="primary"
    style="margin-bottom: 20px"
    @click="loadDefualtValue"
    >加载默认值</el-button
  >
  <el-form label-width="230px">
    <el-form-item label="模板名称:" prop="up_template">
      <el-select
        v-model="newFormInline.up_template"
        clearable
        @change="changeTemplate"
      >
        <el-option
          v-for="item in templateNameLists"
          :key="item"
          :label="item"
          :value="item"
        />
      </el-select>
    </el-form-item>
    <PlusForm
      v-model="valueList"
      :columns="columns"
      label-position="right"
      labelWidth="230px"
      :hasFooter="false"
      @change="handleColumnChange"
    />
    <el-form-item label="注意事项:">
      <div style="color: red">
        <h4>{{ remarkString }}</h4>
      </div>
    </el-form-item>
  </el-form>
</template>
