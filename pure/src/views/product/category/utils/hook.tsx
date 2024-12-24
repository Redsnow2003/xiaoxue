import editForm from "../form.vue";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type { FormItemProps } from "../utils/types";
import type { PaginationProps } from "@pureadmin/table";
import { getKeyList, deviceDetection } from "@pureadmin/utils";
import {
  getProductCategoryList,
  addProductCategory,
  updateProductCategory,
  deleteProductCategory
} from "@/api/product";
import { type Ref, reactive, ref, onMounted, h } from "vue";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    category_name: ""
  });
  const curRow = ref();
  const formRef = ref();
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
      label: "类别ID",
      prop: "id"
    },
    {
      label: "类别名称",
      prop: "category_name"
    },
    {
      label: "操作",
      fixed: "right",
      width: 210,
      slot: "operation"
    }
  ];

  async function handleDelete(row) {
    var ids = [row.id];
    await deleteProductCategory(ids).then(res => {
      if (res.success) {
        message("删除分类成功", { type: "success" });
      } else {
        message(res.message, { type: "error" });
      }
    });
    onSearch();
  }

  function handleSizeChange(val: number) {
    pagination.pageSize = val;
    onSearch();
  }

  function handleCurrentChange(val: number) {
    pagination.currentPage = val;
    onSearch();
  }

  /** 当CheckBox选择项发生变化时会触发该事件 */
  function handleSelectionChange(val) {
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

  /** 批量删除 */
  async function onbatchDel() {
    // 返回当前选中的行
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await deleteProductCategory(ids).then(res => {
      if (res.success) {
        message("删除分类成功", { type: "success" });
      } else {
        message(res.message, { type: "error" });
      }
    });
    tableRef.value.getTableRef().clearSelection();
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    var requestData = { ...form, ...pagination };
    const { data } = await getProductCategoryList(requestData);
    console.log("data", data);
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}产品类别`,
      props: {
        formInline: {
          id: row?.id ?? 0,
          category_name: row?.category_name ?? ""
        }
      },
      width: "40%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            console.log("curData", curData);
            // 表单规则校验通过
            if (title === "新增") {
              await addProductCategory(curData).then(res => {
                if (res.success) {
                  message("新增分类成功", { type: "success" });
                } else {
                  message(res.message, { type: "error" });
                }
              });
              chores();
            } else {
              await updateProductCategory(curData).then(res => {
                if (res.success) {
                  message("更新分类成功", { type: "success" });
                } else {
                  message(res.message, { type: "error" });
                }
              });
              chores();
            }
          }
        });
      }
    });
  }

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

  return {
    form,
    curRow,
    loading,
    columns,
    rowStyle,
    dataList,
    onbatchDel,
    pagination,
    selectedNum,
    onSearch,
    resetForm,
    openDialog,
    handleDelete,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange
  };
}
