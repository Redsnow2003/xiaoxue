<script setup lang="ts">
import { useCategory } from "./utils/hook";
import { ref, computed, nextTick, onMounted } from "vue";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { delay, deviceDetection, useResizeObserver } from "@pureadmin/utils";
import { useProductHandlers } from "./utils/hook";
import Delete from "@iconify-icons/ep/delete";
import EditPen from "@iconify-icons/ep/edit-pen";
import Refresh from "@iconify-icons/ep/refresh";
import AddFill from "@iconify-icons/ri/add-circle-line";
import More2Fill from "@iconify-icons/ri/more-2-fill";
defineOptions({
  name: "AgentAccount"
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
const { handleWhitelist, handleProductConfig } = useProductHandlers();
const {
  form,
  loading,
  columns,
  rowStyle,
  dataList,
  pagination,
  selectedNum,
  handleBatchChange,
  handleBatchUpdate,
  handleUpdataInfo,
  handleChangeFund,
  handleChangeFundLog,
  handleDirectOrder,
  handleCheckAccount,
  handleSales,
  handleCreateAccount,
  onSearch,
  resetForm,
  handleAdd,
  handleDelete,
  handleSizeChange,
  onSelectionCancel,
  handleCurrentChange,
  handleSelectionChange,
  agentItemLists
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
      <el-form-item label="代理商" prop="id">
        <el-select
          v-model="form.id"
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
      <el-form-item label="通知方式" prop="notification_method">
        <el-select
          v-model="form.notification_method"
          placeholder="请选择通知方式"
          clearable
          class="!w-[180px]"
        >
          <el-option :key="1" label="可靠通知" :value="0" />
          <el-option :key="1" label="广播通知" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select
          v-model="form.status"
          placeholder="请选择状态"
          clearable
          class="!w-[180px]"
        >
          <el-option :key="1" label="维护" :value="0" />
          <el-option :key="1" label="上架" :value="1" />
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
            @click="handleAdd()"
          >
            新增代理商
          </el-button>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            :disabled="selectedNum === 0"
            @click="handleBatchChange()"
          >
            批量修改状态
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
                      @click="handleUpdataInfo(row)"
                    >
                      修改
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleDelete(row)"
                    >
                      删除
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleChangeFund(row)"
                    >
                      资金操作
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleChangeFundLog(row)"
                    >
                      资金操作日志
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleWhitelist(row)"
                    >
                      白名单
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleProductConfig(row)"
                    >
                      产品配置
                    </el-dropdown-item>
                    <el-dropdown-item :icon="useRenderIcon(EditPen)">
                      通道配置
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleSales(row)"
                    >
                      供货通道
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleCheckAccount(row)"
                    >
                      资金流水
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleDirectOrder(row)"
                    >
                      直充订单
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleDirectOrder(row)"
                    >
                      更新缓存金额
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleDirectOrder(row)"
                    >
                      更换密钥
                    </el-dropdown-item>
                    <el-dropdown-item
                      :icon="useRenderIcon(EditPen)"
                      @click="handleDirectOrder(row)"
                    >
                      开户信息
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
