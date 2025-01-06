<script setup lang="ts">
import { ref } from "vue";
import { formRules } from "../utils/rule";
import { AgentProductProps } from "../utils/types";
import {
  SupplyStrategyList,
  ProvinceList,
  OperatorListAll
} from "@/api/constdata";

const props = withDefaults(defineProps<AgentProductProps>(), {
  formInline: () => ({
    /** 代理商产品ID */
    id: 0,
    /** 业务类型 */
    business_type: 0,
    /** 代理商ID */
    agent_id: 0,
    /** 代理商名称 */
    agent_name: "",
    /** 产品ID */
    product_id: 0,
    /** 产品名称 */
    product_name: "",
    /** 产品类别 */
    product_category: 0,
    /** 运营商 */
    operator: 0,
    /** 基础价格 */
    base_price: 0,
    /** 供货策略 */
    supply_strategy: 0,
    /** 备用通道供货策略 */
    backup_channel_strategy: 0,
    /** 折扣类型 */
    discount_type: 0,
    /** 折扣 */
    discount: 0,
    /** 超时时间 */
    timeout: 0,
    /** 超时不缓存 */
    timeout_not_cache: 0,
    /** 自动提交备份 */
    auto_submit_backup: 0,
    /** 内部时间 */
    interal_time: 0,
    /** 支持缓存 */
    support_cache: 0,
    /** 转网检测 */
    transfer_check: 0,
    /** 空号检测 */
    empty_check: 0,
    /** 禁用地区 */
    disabled_area: [],
    /** 可用地区 */
    scope: 0,
    enabled_area: [],
    /** 限定运营商 */
    limit_operator: [],
    /** 状态 */
    status: 0,
    /** 备注 */
    remark: "",
    /** 定时更改时间 */
    timing_change_time: ""
  })
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);
console.log("newFormInline", newFormInline);
if (newFormInline.value.product_name === undefined) {
  newFormInline.value.product_name = "全部所选产品";
}
function getRef() {
  return ruleFormRef.value;
}
defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="130px"
  >
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form-item label="产品名称" prop="product_name">
          <el-input v-model="newFormInline.product_name" disabled />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="折扣方式" prop="discount_type">
          <el-select v-model="newFormInline.discount_type">
            <el-option :key="0" label="按折扣" :value="0" />
            <el-option :key="1" label="按金额" :value="1" />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item
          v-if="newFormInline.discount_type === 0"
          label="产品折扣"
          prop="discount"
        >
          <el-input
            v-model="newFormInline.discount"
            placeholder="请输入产品折扣"
          />
        </el-form-item>
        <el-form-item
          v-if="newFormInline.discount_type === 1"
          label="代理金额"
          prop="discount"
        >
          <el-input
            v-model="newFormInline.discount"
            placeholder="请输入产品金额"
          />
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form-item label="状态" prop="status">
          <el-select v-model="newFormInline.status">
            <el-option :key="0" label="维护" :value="0" />
            <el-option :key="1" label="上架" :value="1" />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="供货策略" prop="supply_strategy">
          <el-select v-model="newFormInline.supply_strategy">
            <el-option
              v-for="item in SupplyStrategyList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="是否支持缓存" prop="support_cache">
          <el-select v-model="newFormInline.support_cache">
            <el-option :key="0" label="支持" :value="0" />
            <el-option :key="1" label="不支持" :value="1" />
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form-item label="转网检测" prop="transfer_check">
          <el-select v-model="newFormInline.transfer_check">
            <el-option :key="0" label="不启用" :value="0" />
            <el-option :key="1" label="启用" :value="1" />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="空号检测" prop="empty_check">
          <el-select v-model="newFormInline.empty_check">
            <el-option :key="0" label="不启用" :value="0" />
            <el-option :key="1" label="启用" :value="1" />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="超时不缓存" prop="timeout_not_cache">
          <el-select v-model="newFormInline.timeout_not_cache">
            <el-option :key="0" label="启动" :value="0" />
            <el-option :key="1" label="不启动" :value="1" />
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form-item label="超时时间(秒)" prop="timeout">
          <el-input
            v-model="newFormInline.timeout"
            placeholder="请输入超时时间"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="备用通道供货策略" prop="backup_channel_strategy">
          <el-select v-model="newFormInline.backup_channel_strategy">
            <el-option
              v-for="item in SupplyStrategyList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="自动提交备用通道" prop="auto_submit_backup">
          <el-select v-model="newFormInline.auto_submit_backup">
            <el-option :key="0" label="不启用" :value="0" />
            <el-option :key="1" label="启用" :value="1" />
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form-item label="禁用地区" prop="disable_area">
          <el-select v-model="newFormInline.disabled_area" multiple>
            <el-option
              v-for="item in ProvinceList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="可用地区" prop="enabled_area">
          <el-select v-model="newFormInline.enabled_area" multiple>
            <el-option
              v-for="item in ProvinceList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="限定运营商" prop="limit_operator">
          <el-select v-model="newFormInline.limit_operator" multiple>
            <el-option
              v-for="item in OperatorListAll"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form-item label="备注" prop="remark">
          <el-input
            v-model="newFormInline.remark"
            placeholder="请输入产品备注"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="定时更改时间" prop="timing_change_time">
          <el-date-picker
            v-model="newFormInline.timing_change_time"
            type="datetime"
            placeholder="选择日期时间"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>
