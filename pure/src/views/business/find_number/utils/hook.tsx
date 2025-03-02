import { OperatorListAll, OrderStatusList } from "@/api/constdata";
import { orderFindPhone } from "@/api/order";
import type { PaginationProps } from "@pureadmin/table";
import { reactive, ref } from "vue";

export function useCategory() {
  const form = reactive({
    recharge_number: "",
    create_time: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const loading = ref(false);

  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "充值号码",
      prop: "recharge_number"
    },
    {
      label: "面值",
      prop: "base_price"
    },
    {
      label: "运营商",
      prop: "operator",
      cellRenderer: ({ row }) => (
        <span>
          {OperatorListAll.find(item => item.value === row.operator)?.label}
        </span>
      )
    },
    {
      label: "归属地",
      prop: "location"
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
      label: "创建时间",
      prop: "create_time"
    },
    {
      label: "完成时间",
      prop: "finish_time"
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

  async function onSearch() {
    loading.value = true;
    var params = { ...form, ...pagination };
    const { data } = await orderFindPhone(params);
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

  return {
    form,
    curRow,
    loading,
    columns,
    rowStyle,
    dataList,
    pagination,
    onSearch,
    resetForm,
    handleSizeChange,
    handleCurrentChange
  };
}
