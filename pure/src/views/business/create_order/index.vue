<script setup lang="ts">
import { ref } from "vue";
import { BusinessTypeList } from "@/api/constdata";
import { AgentIdName } from "@/api/types";
import { ProductIdName } from "@/api/types";
const form = {
  business_type: 0,
  agent_id: 0,
  product_id: 0,
  add_method: 0,
  recharge_number: "",
  base_price: 0,
  password2: ""
};
const agentItemLists = ref([] as AgentIdName[]);
const productBaseInfoList = ref([] as ProductIdName[]);
</script>

<template>
  <el-form label-width="82px">
    <el-form-item label="业务类型" prop="business_type">
      <el-select
        v-model="form.business_type"
        placeholder="请选择业务类型"
        clearable
        class="!w-[380px]"
      >
        <el-option
          v-for="item in BusinessTypeList"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="代理商" prop="agent_id">
      <el-select
        v-model="form.agent_id"
        placeholder="请选择代理商"
        clearable
        class="!w-[180px]"
      >
        <el-option
          v-for="item in agentItemLists"
          :key="item.id"
          :label="item.id + ' ' + item.name"
          :value="item.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="产品" prop="product_id">
      <el-select
        v-model="form.product_id"
        placeholder="请选择产品"
        clearable
        class="!w-[180px]"
      >
        <el-option
          v-for="item in productBaseInfoList"
          :key="item.id"
          :label="item.id + ' ' + item.name"
          :value="item.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="添加方式" prop="add_method">
      <el-select
        v-model="form.add_method"
        placeholder="请选择添加方式"
        clearable
        class="!w-[180px]"
      >
        <el-option :key="0" label="单号码添加" :value="0" />
        <el-option :key="1" label="批量添加" :value="1" />
      </el-select>
    </el-form-item>
    <el-form-item label="充值号码" prop="recharge_number">
      <el-input
        v-model="form.recharge_number"
        placeholder="请输入充值号码"
        clearable
        class="!w-[180px]"
      />
    </el-form-item>
    <el-form-item label="面额" prop="base_price">
      <el-input
        v-model="form.base_price"
        type="number"
        placeholder="请输入面额"
        clearable
        class="!w-[180px]"
      />
    </el-form-item>
    <el-form-item label="二级密码" prop="password2">
      <el-input
        v-model="form.base_price"
        type="password"
        placeholder="请输入二级密码"
        clearable
        class="!w-[180px]"
      />
    </el-form-item>
  </el-form>
</template>
