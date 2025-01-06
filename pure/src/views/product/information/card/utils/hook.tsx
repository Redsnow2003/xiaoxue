import editForm from "../form.vue";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type { FormItemProps, CategoryProps } from "../utils/types";
import type { PaginationProps } from "@pureadmin/table";
import { getKeyList, deviceDetection } from "@pureadmin/utils";
import {
  getProductCategoryList,
  getProductInformationList,
  addProductInformation,
  updateProductInformation,
  deleteProductInformation,
  importProductInformation,
  exportProductInformation
} from "@/api/product";
import { type Ref, reactive, ref, onMounted, h } from "vue";
import { OperatorListTelecom, UnitList } from "@/api/constdata";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    /** 类别ID */
    id: "",
    type: 3,
    /** 产品类别 */
    category: "",
    /** 产品名称 */
    name: "",
    /** 面额 */
    price: ""
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const selectedNum = ref(0);
  const loading = ref(true);
  const productCategoryList = ref([] as CategoryProps);
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
      label: "产品ID",
      prop: "id"
    },
    {
      label: "产品名称",
      prop: "name"
    },
    {
      label: "产品类别",
      prop: "category",
      cellRenderer: ({ row }) => (
        <span>
          {
            productCategoryList.value.find(item => item.id == row.category)
              ?.category_name
          }
        </span>
      )
    },
    {
      label: "产品运营商",
      prop: "operator",
      cellRenderer: ({ row }) => (
        <span>
          {OperatorListTelecom.find(item => item.value == row.operator)?.label}
        </span>
      )
    },
    {
      label: "面额",
      prop: "price"
    },
    {
      label: "单位",
      prop: "unit",
      cellRenderer: ({ row }) => (
        <span>{UnitList.find(unit => unit.value == row.unit)?.label}</span>
      )
    },
    {
      label: "基础价（元）",
      prop: "base_price"
    },
    {
      label: "待售库存",
      prop: "sale_inventory"
    },
    {
      label: "销售库存总金额",
      prop: "sale_inventory_amount"
    },
    {
      label: "API提单数量限制",
      prop: "api_limit"
    },
    {
      label: "备注",
      prop: "remark"
    },
    {
      label: "操作",
      fixed: "right",
      width: 210,
      slot: "operation"
    }
  ];

  function handleDelete(row) {
    var ids = [row.id];
    deleteProductInformation(ids).then(res => {
      if (res.success) {
        message("删除产品成功", { type: "success" });
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

  /** 导入 */
  async function handleImport() {
    const res = await importProductInformation();
    if (res.success) {
      message(res.message, { type: "success" });
    } else {
      message(res.message, { type: "error" });
    }
    onSearch();
  }

  /** 导出 */
  async function handleExport() {
    const res = await exportProductInformation();
    if (res.success) {
      message(res.message, { type: "success" });
    } else {
      message(res.message, { type: "error" });
    }
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
    await deleteProductInformation(ids).then(res => {
      if (res.success) {
        message("批量删除产品成功", { type: "success" });
      } else {
        message(res.message, { type: "error" });
      }
    });
    tableRef.value.getTableRef().clearSelection();
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    var params = { ...form, ...pagination };
    console.log("params", params);
    const { data } = await getProductInformationList(params);
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

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}产品信息`,
      props: {
        formInline: {
          id: row?.id ?? 0,
          type: 3,
          category: row?.category ?? "",
          name: row?.name ?? "",
          operator: row?.operator ?? "",
          unit: 6,
          price: row?.price ?? "",
          base_price: row?.base_price ?? "",
          sale_inventory: row?.sale_inventory ?? "",
          sale_inventory_amount: row?.sale_inventory_amount ?? "",
          api_limit: row?.api_limit ?? "",
          remark: row?.remark ?? ""
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(editForm, {
          ref: formRef,
          formInline: null
        }),
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
            var data = {
              ...curData,
              base_price: Number(curData.base_price),
              price: Number(curData.price),
              api_limit: Number(curData.api_limit)
            };
            // 表单规则校验通过
            if (title === "新增" || title === "复制") {
              data.id = 0;
              await addProductInformation(data).then(res => {
                if (res.success) {
                  message("新增成功", { type: "success" });
                } else {
                  message(res.message, { type: "error" });
                }
              });
              chores();
            } else {
              await updateProductInformation(data).then(res => {
                if (res.success) {
                  message("修改成功", { type: "success" });
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
    // 动态获取产品类别列表
    var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
    const response = await getProductCategoryList(requestData);
    productCategoryList.value = response.data.list;
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
    handleImport,
    handleExport,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange,
    productCategoryList
  };
}
