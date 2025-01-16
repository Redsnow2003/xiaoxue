<script setup lang="ts">
import { ref } from "vue";
import { useDept } from "./utils/hook";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { BusinessTypeList, OperatorListAll } from "@/api/constdata";
import { delay, deviceDetection, useResizeObserver } from "@pureadmin/utils";
import { useProductHandlers } from "./utils/hook";
import Delete from "@iconify-icons/ep/delete";
import EditPen from "@iconify-icons/ep/edit-pen";
import Refresh from "@iconify-icons/ep/refresh";
import AddFill from "@iconify-icons/ri/add-circle-line";

defineOptions({
  name: "AgentProduct"
});

const formRef = ref();
const tableRef = ref();
const { handleProductChannel } = useProductHandlers();
const {
  form,
  loading,
  columns,
  dataList,
  selectedNum,
  pagination,
  agentItemLists,
  productCategoryList,
  productBaseInfoList,
  onSearch,
  resetForm,
  newProcuct,
  changeProduct,
  batchChangeProduct,
  batchChangeDiscount,
  handleDelete,
  handleBatchDelete,
  handleSelectionChange,
  onSelectionCancel
} = useDept(tableRef);

function onFullscreen() {
  // 重置表格高度
  tableRef.value.setAdaptive();
}
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
      <el-form-item label="产品名称" prop="product_name">
        <el-input
          v-model="form.product_name"
          placeholder="请输入产品名称"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="产品运营商" prop="operator">
        <el-select
          v-model="form.operator"
          placeholder="请选择产品运营商"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in OperatorListAll"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="状态：" prop="status">
        <el-select
          v-model="form.status"
          placeholder="请选择状态"
          clearable
          class="!w-[180px]"
        >
          <el-option label="维护" :value="0" />
          <el-option label="上架" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="支持缓存" prop="support_cache">
        <el-select
          v-model="form.support_cache"
          placeholder="请选择支持缓存"
          clearable
          class="!w-[180px]"
        >
          <el-option label="支持" :value="0" />
          <el-option label="不支持" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="转网检测" prop="transfer_check">
        <el-select
          v-model="form.transfer_check"
          placeholder="请选择转网检测"
          clearable
          class="!w-[180px]"
        >
          <el-option label="不启用" :value="0" />
          <el-option label="启用" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="空号检测" prop="empty_check">
        <el-select
          v-model="form.empty_check"
          placeholder="请选择空号检测"
          clearable
          class="!w-[180px]"
        >
          <el-option label="不启用" :value="0" />
          <el-option label="启用" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="自动提交备用通道" prop="auto_submit_backup">
        <el-select
          v-model="form.auto_submit_backup"
          placeholder="请选择空号检测"
          clearable
          class="!w-[180px]"
        >
          <el-option label="不启用" :value="0" />
          <el-option label="启用" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="超时不缓存" prop="timeout_not_cache">
        <el-select
          v-model="form.timeout_not_cache"
          placeholder="请选择超时不缓存"
          clearable
          class="!w-[180px]"
        >
          <el-option label="启动" :value="0" />
          <el-option label="不启动" :value="1" />
        </el-select>
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

    <PureTableBar
      title="供应商提供的产品列表"
      :columns="columns"
      :class="[!deviceDetection() ? '!w-full' : 'w-full']"
      @refresh="onSearch"
      @fullscreen="onFullscreen"
    >
      <template #buttons>
        <el-button
          type="primary"
          :icon="useRenderIcon(AddFill)"
          @click="newProcuct()"
        >
          新增
        </el-button>
        <el-button
          type="primary"
          :icon="useRenderIcon(AddFill)"
          :disabled="selectedNum <= 0"
          @click="batchChangeProduct()"
        >
          批量修改
        </el-button>
        <el-button
          type="primary"
          :icon="useRenderIcon(AddFill)"
          @click="batchChangeDiscount()"
        >
          批量折扣
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
          <el-popconfirm title="是否确认删除?" @confirm="handleBatchDelete">
            <template #reference>
              <el-button type="danger" text class="mr-1"> 批量删除 </el-button>
            </template>
          </el-popconfirm>
        </div>
        <pure-table
          ref="tableRef"
          adaptive
          :adaptiveConfig="{ offsetBottom: 45 }"
          align-whole="center"
          row-key="id"
          showOverflowTooltip
          table-layout="auto"
          default-expand-all
          :loading="loading"
          :size="size"
          :data="dataList"
          :columns="dynamicColumns"
          :pagination="{ ...pagination, size }"
          :header-cell-style="{
            background: 'var(--el-fill-color-light)',
            color: 'var(--el-text-color-primary)'
          }"
          @selection-change="handleSelectionChange"
        >
          <template #operation="{ row }">
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(EditPen)"
              @click="changeProduct(row)"
            >
              修改
            </el-button>
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(Delete)"
              @click="handleDelete(row)"
            >
              删除
            </el-button>
            <el-button
              class="reset-margin"
              link
              type="primary"
              :size="size"
              :icon="useRenderIcon(Delete)"
              @click="handleProductChannel(row)"
            >
              供货通道
            </el-button>
          </template>
        </pure-table>
      </template>
    </PureTableBar>
  </div>
</template>

<style lang="scss" scoped>
:deep(.el-table__inner-wrapper::before) {
  height: 0;
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
