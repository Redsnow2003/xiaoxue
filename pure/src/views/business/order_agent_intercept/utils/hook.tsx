import type { ProductIdName } from "@/api/types";
import type { AgentIdName } from "@/api/types";
import type { PaginationProps } from "@pureadmin/table";
import { type Ref, reactive, ref, onMounted } from "vue";

import { getAgentSimpleList } from "@/api/agent";
import { getInterceptOrderInfo } from "@/api/order";
import { getProductInformationIdAndName } from "@/api/product";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    down_id: "",
    recharge_number: "",
    agent_id: "",
    product_id: "",
    create_time: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const selectedNum = ref(0);
  const loading = ref(true);
  const agentItemLists = ref([] as AgentIdName[]);
  const productBaseInfoList = ref([] as ProductIdName[]);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "ID",
      prop: "id"
    },
    {
      label: "下游订单号",
      prop: "down_id"
    },
    {
      label: "代理商",
      prop: "agent_id",
      cellRenderer: ({ row }) => (
        <span>
          {row.agent_id}
          <br />
          <hr style="border-color: lightgray;" />
          {row.agent_name}
        </span>
      )
    },
    {
      label: "产品ID|产品名称",
      prop: "product_id",
      cellRenderer: ({ row }) => (
        <span>
          {row.product_id}
          <br />
          <hr style="border-color: lightgray;" />
          {row.product_name}
        </span>
      )
    },
    {
      label: "充值号码",
      prop: "recharge_number"
    },
    {
      label: "校验价格",
      prop: "check_price"
    },
    {
      label: "请求IP",
      prop: "request_ip"
    },
    {
      label: "请求时间",
      prop: "request_time"
    },
    {
      label: "请求参数",
      prop: "request_params"
    },
    {
      label: "我方返回信息",
      prop: "response_information"
    },
    {
      label: "返回时间",
      prop: "response_time"
    },
    {
      label: "备注",
      prop: "remark"
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
    const { data } = await getInterceptOrderInfo(params);
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
    const response3 = await getProductInformationIdAndName();
    productBaseInfoList.value = response3.data;
  });

  function handleExportCSV() {
    // 导出CSV文件的逻辑
    console.log("handleExportCSV");
  }

  return {
    form,
    curRow,
    loading,
    columns,
    rowStyle,
    dataList,
    pagination,
    handleExportCSV,
    onSearch,
    resetForm,
    handleSizeChange,
    handleCurrentChange,
    onSelectionCancel,
    handleSelectionChange,
    agentItemLists,
    productBaseInfoList
  };
}
