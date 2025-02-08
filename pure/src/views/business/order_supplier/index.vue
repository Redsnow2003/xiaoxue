<script setup lang="ts">
import { useCategory } from "./utils/hook";
import { ref, computed, nextTick, onMounted } from "vue";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { getPickerShortcuts } from "@/api/utils";
import { delay, deviceDetection, useResizeObserver } from "@pureadmin/utils";
import { useProductHandlers } from "./utils/hook";
import {
  BusinessTypeList,
  SupplierOrderStatusList,
  ProvinceList
} from "@/api/constdata";
import Delete from "@iconify-icons/ep/delete";
import EditPen from "@iconify-icons/ep/edit-pen";
import Refresh from "@iconify-icons/ep/refresh";
import AddFill from "@iconify-icons/ri/add-circle-line";
import More2Fill from "@iconify-icons/ri/more-2-fill";
defineOptions({
  name: "SupplierOrderList"
});

const iconClass = computed(() => {
  return [
    "w-[22px]",
    "h-[22px]",
    "flex",
    "justify-center",
    "items-center",
    "outline-none",
    "rounded-[4px]",
    "cursor-pointer",
    "transition-colors",
    "hover:bg-[#0000000f]",
    "dark:hover:bg-[#ffffff1f]",
    "dark:hover:text-[#ffffffd9]"
  ];
});

const formRef = ref();
const tableRef = ref();
const contentRef = ref();
const { handleAgentproduct } = useProductHandlers();
const {
  form,
  loading,
  columns,
  rowStyle,
  dataList,
  pagination,
  selectedNum,
  handleExportCSV,
  handleBatchAutoCheckOrder,
  handleBatchManualCheckOrder,
  handleBatchCheckTemplate,
  handleBatchUrgeTemplate,
  handleBatchBackupSubmit,
  handleBatchCancelOrder,
  handleLookOrder,
  handleChangeRemark,
  handleQueryUpOrder,
  handleFailureToSuccess,
  handleBackupSubmit,
  handleSupplierProduct,
  handleOrderSubmitLog,
  handleOrderQueryLog,
  handleOrderCancelLog,
  handleOrderCallbackLog,
  onSearch,
  resetForm,
  handleSizeChange,
  onSelectionCancel,
  handleCurrentChange,
  handleSelectionChange,
  agentItemLists,
  supplierItemLists,
  productBaseInfoList,
  productCategoryList
} = useCategory(tableRef);

onMounted(async () => {
  useResizeObserver(contentRef, async () => {
    await nextTick();
    delay(60).then(() => {
      tableRef.value.setAdaptive();
    });
  });
});
</script>

