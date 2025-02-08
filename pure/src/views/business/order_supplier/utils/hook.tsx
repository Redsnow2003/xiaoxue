import BatchCheckTemplateForm from "../form/batchchecktemplate.vue";
import BatchUrgeTemplateForm from "../form/batchurgetemplate.vue";
import BatchBackupSubmitForm from "../form/batchbackup.vue";
import OrderInfoForm from "../form/orderinfo.vue";
import OrderUpInfoForm from "../form/orderupinfo.vue";
import ChangeRemarkForm from "../form/changeremark.vue";
import BackupLogForm from "../form/backuplog.vue";
import OrderSubmitLogForm from "../form/ordersubmitlog.vue";
import OrderQueryLogForm from "../form/orderquerylog.vue";
import OrderCallbackLogForm from "../form/ordercallbacklog.vue";
import OrderCancelLogForm from "../form/ordercancellog.vue";
import { addDialog } from "@/components/ReDialog";
import type {
  FormItemProps,
  CategoryProps,
  ChangeRemarkProps,
  BatchBackupProps
} from "../utils/types";
import type { ProductIdName } from "@/api/types";
import type { AgentIdName } from "@/api/types";
import type { SupplierIdName } from "@/api/types";
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
  getSupplierOrderList,
  batchSupplierCancel,
  batchSupplierManualCheck,
  batchSupplierAutoCheck,
  updateSupplierOrderRemark,
  batchSupplierBackupSubmit,
  supplierOrderFailToSuccess
} from "@/api/order";
import { getSupplierSimpleList } from "@/api/supplier";
import {
  getProductCategoryList,
  getProductInformationIdAndName
} from "@/api/product";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    business_type: "",
    id: "",
    order_id: "",
    up_id: "",
    recharge_number: "",
    agent_id: "",
    product_category: "",
    product_id: "",
    base_price: "",
    remark: "",
    status: "",
    supplier: "",
    order_time: "",
    create_time: "",
    finish_time: ""
  });
  const curRow = ref();
  const formRef = ref();
  const dataList = ref([]);
  const selectedNum = ref(0);
  const loading = ref(true);
  const agentItemLists = ref([] as AgentIdName[]);
  const supplierItemLists = ref([] as SupplierIdName[]);
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
      label: "订单号|供货单号|上游单号",
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
    const { data } = await getSupplierOrderList(params);
    console.log(data);
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
    const supplierResponse = await getSupplierSimpleList();
    supplierItemLists.value = supplierResponse.data;
    var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
    const response2 = await getProductCategoryList(requestData);
    productCategoryList.value = response2.data.list;
    const response3 = await getProductInformationIdAndName();
    productBaseInfoList.value = response3.data;
  });

  function handleExportCSV() {
    // 导出CSV文件的逻辑
    console.log("handleExportCSV");
  }

  async function handleBatchAutoCheckOrder() {
    // 批量缓存的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await batchSupplierAutoCheck(ids);
    onSearch();
  }

  async function handleBatchManualCheckOrder() {
    // 批量通知的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await batchSupplierManualCheck(ids);
    onSearch();
  }

  function handleBatchCheckTemplate() {
    // 批量备份的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    addDialog({
      title: `批量核单客服模板`,
      props: {
        ids: ids
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(BatchCheckTemplateForm),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  async function handleBatchUrgeTemplate() {
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    addDialog({
      title: `批量催单客服模板`,
      props: {
        ids: ids
      },
      width: "30%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(BatchUrgeTemplateForm),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleBatchBackupSubmit() {
    // 批量超时的逻辑
    addDialog({
      title: `批量备用通道重新提交`,
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
        h(BatchBackupSubmitForm, { ref: formRef, formInline: null }),
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
            await batchSupplierBackupSubmit(curData);
            chores();
          }
        });
      }
    });
  }

  async function handleBatchCancelOrder() {
    // 批量取消的逻辑
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await batchSupplierCancel(ids);
    onSearch();
  }

  async function handleLookOrder(row?: FormItemProps) {
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
      contentRenderer: () => h(OrderInfoForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleQueryUpOrder(row?: FormItemProps) {
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
        h(OrderUpInfoForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  async function handleFailureToSuccess(row?: FormItemProps) {
    var supplierOrderId = {
      supplier_order_id: row.id
    };
    await supplierOrderFailToSuccess(supplierOrderId);
    onSearch();
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
            await updateSupplierOrderRemark(formData);
            chores();
          }
        });
      }
    });
  }

  function handleBackupSubmit(row?: FormItemProps) {
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

  function handleSupplierProduct(row?: FormItemProps) {
    // 提交日志的逻辑
    console.log("handleSupplierProduct", row);
  }

  function handleOrderSubmitLog(row?: FormItemProps) {
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
        h(OrderSubmitLogForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleOrderQueryLog(row?: FormItemProps) {
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
        h(OrderQueryLogForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleOrderCallbackLog(row?: FormItemProps) {
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
        h(OrderCallbackLogForm, { ref: formRef, order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleOrderCancelLog(row?: FormItemProps) {
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
        h(OrderCancelLogForm, { ref: formRef, order_id: null }),
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
    handleBatchAutoCheckOrder,
    handleBatchManualCheckOrder,
    handleBatchCheckTemplate,
    handleBatchUrgeTemplate,
    handleBatchBackupSubmit,
    handleBatchCancelOrder,
    handleLookOrder,
    handleChangeRemark,
    handleQueryUpOrder,
    handleFailureToSuccess,
    handleBackupSubmit,
    handleSupplierProduct,
    handleOrderSubmitLog,
    handleOrderQueryLog,
    handleOrderCancelLog,
    handleOrderCallbackLog,
    onSearch,
    resetForm,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange,
    agentItemLists,
    supplierItemLists,
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
