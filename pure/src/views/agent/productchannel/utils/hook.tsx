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
import type { FormItemProps, CategoryProps } from "../utils/types";
import type { ProductIdName } from "@/api/types";
import type { AgentIdName } from "@/api/types";
import type { SupplierIdName } from "@/api/types";
import { useRoute } from "vue-router";
import { message } from "@/utils/message";
import {
  OperatorListAll,
  SupplyStrategyList,
  BusinessTypeList,
  ProvinceList
} from "@/api/constdata";

export function useCategory(tableRef: Ref) {
  const form = reactive({
    /** 代理商 */
    agent_id: "" as any,
    agent_name: "",
    /** 供货商 */
    supplier_id: "",
    supplier_name: "",
    business_type: "",
    product_id: "",
    product_category: "",
    operator: "",
    agent_product_id: "" as any,
    supplier_product_id: "" as any,
    up_product_id: "",
    status: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const selectedNum = ref(0);
  const agentItemLists = ref([] as AgentIdName[]);
  const supplierItemLists = ref([] as SupplierIdName[]);
  const productCategoryList = ref([] as CategoryProps);
  const productBaseInfoList = ref([] as ProductIdName[]);
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
      label: "通道ID",
      prop: "id"
    },
    {
      label: "产品",
      prop: "product_id",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {row.product_id} <br /> <hr style="border-color: lightgray;" />{" "}
          {row.product_name} <br /> <hr style="border-color: lightgray;" />
          {row.base_price}
        </span>
      )
    },
    {
      label: "产品类别",
      prop: "product_category",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {
            productCategoryList.value.find(
              item => item.id === row.product_category
            )?.category_name
          }
        </span>
      )
    },
    {
      label: "产品运营商",
      prop: "operator",
      minWidth: 100,
      cellRenderer: ({ row }) => (
        <span>
          {OperatorListAll.find(item => item.value === row.operator)?.label}
        </span>
      )
    },
    {
      label: "代理商产品ID",
      prop: "agent_product_id"
    },
    {
      label: "代理商",
      prop: "agent_id",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {row.agent_id} <br /> <hr style="border-color: lightgray;" />{" "}
          {row.agent_name}
        </span>
      )
    },
    {
      label: "代理折扣",
      prop: "agent_discount",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {row.agent_discount}
          <br /> <hr style="border-color: lightgray;" />
          {(row.agent_discount * row.base_price).toFixed(2)}
        </span>
      )
    },
    {
      label: "供货策略",
      prop: "supply_strategy",
      cellRenderer: ({ row }) => (
        <span>
          {
            SupplyStrategyList.find(item => item.value === row.supply_strategy)
              ?.label
          }
        </span>
      )
    },
    {
      label: "供货商产品ID",
      prop: "supplier_product_id"
    },
    {
      label: "供货商",
      prop: "supplier_id",
      minWidth: 150,
      cellRenderer: ({ row }) => (
        <span>
          {row.supplier_id} <br /> <hr style="border-color: lightgray;" />{" "}
          {row.supplier_name}
        </span>
      )
    },
    {
      label: "供货折扣",
      prop: "supplier_discount",
      cellRenderer: ({ row }) => (
        <span>
          {row.supplier_discount}
          <br /> <hr style="border-color: lightgray;" />
          {(row.supplier_discount * row.base_price).toFixed(2)}
        </span>
      )
    },
    {
      label: "上游产品ID",
      prop: "up_product_id",
      minWidth: 100
    },
    {
      label: "执行次数",
      prop: "execute_count"
    },
    {
      label: "权重",
      prop: "weight"
    },
    {
      label: "优先级",
      prop: "priority"
    },
    {
      label: "分省",
      prop: "province",
      cellRenderer: ({ row }) => (
        <span>
          {ProvinceList.find(item => item.value === row.province)?.label}
        </span>
      )
    },
    {
      label: "状态",
      prop: "status",
      cellRenderer: ({ row }) => (
        <span>{row.status === 0 ? "维护" : "上架"}</span>
      )
    },
    {
      label: "是否可通",
      prop: "is_connect",
      cellRenderer: ({ row }) => (
        <span>{row.is_connect === 1 ? "是" : "否"}</span>
      )
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
    if (params.agent_product_id != "") {
      params.agent_product_id = Number(params.agent_product_id);
    }
    if (params.supplier_product_id != "") {
      params.supplier_product_id = Number(params.supplier_product_id);
    }
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
    form.agent_id = id;
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
