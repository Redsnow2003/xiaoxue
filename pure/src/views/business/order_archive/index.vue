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
  NotifyStatusList,
  OrderStatusList,
  ProvinceList
} from "@/api/constdata";
import Delete from "@iconify-icons/ep/delete";
import EditPen from "@iconify-icons/ep/edit-pen";
import Refresh from "@iconify-icons/ep/refresh";
import AddFill from "@iconify-icons/ri/add-circle-line";
import More2Fill from "@iconify-icons/ri/more-2-fill";
defineOptions({
  name: "OrderArchiveList"
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
  handleBatchStatus,
  handleBatchNotify,
  handleBatchBackup,
  handleBatchCancelBackup,
  handleBatchTimeout,
  handleBatchCancel,
  handleBatchManual,
  handleLookSupplierOrder,
  handleNotifyStatus,
  handleChangeRemark,
  handleBackupLog,
  handleSubmitLog,
  handleQueryLog,
  handleNotifyLog,
  onSearch,
  resetForm,
  handleSizeChange,
  onSelectionCancel,
  handleCurrentChange,
  handleSelectionChange,
  agentItemLists,
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
      <el-form-item label="订单ID" prop="id">
        <el-input
          v-model="form.id"
          placeholder="多个用,或者空格隔开"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="下游订单号" prop="down_id">
        <el-input
          v-model="form.down_id"
          placeholder="多个用,或者空格隔开"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="通知状态" prop="notify_status">
        <el-select
          v-model="form.notify_status"
          placeholder="请选择通知状态"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in NotifyStatusList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
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
      <el-form-item label="订单状态" prop="status">
        <el-select
          v-model="form.status"
          placeholder="请选择订单状态"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in OrderStatusList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="是否超时" prop="is_timeout">
        <el-select
          v-model="form.is_timeout"
          placeholder="请选择是否超时"
          clearable
          class="!w-[180px]"
        >
          <el-option :key="0" label="否" :value="0" />
          <el-option :key="1" label="是" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="是否退单" prop="is_cancel">
        <el-select
          v-model="form.is_cancel"
          placeholder="请选择是否退单"
          clearable
          class="!w-[180px]"
        >
          <el-option :key="0" label="否" :value="0" />
          <el-option :key="1" label="是" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="号码归属地" prop="location">
        <el-select
          v-model="form.location"
          placeholder="请选择号码归属地"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in ProvinceList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="特殊参数" prop="special_params">
        <el-input
          v-model="form.special_params"
          placeholder="请输入特殊参数"
          clearable
          number
          class="!w-[180px]"
        />
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
            @click="handleBatchStatus()"
          >
            批量处理缓存单
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchNotify()"
          >
            批量通知
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchBackup()"
          >
            批量备用通道重新提交
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchCancelBackup()"
          >
            批量备用通道取消
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchTimeout()"
          >
            批量超时
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchCancel()"
          >
            批量退单
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchManual()"
          >
            批量转人工处理
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
                      @click="handleLookSupplierOrder(row)"
                    >
                      查看供货单
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleNotifyStatus(row)"
                    >
                      通知状态
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleChangeRemark(row)"
                    >
                      修改备注
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleBackupLog(row)"
                    >
                      备用通道记录
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleAgentproduct(row)"
                    >
                      代理商产品配置
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleSubmitLog(row)"
                    >
                      提单日志
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleQueryLog(row)"
                    >
                      查单日志
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleNotifyLog(row)"
                    >
                      通知日志
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
