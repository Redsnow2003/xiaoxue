<script setup lang="ts">
import { onMounted, ref } from "vue";
import { BusinessTypeList } from "@/api/constdata";
import { AgentIdName } from "@/api/types";
import { ProductIdName } from "@/api/types";
import { getAgentSimpleList } from "@/api/agent";
import { getProductInformationIdAndName } from "@/api/product";
const form = {
  business_type: 0,
  agent_id: "",
  product_id: "",
  add_method: 0,
  recharge_number: "",
  base_price: 0,
  password2: ""
};
const agentItemLists = ref([] as AgentIdName[]);
const productBaseInfoList = ref([] as ProductIdName[]);

onMounted(async () => {
  const response = await getAgentSimpleList();
  agentItemLists.value = response.data;
  var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
  const response3 = await getProductInformationIdAndName();
  productBaseInfoList.value = response3.data;
});

function handleSubmit() {
  console.log("Form submitted", form);
}
</script>

<template>
  <el-tabs>
    <el-tab-pane label="订单信息">
      <el-form label-width="82px">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="业务类型" prop="business_type">
              <el-select
                v-model="form.business_type"
                placeholder="请选择业务类型"
                clearable
                class="!w-[100%]"
              >
                <el-option
                  v-for="item in BusinessTypeList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="代理商" prop="agent_id">
              <el-select
                v-model="form.agent_id"
                placeholder="请选择代理商"
                clearable
                class="!w-[100%]"
              >
                <el-option
                  v-for="item in agentItemLists"
                  :key="item.id"
                  :label="item.id + ' ' + item.name"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="产品" prop="product_id">
              <el-select
                v-model="form.product_id"
                placeholder="请选择产品"
                clearable
                class="!w-[100%]"
              >
                <el-option
                  v-for="item in productBaseInfoList"
                  :key="item.id"
                  :label="item.id + ' ' + item.name"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="添加方式" prop="add_method">
              <el-select
                v-model="form.add_method"
                placeholder="请选择添加方式"
                clearable
                class="!w-[100%]"
              >
                <el-option :key="0" label="单号码添加" :value="0" />
                <el-option :key="1" label="批量添加" :value="1" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="充值号码" prop="recharge_number">
              <el-input
                v-model="form.recharge_number"
                placeholder="请输入充值号码"
                clearable
                class="!w-[100%]"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="面额" prop="base_price">
              <el-input
                v-model="form.base_price"
                type="number"
                placeholder="请输入面额"
                clearable
                class="!w-[100%]"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="二级密码" prop="password2">
              <el-input
                v-model="form.password2"
                type="password"
                placeholder="请输入二级密码"
                clearable
                class="!w-[100%]"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item>
              <el-button type="primary" @click="handleSubmit">提交</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-tab-pane>
  </el-tabs>
</template>
