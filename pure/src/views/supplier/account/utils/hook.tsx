import addForm from "../form/addsupplier.vue";
import batchChangeForm from "../form/batchchange.vue";
import balanceLogForm from "../form/balancelog.vue";
import changeInfoForm from "../form/changeinfo.vue";
import changeFundForm from "../form/changefund.vue";
import templateForm from "../form/template.vue";
import fundLogForm from "../form/fundlog.vue";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type { FormItemProps } from "../utils/types";
import type { SupplierIdName } from "@/api/types";
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
import { useRouter } from "vue-router";

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
  const supplierItemLists = ref([] as SupplierIdName[]);
  const templateNameLists = ref([]);
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
      prop: "up_template",
      cellRenderer: ({ row }) => (
        <span>
          {
            <el-button type="text" onClick={() => handleConfigTemplate(row)}>
              {row.up_template}
            </el-button>
          }
        </span>
      )
    },
    {
      label: "备注",
      prop: "remark"
    },
    {
      label: "状态",
      prop: "status",
      cellRenderer: ({ row }) => (
        <span>{row.status === 0 ? "维护" : "上架"}</span>
      )
    },
    {
      label: "状态信息",
      prop: "status_info"
    },
    {
      label: "操作",
      fixed: "right",
      width: 100,
      slot: "operation"
    }
  ];

  async function handleDelete(row) {
    var ids = [row.id];
    await deleteSupplier(ids).then(res => {
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
        formInline: {
          name: "",
          dept: "",
          phone: "",
          email: "",
          nickname: "",
          remark: ""
        }
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
              ...curData,
              our_balance: 0,
              up_balance: 0,
              up_balance_update_time: null,
              status: 1
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
        formInline: {
          status: 1
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(batchChangeForm, {
          formInline: null
        }),
      beforeSure: async (done, { options }) => {
        const curSelected = tableRef.value.getTableRef().getSelectionRows();
        var ids = getKeyList(curSelected, "id");
        const curData = {
          status: options.props.formInline.status as number,
          ids: ids
        };
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
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
          supplier_id: row?.id ?? "",
          supplier_name: row?.name ?? "",
          fund_action: "add",
          amount: 0,
          confirm_amount: 0,
          file: "",
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
        const curData = options.props.formInline;
        const data = {
          ...curData,
          amount: Number(curData.amount),
          confirmAmount: Number(curData.confirmAmount)
        };
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async (valid: any) => {
          if (valid) {
            // 表单规则校验通过
            await supplierChangeFund(data).then(res => {
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

  function handleConfigTemplate(row?: FormItemProps) {
    addDialog({
      title: `上游模板`,
      props: {
        formInline: {
          up_template: row?.up_template ?? "",
          template_json: row?.template_json ?? ""
        }
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(templateForm, {
          formInline: null
        }),
      beforeSure: async (done, { options }) => {
        const curData = {
          ...row,
          ...options.props.formInline
        };
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        await updateSupplier(curData).then(res => {
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

  function handleChangeFundLog(row?: FormItemProps) {
    addDialog({
      title: `资金操作记录`,
      props: {
        supplier_id: row?.id ?? ""
      },
      width: "70%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(fundLogForm, {
          ref: formRef,
          supplier_id: 0
        }),
      beforeSure: done => {
        done(); // 关闭弹框
        onSearch(); // 刷新表格数据
      }
    });
  }

  function handleBalanceLog(row?: FormItemProps) {
    addDialog({
      title: `余额更新记录`,
      props: {
        supplier_id: row?.id ?? ""
      },
      width: "70%",
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
          remark: row?.remark ?? "",
          status: row?.status ?? 0
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
        const curData = {
          ...row,
          ...options.props.formInline
        };
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            var data = {
              ...curData,
              status: curData.status as number
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
    handleChangeFundLog,
    handleUpdataInfo,
    handleUpdateBalance,
    handleDirectOrder,
    handleCheckAccount,
    handleSales,
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

function useProductHandlers() {
  const router = useRouter();
  function handleConfigProduct(row?: FormItemProps) {
    console.log("handleConfigProduct", row);

    if (row && row.id) {
      // 携带参数 row.id 跳转到产品配置页面
      router.push({
        path: "/supplier/product/index",
        query: { supplier_id: row.id }
      });
    } else {
      console.error("Row or Row ID is missing");
    }
  }
  return { handleConfigProduct };
}
export { useProductHandlers };
