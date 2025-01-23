<script setup lang="ts">
import { useCategory } from "./utils/hook";
import { ref, computed, nextTick, onMounted } from "vue";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { BusinessTypeList, OperatorListAll } from "@/api/constdata";
import { delay, deviceDetection, useResizeObserver } from "@pureadmin/utils";
import Refresh from "@iconify-icons/ep/refresh";
import Delete from "@iconify-icons/ep/delete";
import EditPen from "@iconify-icons/ep/edit-pen";
import AddFill from "@iconify-icons/ri/add-circle-line";
defineOptions({
  name: "AgentProductChannel"
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

const {
  form,
  curRow,
  disabled,
  selectedNum,
  loading,
  columns,
  rowStyle,
  dataList,
  pagination,
  agentItemLists,
  supplierItemLists,
  productCategoryList,
  productBaseInfoList,
  onSearch,
  resetForm,
  openDialog,
  handleSizeChange,
  handleDelete,
  handleCurrentChange,
  handleSelectionChange,
  onSelectionCancel
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
      <el-form-item label="代理商产品ID" prop="agent_product_id">
        <el-input
          v-model="form.agent_product_id"
          type="number"
          placeholder="请输入产品ID"
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
          :disabled="disabled"
        >
          <el-option
            v-for="item in agentItemLists"
            :key="item.id"
            :label="item.id + ' ' + item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="供货商产品ID" prop="supplier_product_id">
        <el-input
          v-model="form.supplier_product_id"
          type="number"
          placeholder="请输入产品ID"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="供货商" prop="supplier_id">
        <el-select
          v-model="form.supplier_id"
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
      <el-form-item label="上游产品ID" prop="up_product_id">
        <el-input
          v-model="form.up_product_id"
          placeholder="请输入产品ID"
          clearable
          class="!w-[180px]"
        />
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
            @click="openDialog()"
          >
            新增
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
              <el-popconfirm
                :title="`是否确认删除${row.ip}`"
                @confirm="handleDelete(row)"
              >
                <template #reference>
                  <el-button
                    class="reset-margin"
                    link
                    type="primary"
                    :size="size"
                    :icon="useRenderIcon(Delete)"
                  >
                    删除
                  </el-button>
                </template>
              </el-popconfirm>
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
