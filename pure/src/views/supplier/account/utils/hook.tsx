import addForm from "../form/addsupplier.vue";
import batchChangeForm from "../form/batchchange.vue";
import balanceLogForm from "../form/balancelog.vue";
import changeInfoForm from "../form/changeinfo.vue";
import changeFundForm from "../form/changefund.vue";
import { createFormData } from "@pureadmin/utils";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type {
  FormItemProps,
  SupplierSimpleItem,
  TemplateNameItem
} from "../utils/types";
import type { PaginationProps } from "@pureadmin/table";
import { deviceDetection, getKeyList } from "@pureadmin/utils";
import {
  getSupplierSimpleList,
  getSupplierTemplateNameList,
  addSupplier,
  updateSupplier,
  deleteSupplier,
  getSupplierList,
  batchUpdateSupplierStatus,
  batchUpdateSupplierBalance,
  supplierChangeFund
} from "@/api/supplier";
import { type Ref, reactive, ref, onMounted, h } from "vue";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    /** 供货商 */
    name: "",
    /** 上游模板 */
    up_template: "",
    /** 状态 */
    status: ""
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const selectedNum = ref(0);
  const loading = ref(true);
  const supplierItemLists = ref([] as SupplierSimpleItem[]);
  const templateNameLists = ref([] as TemplateNameItem[]);
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
      label: "供应商ID",
      prop: "id"
    },
    {
      label: "供应商名称",
      prop: "name"
    },
    {
      label: "我方平台账户余额(元)",
      prop: "our_balance"
    },
    {
      label: "上游平台账户余额(元)",
      prop: "up_balance"
    },
    {
      label: "上游平台账户余额更新时间",
      prop: "up_balance_update_time"
    },
    {
      label: "上游模板配置",
      prop: "up_template"
    },
    {
      label: "备注",
      prop: "remark"
    },
    {
      label: "状态",
      prop: "status"
    },
    {
      label: "状态信息",
      prop: "status_info"
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
    deleteSupplier(ids).then(res => {
      if (res.success) {
        message("删除供应商成功", { type: "success" });
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

  async function onSearch() {
    loading.value = true;
    var params = { ...form, ...pagination };
    console.log("params", params);
    const { data } = await getSupplierList(params);
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

  function handleAdd() {
    addDialog({
      title: `新增供应商`,
      props: {
        formInline: null
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(addForm, {
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
              ...curData
            };
            // 表单规则校验通过
            await addSupplier(data).then(res => {
              if (res.success) {
                message("新增成功", { type: "success" });
              } else {
                message(res.message, { type: "error" });
              }
            });
            chores();
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
    const response = await getSupplierSimpleList();
    supplierItemLists.value = response.data;
    const response2 = await getSupplierTemplateNameList();
    templateNameLists.value = response2.data;
  });

  function handleBatchChange() {
    addDialog({
      title: `批量修改供应商状态`,
      props: {
        status: 0
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(batchChangeForm, {
          ref: formRef,
          status: 0
        }),
      beforeSure: (done, { options }) => {
        const curSelected = tableRef.value.getTableRef().getSelectionRows();
        var ids = getKeyList(curSelected, "id");
        const FormRef = formRef.value.getRef();
        const curData = { status: options.props.status as number, ids: ids };
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            // 表单规则校验通过
            await batchUpdateSupplierStatus(curData).then(res => {
              if (res.success) {
                message("修改成功", { type: "success" });
              } else {
                message(res.message, { type: "error" });
              }
            });
            chores();
          }
        });
      }
    });
  }

  async function handleBatchUpdate() {
    await batchUpdateSupplierBalance().then(res => {
      if (res.success) {
        message("批量更新供应商余额成功", { type: "success" });
      } else {
        message(res.message, { type: "error" });
      }
    });
  }

  function handleChangeFund(row?: FormItemProps) {
    addDialog({
      title: `操作资金`,
      props: {
        formInline: {
          supplierId: row?.id ?? "",
          supplierName: row?.name ?? "",
          fundAction: "",
          amount: 0,
          confirmAmount: 0,
          fileList: [],
          remark: ""
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(changeFundForm, {
          ref: formRef,
          formInline: null
        }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = createFormData({
          ...options.props.formInline,
          fileList: options.props.formInline.fileList.map(file => ({
            raw: file.raw,
            file: file.file
          }))
        });
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            // 表单规则校验通过
            await supplierChangeFund(curData).then(res => {
              if (res.success) {
                message("修改成功", { type: "success" });
              } else {
                message(res.message, { type: "error" });
              }
            });
            chores();
          }
        });
      }
    });
  }

  function handleUpdateBalance(row?: FormItemProps) {
    console.log("handleUpdateBalance", row);
  }

  function handleDirectOrder(row?: FormItemProps) {
    console.log("handleDirectOrder", row);
  }

  function handleCheckAccount(row?: FormItemProps) {
    console.log("handleCheckAccount", row);
  }

  function handleSales(row?: FormItemProps) {
    console.log("handleSales", row);
  }

  function handleConfigProduct(row?: FormItemProps) {
    console.log("handleConfigProduct", row);
  }

  function handleBalanceLog(row?: FormItemProps) {
    addDialog({
      title: `余额更新记录`,
      props: {
        supplier_id: row?.id ?? ""
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(balanceLogForm, {
          ref: formRef,
          supplier_id: 0
        }),
      beforeSure: done => {
        done(); // 关闭弹框
        onSearch(); // 刷新表格数据
      }
    });
  }

  function handleUpdataInfo(row?: FormItemProps) {
    addDialog({
      title: `修改供应商信息`,
      props: {
        formInline: {
          id: row?.id ?? 0,
          name: row?.name ?? "",
          dept: row?.dept ?? "",
          phone: row?.phone ?? "",
          email: row?.email ?? "",
          nikename: row?.nickname ?? "",
          remark: row?.remark ?? ""
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(changeInfoForm, {
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
              ...curData
            };
            // 表单规则校验通过
            await updateSupplier(data).then(res => {
              if (res.success) {
                message("修改成功", { type: "success" });
              } else {
                message(res.message, { type: "error" });
              }
            });
            chores();
          }
        });
      }
    });
  }

  return {
    form,
    curRow,
    loading,
    columns,
    rowStyle,
    dataList,
    pagination,
    selectedNum,
    handleBatchChange,
    handleBatchUpdate,
    handleChangeFund,
    handleUpdataInfo,
    handleUpdateBalance,
    handleDirectOrder,
    handleCheckAccount,
    handleSales,
    handleConfigProduct,
    handleBalanceLog,
    onSearch,
    resetForm,
    handleAdd,
    handleDelete,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange,
    supplierItemLists,
    templateNameLists
  };
}
