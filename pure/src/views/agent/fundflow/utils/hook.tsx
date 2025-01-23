import type { PaginationProps } from "@pureadmin/table";
import {
  getAgentProductChannelList,
  getAgentSimpleList,
  deleteAgentChannel
} from "@/api/agent";
import {
  getProductCategoryList,
  getProductInformationIdAndName
} from "@/api/product";
import { getSupplierSimpleList } from "@/api/supplier";
import { reactive, ref, onMounted, type Ref } from "vue";
import type {
  AgentSimpleItem,
  FormItemProps,
  CategoryProps,
  ProductBaseInfoArray
} from "../utils/types";
import type { SupplierIdName } from "@/api/types";
import { useRoute } from "vue-router";
import { message } from "@/utils/message";
import { FundOperationTypeList } from "@/api/constdata";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    /** 代理商 */
    agent_id: "" as any,
    agent_name: "",
    order_id: "",
    recharge_number: "",
    action: "",
    create_time: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const selectedNum = ref(0);
  const agentItemLists = ref([] as AgentSimpleItem[]);
  const supplierItemLists = ref([] as SupplierIdName[]);
  const productCategoryList = ref([] as CategoryProps);
  const productBaseInfoList = ref([] as ProductBaseInfoArray);
  const route = useRoute();
  const disabled = ref(false);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "流水ID",
      prop: "id"
    },
    {
      label: "代理商ID",
      prop: "agent_id"
    },
    {
      label: "代理商名称",
      prop: "agent_name"
    },
    {
      label: "操作类型",
      prop: "action",
      cellRenderer: ({ row }) => (
        <span>
          {FundOperationTypeList.find(item => item.value === row.action)?.label}
        </span>
      )
    },
    {
      label: "操作金额",
      prop: "amount"
    },
    {
      label: "开始金额",
      prop: "before_amount"
    },
    {
      label: "结束金额",
      prop: "after_amount"
    },
    {
      label: "订单信息",
      prop: "order_id",
      minWidth: 100,
      cellRenderer: ({ row }) => (
        <span>
          {row.order_id} <br /> <hr style="border-color: lightgray;" />{" "}
          {row.recharge_number} <br /> <hr style="border-color: lightgray;" />
          {row.order_create_time}
        </span>
      )
    },
    {
      label: "创建时间",
      prop: "create_time"
    },
    {
      label: "凭证",
      prop: "cert_pic",
      slot: "image"
    },
    {
      label: "备注",
      prop: "remark"
    }
  ];

  function handleSizeChange(val: number) {
    pagination.pageSize = val;
    console.log("handleSizeChange");
    onSearch();
  }

  function handleCurrentChange(val: number) {
    pagination.currentPage = val;
    console.log("handleCurrentChange");
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
    const { data } = await getAgentProductChannelList(params);
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
    var id = form.agent_id;
    formEl.resetFields();
    if (disabled.value) {
      form.agent_id = id;
    }
    console.log("resetForm");
    onSearch();
  };

  /** 高亮当前权限选中行 */
  function rowStyle({ row: { id } }) {
    return {
      cursor: "pointer",
      background: id === curRow.value?.id ? "var(--el-fill-color-light)" : ""
    };
  }

  function openDialog() {}

  async function handleDelete(row?: FormItemProps) {
    var ids = [row.id];
    await deleteAgentChannel(ids).then(res => {
      if (res.success) {
        message("删除通道成功", { type: "success" });
      } else {
        message(res.message, { type: "error" });
      }
    });
    console.log("handleDelete");
    onSearch();
  }

  onMounted(async () => {
    // 动态获取产品类别列表
    const response = await getAgentSimpleList();
    agentItemLists.value = response.data;
    const supplierResponse = await getSupplierSimpleList();
    supplierItemLists.value = supplierResponse.data;
    disabled.value = false;
    if (route.query.agent_id !== undefined) {
      form.agent_id = Number(route.query.agent_id);
      form.agent_name = String(route.query.agent_name);
      disabled.value = true;
    }
    var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
    const response2 = await getProductCategoryList(requestData);
    productCategoryList.value = response2.data.list;
    const response3 = await getProductInformationIdAndName();
    productBaseInfoList.value = response3.data;
    onSearch();
  });

  return {
    form,
    disabled,
    curRow,
    loading,
    selectedNum,
    columns,
    rowStyle,
    dataList,
    pagination,
    agentItemLists,
    supplierItemLists,
    productCategoryList,
    productBaseInfoList,
    onSearch,
    openDialog,
    resetForm,
    handleSizeChange,
    handleDelete,
    handleCurrentChange,
    handleSelectionChange,
    onSelectionCancel
  };
}
