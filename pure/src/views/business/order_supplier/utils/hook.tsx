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
  OrederSupplierItemProps,
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
  OperatorListAll,
  SupplierOrderStatusList
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
import { ElMessageBox } from "element-plus";

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
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {row.order_id}
          <br />
          <hr style="border-color: lightgray;" />
          {row.id}
          <br />
          <hr style="border-color: lightgray;" />
          {row.up_id}
        </span>
      )
    },
    {
      label: "代理商",
      prop: "agent_id",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span
          style="cursor: pointer;color: #409EFF; text-decoration: underline;"
          onClick={() => gotoAgentInfo(row)}
        >
          {row.agent_id}
          <br />
          <hr style="border-color: lightgray;" />
          {row.agent_name}
        </span>
      )
    },
    {
      label: "代理折扣|数量|订单总价",
      prop: "agent_discount",
      minWidth: 150,
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
      label: "供货商",
      prop: "supplier_id",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span
          style="cursor: pointer;color: #409EFF; text-decoration: underline;"
          onClick={() => gotoAgentInfo(row)}
        >
          {row.supplier_id}
          <br />
          <hr style="border-color: lightgray;" />
          {row.supplier_name}
        </span>
      )
    },
    {
      label: "产品ID|产品名称|基础价|运营商",
      prop: "product_id",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span
          style="cursor: pointer;color: #409EFF; text-decoration: underline;"
          onClick={() => gotoProductInfo(row)}
        >
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
      label: "供货折扣|数量|供货总价",
      prop: "supplier_discount",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {row.supplier_discount}
          <br />
          <hr style="border-color: lightgray;" />
          {row.count}
          <br />
          <hr style="border-color: lightgray;" />
          {(row.base_price * row.supplier_discount * row.count).toFixed(2)}
        </span>
      )
    },
    {
      label: "充值号码|运营商|归属地",
      prop: "recharge_number",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {row.recharge_number}
          <br />
          <hr style="border-color: lightgray;" />
          {row.operator}
          <br />
          <hr style="border-color: lightgray;" />
          {row.location}
        </span>
      )
    },
    {
      label: "订单时间",
      minWidth: 150,
      prop: "order_time"
    },
    {
      label: "供货单状态",
      prop: "status",
      cellRenderer: ({ row }) => (
        <span>
          {
            SupplierOrderStatusList.find(item => item.value === row.status)
              ?.label
          }
        </span>
      )
    },
    {
      label: "创建时间|完成时间",
      prop: "create_time",
      minWidth: 150,
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
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>{timeDiff(row.create_time, row.finish_time)}</span>
      )
    },
    {
      label: "上游信息",
      prop: "up_information"
    },
    {
      label: "更新时间",
      prop: "update_time"
    },
    {
      label: "备注",
      prop: "remark"
    },
    {
      label: "备用通道重试",
      prop: "is_backup",
      cellRenderer: ({ row }) => (row.is_backup ? "是" : "否")
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

  //点击跳转到对应的代理商页面
  function gotoAgentInfo(row?: OrederSupplierItemProps) {
    const router = useRouter();
    console.log("handleAgentproduct", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到产品配置页面
      router.push({});
    } else {
      console.error("Row or Row ID is missing");
    }
  }

  //点击跳转到对应产品配置页面
  function gotoProductInfo(row?: OrederSupplierItemProps) {
    const router = useRouter();
    console.log("handleAgentproduct", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到产品配置页面
      router.push({});
    } else {
      console.error("Row or Row ID is missing");
    }
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

  async function handleLookOrder(row?: OrederSupplierItemProps) {
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

  function handleQueryUpOrder(row?: OrederSupplierItemProps) {
    // 查看供应商订单的逻辑
    addDialog({
      title: `上游订单信息`,
      props: {
        up_id: row.up_id,
        supplier_id: row.supplier_id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(OrderUpInfoForm, { up_id: null, supplier_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleFailureToSuccess(row?: OrederSupplierItemProps) {
    // 通知状态的逻辑
    ElMessageBox.confirm(`是否确认将供货单${row.id}状态改为成功`, "确认", {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
      type: "warning"
    })
      .then(() => {
        var supplierOrderId = {
          id: row.id
        };
        supplierOrderFailToSuccess(supplierOrderId).then(() => {
          onSearch();
        });
      })
      .catch(() => {});
  }

  function handleChangeRemark(row?: OrederSupplierItemProps) {
    // 修改备注的逻辑
    addDialog({
      title: `修改订单备注`,
      props: {
        formInline: {
          /** 订单ID */
          order_id: row.id,
          /** 备注 */
          remark: row.remark
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

  function handleBackupSubmitLog(row?: OrederSupplierItemProps) {
    // 备份日志的逻辑
    addDialog({
      title: `备用通道重新提交记录`,
      props: {
        order_id: row.order_id
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

  function handleOrderSubmitLog(row?: OrederSupplierItemProps) {
    // 查询日志的逻辑
    addDialog({
      title: `提单日志`,
      props: {
        supplier_order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(OrderSubmitLogForm, { ref: formRef, supplier_order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleOrderQueryLog(row?: OrederSupplierItemProps) {
    // 通知日志的逻辑
    addDialog({
      title: `查单日志`,
      props: {
        supplier_order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(OrderQueryLogForm, { ref: formRef, supplier_order_id: null }),
      beforeSure: done => {
        done(); // 关闭弹框
      }
    });
  }

  function handleOrderCallbackLog(row?: OrederSupplierItemProps) {
    // 通知日志的逻辑
    addDialog({
      title: `回调日志`,
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

  function handleOrderCancelLog(row?: OrederSupplierItemProps) {
    // 通知日志的逻辑
    addDialog({
      title: `撤单日志`,
      props: {
        supplier_order_id: row.id
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(OrderCancelLogForm, { ref: formRef, supplier_order_id: null }),
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
    handleBackupSubmitLog,
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

  function handleSupplierProduct(row?: OrederSupplierItemProps) {
    console.log("handleAgentChannel", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到产品配置页面
      router.push({
        path: "/supplier/product/index",
        query: {
          product_id: row.product_id,
          product_name: row.product_name
        }
      });
    } else {
    }
  }

  return {
    handleSupplierProduct
  };
}
export { useProductHandlers };
