<script setup lang="ts">
import { useCategory } from "./utils/hook";
import { ref, computed, nextTick, onMounted } from "vue";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { delay, deviceDetection, useResizeObserver } from "@pureadmin/utils";
import { OperatorList1, ProvinceList } from "@/api/constdata";
import Delete from "@iconify-icons/ep/delete";
import EditPen from "@iconify-icons/ep/edit-pen";
import Refresh from "@iconify-icons/ep/refresh";
import AddFill from "@iconify-icons/ri/add-circle-line";

defineOptions({
  name: "ProductInformation-PowerBill"
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
  selectedNum,
  handleImport,
  handleExport,
  onSearch,
  onbatchDel,
  resetForm,
  openDialog,
  handleDelete,
  handleSizeChange,
  onSelectionCancel,
  handleCurrentChange,
  handleSelectionChange,
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
      <el-form-item label="产品ID：" prop="id">
        <el-input
          v-model="form.id"
          placeholder="请输入产品ID"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="产品名称：" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入产品名称"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="产品类别" prop="category">
        <el-select
          v-model="form.category"
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
      <el-form-item label="运营商" prop="operator">
        <el-select
          v-model="form.operator"
          placeholder="请选择运营商"
          clearable
          class="!w-[180px]"
        >
          <el-option
            v-for="item in OperatorList1"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="面额" prop="price">
        <el-input
          v-model="form.price"
          placeholder="请输入面额"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="使用范围" prop="scope">
        <el-select
          v-model="form.scope"
          placeholder="请选使用范围"
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
            新增产品
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            @click="handleImport()"
          >
            从系统模板新增
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            @click="handleExport()"
          >
            导出
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
            <el-popconfirm title="是否确认删除?" @confirm="onbatchDel">
              <template #reference>
                <el-button type="danger" text class="mr-1">
                  批量删除
                </el-button>
              </template>
            </el-popconfirm>
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
              <el-button
                class="reset-margin"
                link
                type="primary"
                :size="size"
                :icon="useRenderIcon(EditPen)"
                @click="openDialog('修改', row)"
              >
                修改
              </el-button>
              <el-popconfirm
                :title="`是否确认删除话费类别${row.name}`"
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
              <el-button
                class="reset-margin"
                link
                type="primary"
                :size="size"
                :icon="useRenderIcon(EditPen)"
                @click="openDialog('复制', row)"
              >
                复制
              </el-button>
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