<template>
  <div class="main">
    <el-form
      ref="formRef"
      :inline="true"
      :model="form"
      class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] overflow-auto"
    >
      <el-form-item label="业务类型" prop="business_type">
        <el-select
          v-model="form.business_type"
          placeholder="请选择业务类型"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in BusinessTypeList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="订单号" prop="order_id">
        <el-input
          v-model="form.order_id"
          placeholder="订单ID"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="供货单号" prop="id">
        <el-input
          v-model="form.id"
          placeholder="多个用,或者空格隔开"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="上游单号" prop="up_id">
        <el-input
          v-model="form.up_id"
          placeholder="上游单号"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="充值号码" prop="recharge_number">
        <el-input
          v-model="form.recharge_number"
          placeholder="多个用,或者空格隔开"
          clearable
          class="!w-[180px]"
        />
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
      <el-form-item label="产品类别" prop="product_category">
        <el-select
          v-model="form.product_category"
          placeholder="请选择产品类别"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in productCategoryList"
            :key="item.id"
            :label="item.category_name"
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
      <el-form-item label="基础价" prop="base_price">
        <el-input
          v-model="form.base_price"
          placeholder="请输入基础价"
          clearable
          number
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="备注" prop="base_price">
        <el-input
          v-model="form.base_price"
          placeholder="请输入备注"
          clearable
          number
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="供货单状态" prop="status">
        <el-select
          v-model="form.status"
          placeholder="请选择供货单状态"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in SupplierOrderStatusList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="供货商" prop="supplier">
        <el-select
          v-model="form.supplier"
          placeholder="请选择供货商"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in supplierItemLists"
            :key="item.id"
            :label="item.id + ' ' + item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="创建时间" prop="create_time">
        <el-date-picker
          v-model="form.create_time"
          :shortcuts="getPickerShortcuts()"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期时间"
          end-placeholder="结束日期时间"
        />
      </el-form-item>
      <el-form-item label="完成时间" prop="finish_time">
        <el-date-picker
          v-model="form.finish_time"
          :shortcuts="getPickerShortcuts()"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期时间"
          end-placeholder="结束日期时间"
        />
      </el-form-item>
      <el-form-item label="订单时间" prop="order_time">
        <el-date-picker
          v-model="form.order_time"
          :shortcuts="getPickerShortcuts()"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期时间"
          end-placeholder="结束日期时间"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :icon="useRenderIcon('ri:search-line')"
          :loading="loading"
          @click="onSearch"
        >
          搜索
        </el-button>
        <el-button :icon="useRenderIcon(Refresh)" @click="resetForm(formRef)">
          重置
        </el-button>
      </el-form-item>
    </el-form>

    <div
      ref="contentRef"
      :class="['flex', deviceDetection() ? 'flex-wrap' : '']"
    >
      <PureTableBar
        :class="[!deviceDetection() ? '!w-full' : 'w-full']"
        style="transition: width 220ms cubic-bezier(0.4, 0, 0.2, 1)"
        title=""
        :columns="columns"
        @refresh="onSearch"
      >
        <template #buttons>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            @click="handleExportCSV()"
          >
            导出
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchAutoCheckOrder()"
          >
            批量自动核单
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchManualCheckOrder()"
          >
            批量手动核单
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchCheckTemplate()"
          >
            批量核实客服模板
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchUrgeTemplate()"
          >
            批量催单客服模板
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchBackupSubmit()"
          >
            批量备用通道重新提交
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchCancelOrder()"
          >
            批量撤销申请
          </el-button>
        </template>
        <template v-slot="{ size, dynamicColumns }">
          <div
            v-if="selectedNum > 0"
            v-motion-fade
            class="bg-[var(--el-fill-color-light)] w-full h-[46px] mb-2 pl-4 flex items-center"
          >
            <div class="flex-auto">
              <span
                style="font-size: var(--el-font-size-base)"
                class="text-[rgba(42,46,54,0.5)] dark:text-[rgba(220,220,242,0.5)]"
              >
                已选 {{ selectedNum }} 项
              </span>
              <el-button type="primary" text @click="onSelectionCancel">
                取消选择
              </el-button>
            </div>
          </div>
          <pure-table
            ref="tableRef"
            row-key="id"
            align-whole="center"
            showOverflowTooltip
            table-layout="auto"
            :loading="loading"
            :size="size"
            adaptive
            :row-style="rowStyle"
            :adaptiveConfig="{ offsetBottom: 108 }"
            :data="dataList"
            :columns="dynamicColumns"
            :pagination="{ ...pagination, size }"
            :header-cell-style="{
              background: 'var(--el-fill-color-light)',
              color: 'var(--el-text-color-primary)'
            }"
            @selection-change="handleSelectionChange"
            @page-size-change="handleSizeChange"
            @page-current-change="handleCurrentChange"
          >
            <template #operation="{ row }">
              <el-dropdown trigger="click">
                <IconifyIconOffline :icon="More2Fill" class="text-[24px]" />
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleLookOrder(row)"
                    >
                      查看订单
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleChangeRemark(row)"
                    >
                      修改备注
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleQueryUpOrder(row)"
                    >
                      查询上游订单
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleFailureToSuccess(row)"
                    >
                      失败转成功
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleBackupSubmit(row)"
                    >
                      备用通道重新提交
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleSupplierProduct(row)"
                    >
                      供货商产品配置
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleOrderSubmitLog(row)"
                    >
                      提单日志
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleOrderQueryLog(row)"
                    >
                      查单日志
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleOrderCancelLog(row)"
                    >
                      撤单日志
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleOrderCallbackLog(row)"
                    >
                      回调日志
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </template>
          </pure-table>
        </template>
      </PureTableBar>
    </div>
  </div>
</template>

<style scoped lang="scss">
:deep(.el-dropdown-menu__item i) {
  margin: 0;
}

.main-content {
  margin: 24px 24px 0 !important;
}

.search-form {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
