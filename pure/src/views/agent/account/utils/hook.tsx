import addForm from "../form/addagent.vue";
import batchChangeForm from "../form/batchchange.vue";
import changeInfoForm from "../form/changeinfo.vue";
import changeFundForm from "../form/changefund.vue";
import fundLogForm from "../form/fundlog.vue";
import createAccountForm from "../form/createaccount.vue";
import { message } from "@/utils/message";
import { addDialog } from "@/components/ReDialog";
import type { FormItemProps, AgentSimpleItem } from "../utils/types";
import type { PaginationProps } from "@pureadmin/table";
import { deviceDetection, getKeyList } from "@pureadmin/utils";
import {
  getAgentSimpleList,
  getAgentList,
  addAgent,
  updateAgent,
  deleteAgent,
  changeAgentFund,
  batchUpdateAgentStatus
} from "@/api/agent";
import { type Ref, reactive, ref, onMounted, h } from "vue";
import { useRouter } from "vue-router";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    /** 代理商 */
    id: "",
    /** 通知方式 */
    notification_method: "",
    /** 状态 */
    status: ""
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const selectedNum = ref(0);
  const loading = ref(true);
  const agentItemLists = ref([] as AgentSimpleItem[]);
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
      label: "代理商ID",
      prop: "id"
    },
    {
      label: "代理商名称",
      prop: "name"
    },
    {
      label: "账户余额(元)",
      prop: "fund_balance"
    },
    {
      label: "授信余额(元)",
      prop: "credit_balance"
    },
    {
      label: "冻结金额(元)",
      prop: "frozen_amount"
    },
    {
      label: "缓存可用金额(元)",
      prop: "cache_amount"
    },
    {
      label: "密钥",
      prop: "secret_key"
    },
    {
      label: "通知方式",
      prop: "notification_method",
      cellRenderer: ({ row }) => (
        <span>{row.notification_method === 0 ? "可靠通知" : "广播通知"}</span>
      )
    },
    {
      label: "状态",
      prop: "status",
      cellRenderer: ({ row }) => (
        <span>
          {row.status === 0 ? (
            <div style="color:red">维护</div>
          ) : (
            <div style="color:green">上架</div>
          )}
        </span>
      )
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

  async function handleDelete(row?: FormItemProps) {
    //添加确认消息框

    var ids = [row.id];
    await deleteAgent(ids).then(res => {
      if (res.success) {
        message("删除代理商成功", { type: "success" });
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
    const { data } = await getAgentList(params);
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
      title: `新增代理商`,
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
            var data = {
              ...curData,
              notification_method: 0,
              fund_balance: 0,
              credit_balance: 0,
              frozen_amount: 0,
              cache_amount: 0,
              status: 1
            };
            // 表单规则校验通过
            await addAgent(data).then(res => {
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
    const response = await getAgentSimpleList();
    agentItemLists.value = response.data;
  });

  function handleBatchChange() {
    addDialog({
      title: `批量修改代理商状态`,
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
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        // 表单规则校验通过
        await batchUpdateAgentStatus(curData).then(res => {
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

  function handleBatchUpdate() {}

  function handleChangeFund(row?: FormItemProps) {
    addDialog({
      title: `操作资金`,
      props: {
        formInline: {
          agent_id: row?.id ?? "",
          agent_name: row?.name ?? "",
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
            await changeAgentFund(data).then(res => {
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

  function handleChangeFundLog(row?: FormItemProps) {
    addDialog({
      title: `资金操作记录`,
      props: {
        agent_id: row?.id ?? ""
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

  function handleCreateAccount(row?: FormItemProps) {
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
        h(createAccountForm, {
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
      title: `修改代理商信息`,
      props: {
        formInline: {
          id: row?.id ?? 0,
          name: row?.name ?? "",
          nickname: row?.nickname ?? "",
          dept: row?.dept ?? "",
          phone: row?.phone ?? "",
          email: row?.email ?? "",
          secret_key: row?.secret_key ?? "",
          notification_address: row?.notification_address ?? "",
          notification_method: row?.notification_method ?? 0,
          customer: row?.customer ?? "",
          remark: row?.remark ?? "",
          status: row?.status ?? 0,
          fund_balance: row?.fund_balance ?? 0,
          credit_balance: row?.credit_balance ?? 0,
          frozen_amount: row?.frozen_amount ?? 0,
          cache_amount: row?.cache_amount ?? 0
        }
      },
      width: "40%",
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
            await updateAgent(data).then(res => {
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
    handleCreateAccount,
    onSearch,
    resetForm,
    handleAdd,
    handleDelete,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange,
    agentItemLists
  };
}

function useProductHandlers() {
  const router = useRouter();
  function handleWhitelist(row?: FormItemProps) {
    console.log("handleWhitelist", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到白名单配置页面
      router.push({
        path: "/agent/whitelist/index",
        query: {
          agent_id: row.id,
          agent_name: row.name
        }
      });
    } else {
      console.error("Row or Row ID is missing");
    }
  }

  function handleProductConfig(row?: FormItemProps) {
    console.log("handleProductConfig", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到产品配置页面
      router.push({
        path: "/agent/product/index",
        query: {
          agent_id: row.id,
          agent_name: row.name
        }
      });
    } else {
      console.error("Row or Row ID is missing");
    }
  }

  function handleAgentChannel(row?: FormItemProps) {
    console.log("handleAgentChannel", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到产品配置页面
      router.push({
        path: "/agent/channel/index",
        query: {
          agent_id: row.id,
          agent_name: row.name
        }
      });
    } else {
      console.error("Row or Row ID is missing");
    }
  }

  function handleAgentProductChannel(row?: FormItemProps) {
    console.log("handleAgentChannel", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到产品配置页面
      router.push({
        path: "/agent/productchannel/index",
        query: {
          agent_id: row.id,
          agent_name: row.name
        }
      });
    } else {
      console.error("Row or Row ID is missing");
    }
  }

  return {
    handleWhitelist,
    handleProductConfig,
    handleAgentChannel,
    handleAgentProductChannel
  };
}
export { useProductHandlers };
