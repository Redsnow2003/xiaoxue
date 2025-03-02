import { getNumberBlackList, deleteNumberBlackList } from "@/api/order";
import type { PaginationProps } from "@pureadmin/table";
import { type Ref, reactive, ref, onMounted } from "vue";
import type { BacklistItemProps } from "./types";
import { getKeyList } from "@pureadmin/utils";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    recharge_number: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const selectedNum = ref(0);
  const loading = ref(true);

  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "勾选列", // 如果需要表格多选，此处label必须设置
      type: "selection",
      fixed: "left",
      reserveSelection: true // 数据刷新后保留选项
    },
    {
      label: "ID",
      prop: "id"
    },
    {
      label: "黑名单",
      prop: "recharge_number"
    },
    {
      label: "备注",
      prop: "remark"
    },
    {
      label: "操作",
      fixed: "right",
      width: 100,
      slot: "operation"
    }
  ];

  function handleSizeChange(val: number) {
    pagination.pageSize = val;
    onSearch();
  }

  function handleCurrentChange(val: number) {
    pagination.currentPage = val;
    onSearch();
  }

  /** 当CheckBox选择项发生变化时会触发该事件 */
  function handleSelectionChange(val: Array<any>) {
    selectedNum.value = val.length;
    // 重置表格高度
    tableRef.value.setAdaptive();
  }

  /** 取消选择 */
  function onSelectionCancel() {
    selectedNum.value = 0;
    // 用于多选表格，清空用户的选择
    tableRef.value.getTableRef().clearSelection();
  }

  async function onSearch() {
    loading.value = true;
    var params = { ...form, ...pagination };
    const { data } = await getNumberBlackList(params);
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    setTimeout(() => {
      loading.value = false;
    }, 100);
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  /** 高亮当前权限选中行 */
  function rowStyle({ row: { id } }) {
    return {
      cursor: "pointer",
      background: id === curRow.value?.id ? "var(--el-fill-color-light)" : ""
    };
  }
  onMounted(async () => {
    onSearch();
  });

  function handleAdd() {
    console.log("handleAdd");
  }

  function handleImport() {
    console.log("handleAdd");
  }

  function handleImportTemplateDownload() {
    console.log("handleAdd");
  }

  function handleBatchDelete() {
    // 获取选中的数据
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    deleteNumberBlackList(ids).then(() => {
      onSearch();
      onSelectionCancel();
    });
  }

  function handleDelete(row?: BacklistItemProps) {
    var params = [row.id];
    deleteNumberBlackList(params).then(() => {
      onSearch();
      onSelectionCancel();
    });
  }

  return {
    form,
    curRow,
    loading,
    columns,
    selectedNum,
    rowStyle,
    dataList,
    pagination,
    onSearch,
    handleAdd,
    handleImport,
    handleImportTemplateDownload,
    handleBatchDelete,
    resetForm,
    handleDelete,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange
  };
}
