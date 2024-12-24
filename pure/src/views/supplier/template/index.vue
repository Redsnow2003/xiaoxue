<script setup lang="ts">
import { useCategory } from "./utils/hook";
import { ref, computed, nextTick, onMounted } from "vue";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { delay, deviceDetection, useResizeObserver } from "@pureadmin/utils";
import Refresh from "@iconify-icons/ep/refresh";

defineOptions({
  name: "SupplierTemplate"
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
  loading,
  columns,
  rowStyle,
  dataList,
  pagination,
  onSearch,
  resetForm,
  handleSizeChange,
  handleCurrentChange
} = useCategory();

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
      <el-form-item label="模板名称：" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入模板名称"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="提单路径：" prop="submit_address">
        <el-input
          v-model="form.submit_address"
          placeholder="请输入提单路径"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="查单路径：" prop="query_address">
        <el-input
          v-model="form.query_address"
          placeholder="请输入查单路径"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="余额路径：" prop="balance_address">
        <el-input
          v-model="form.balance_address"
          placeholder="请输入余额路径"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="备注：" prop="remark">
        <el-input
          v-model="form.remark"
          placeholder="请输入备注"
          clearable
          class="!w-[180px]"
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
        <template v-slot="{ size, dynamicColumns }">
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
            @page-size-change="handleSizeChange"
            @page-current-change="handleCurrentChange"
          />
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
