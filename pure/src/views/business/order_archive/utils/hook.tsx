import BatchStatusForm from "../form/batchstatus.vue";
import BatchBackupForm from "../form/batchbackup.vue";
import BatchTimeoutForm from "../form/batchtimeout.vue";
import SupplierOrderForm from "../form/supplierorder.vue";
import ChangeRemarkForm from "../form/changeremark.vue";
import BackupLogForm from "../form/backuplog.vue";
import OrderSubmitLogForm from "../form/ordersubmitlog.vue";
import OrderQueryLogForm from "../form/orderquerylog.vue";
import OrderNotifyLogForm from "../form/ordernotifylog.vue";
import { addDialog } from "@/components/ReDialog";
import type {
  FormItemProps,
  CategoryProps,
  BatchStatusProps,
  BatchBackupProps,
  BatchTimeoutProps,
  ChangeRemarkProps
} from "../utils/types";
import type { ProductIdName } from "@/api/types";
import type { AgentIdName } from "@/api/types";
import { timeDiff } from "@/api/utils";
import type { PaginationProps } from "@pureadmin/table";
import { deviceDetection, getKeyList } from "@pureadmin/utils";
import { type Ref, reactive, ref, onMounted, h } from "vue";
import { useRouter } from "vue-router";
import {
  BusinessTypeList,
  NotifyStatusList,
  OperatorListAll,
  OrderStatusList
} from "@/api/constdata";
import { getAgentSimpleList } from "@/api/agent";
import {
  getOrderList,
  batchBackupCancel,
  batchBackupSubmit,
  batchOrderCancel,
  batchOrderManual,
  batchOrderNotice,
  batchOrderTimeout,
  batchUpdateOrderStatusRemark,
  updateOrderRemark
} from "@/api/order";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    business_type: "",
    id: "",
    down_id: "",
    notify_status: "",
    recharge_number: "",
    agent_id: "",
    product_category: "",
    product_id: "",
    base_price: "",
    remark: "",
    status: "",
    is_timeout: "",
    is_cancel: "",
    location: "",
    special_params: "",
    create_time: "",
    finish_time: ""
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const selectedNum = ref(0);
  const loading = ref(true);
  const agentItemLists = ref([] as AgentIdName[]);
  const productBaseInfoList = ref([] as ProductIdName[]);
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
      label: "业务类型",
      prop: "business_type",
      cellRenderer: ({ row }) => (
        <span>
          {
            BusinessTypeList.find(item => item.value === row.business_type)
              ?.label
          }
        </span>
      )
    },
    {
      label: "订单号|下游订单号",
      prop: "id",
      cellRenderer: ({ row }) => (
        <span>
          {row.id}
          <br />
          {row.down_id}
        </span>
      )
    },
    {
      label: "代理商",
      prop: "agent_id",
      cellRenderer: ({ row }) => (
        <span>
          {row.agent_id}
          <br />
          {row.agent_name}
        </span>
      )
    },
    {
      label: "产品ID|产品名称|基础价|运营商",
      prop: "product_id",
      cellRenderer: ({ row }) => (
        <span>
          {row.product_id}
          <br />
          <hr style="border-color: lightgray;" />
          {row.product_name}
          <br />
          <hr style="border-color: lightgray;" />
          {row.base_price}
          <br />
          <hr style="border-color: lightgray;" />
          {OperatorListAll.find(item => item.value === row.operator)?.label}
        </span>
      )
    },
    {
      label: "代理折扣|数量|订单总价",
      prop: "agent_discount",
      cellRenderer: ({ row }) => (
        <span>
          {row.agent_discount}
          <br />
          <hr style="border-color: lightgray;" />
          {row.count}
          <br />
          <hr style="border-color: lightgray;" />
          {(row.base_price * row.agent_discount * row.count).toFixed(2)}
        </span>
      )
    },
    {
      label: "充值号码|运营商|归属地",
      prop: "recharge_number",
      cellRenderer: ({ row }) => (
        <span>
          {row.recharge_number}
          <br />
          <hr style="border-color: lightgray;" />
          {row.operator}
          <br />
          <hr style="border-color: lightgray;" />
          {location}
        </span>
      )
    },
    {
      label: "订单状态",
      prop: "status",
      cellRenderer: ({ row }) => (
        <span>
          {OrderStatusList.find(item => item.value === row.status)?.label}
        </span>
      )
    },
    {
      label: "是否超时",
      prop: "is_timeout",
      cellRenderer: ({ row }) => (
        <span>{row.is_timeout === 1 ? "是" : "否"}</span>
      )
    },
    {
      label: "超时时间(秒)",
      prop: "timeout"
    },
    {
      label: "创建时间|完成时间",
      prop: "create_time",
      cellRenderer: ({ row }) => (
        <span>
          {row.create_time}
          <br />
          <hr style="border-color: lightgray;" />
          {row.finish_time}
        </span>
      )
    },
    {
      label: "用时",
      prop: "finish_time",
      cellRenderer: ({ row }) => (
        <span>{timeDiff(row.create_time, row.finish_time)}</span>
      )
    },
    {
      label: "通知状态",
      prop: "notify_status",
      cellRenderer: ({ row }) => (
        <span>
          {
            NotifyStatusList.find(item => item.value === row.notify_status)
              ?.label
          }
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
    const { data } = await getOrderList(params);
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
    // 动态获取产品类别列表
    const response = await getAgentSimpleList();
    agentItemLists.value = response.data;
  });

  function handleExportCSV() {
    // 导出CSV文件的逻辑
    console.log("handleExportCSV");
  }

  function handleBatchStatus() {
    // 批量缓存的逻辑
    addDialog({
      title: `批量手动核销订单`,
      props: {
        formInline: {
          /** 状态 */
          status: 0,
          /** 备注 */
          remark: ""
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(BatchStatusForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const formData = options.props.formInline as BatchStatusProps;
        const curSelected = tableRef.value.getTableRef().getSelectionRows();
        var ids = getKeyList(curSelected, "id");
        const curData = {
          ...formData,
          ids: ids
        };
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            await batchUpdateOrderStatusRemark(curData);
            chores();
          }
        });
      }
    });
  }

  async function handleBatchNotify() {
    // 批量通知的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await batchOrderNotice(ids);
    onSearch();
  }

  function handleBatchBackup() {
    // 批量备份的逻辑
    addDialog({
      title: `批量备用通道重新提交`,
      props: {
        formInline: {
          /** 次数 */
          count: 1,
          /** 时间间隔 */
          interval: 10
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(BatchBackupForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const formData = options.props.formInline as BatchBackupProps;
        const curSelected = tableRef.value.getTableRef().getSelectionRows();
        var ids = getKeyList(curSelected, "id");
        const curData = {
          ...formData,
          ids: ids
        };
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            await batchBackupSubmit(curData);
            chores();
          }
        });
      }
    });
  }

  async function handleBatchCancelBackup() {
    // 批量取消备份的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await batchBackupCancel(ids);
    onSearch();
  }

  function handleBatchTimeout() {
    // 批量超时的逻辑
    addDialog({
      title: `批量设置订单超时`,
      props: {
        formInline: {
          type: 0,
          seconds: 0,
          time: "",
          create_after: 0
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(BatchTimeoutForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const formData = options.props.formInline as BatchTimeoutProps;
        const curSelected = tableRef.value.getTableRef().getSelectionRows();
        var ids = getKeyList(curSelected, "id");
        const curData = {
          ...formData,
          ids: ids
        };
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            await batchOrderTimeout(curData);
            chores();
          }
        });
      }
    });
  }

  async function handleBatchCancel() {
    // 批量取消的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await batchOrderCancel(ids);
    onSearch();
  }

  async function handleBatchManual() {
    // 批量手动处理的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await batchOrderManual(ids);
    onSearch();
  }

  function handleLookSupplierOrder(row?: FormItemProps) {
    // 查看供应商订单的逻辑
    addDialog({
      title: `相关供货单列表`,
      props: {
        order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(SupplierOrderForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleNotifyStatus(row?: FormItemProps) {
    // 通知状态的逻辑
    console.log("handleNotifyStatus", row);
  }

  function handleChangeRemark(row?: FormItemProps) {
    // 修改备注的逻辑
    addDialog({
      title: `修改订单备注`,
      props: {
        formInline: {
          /** 订单ID */
          order_id: row.id,
          /** 备注 */
          remark: ""
        }
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(ChangeRemarkForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const formData = options.props.formInline as ChangeRemarkProps;
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            await updateOrderRemark(formData);
            chores();
          }
        });
      }
    });
  }

  function handleBackupLog(row?: FormItemProps) {
    // 备份日志的逻辑
    addDialog({
      title: `备用通道重新提交记录`,
      props: {
        order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(BackupLogForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleSubmitLog(row?: FormItemProps) {
    // 提交日志的逻辑
    addDialog({
      title: `提单日志`,
      props: {
        order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(OrderSubmitLogForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleQueryLog(row?: FormItemProps) {
    // 查询日志的逻辑
    addDialog({
      title: `提单日志`,
      props: {
        order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(OrderQueryLogForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleNotifyLog(row?: FormItemProps) {
    // 通知日志的逻辑
    addDialog({
      title: `提单日志`,
      props: {
        order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(OrderNotifyLogForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
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
    handleExportCSV,
    handleBatchStatus,
    handleBatchNotify,
    handleBatchBackup,
    handleBatchCancelBackup,
    handleBatchTimeout,
    handleBatchCancel,
    handleBatchManual,
    handleLookSupplierOrder,
    handleNotifyStatus,
    handleChangeRemark,
    handleBackupLog,
    handleSubmitLog,
    handleQueryLog,
    handleNotifyLog,
    onSearch,
    resetForm,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange,
    agentItemLists,
    productBaseInfoList,
    productCategoryList
  };
}

function useProductHandlers() {
  const router = useRouter();

  function handleAgentproduct(row?: FormItemProps) {
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
    handleAgentproduct
  };
}
export { useProductHandlers };
