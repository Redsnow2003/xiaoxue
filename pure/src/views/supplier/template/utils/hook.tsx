import type { PaginationProps } from "@pureadmin/table";
import { getSupplierTemplateList } from "@/api/supplier";
import { reactive, ref, onMounted } from "vue";

export function useCategory() {
  const form = reactive({
    /** 名称 */
    name: "",
    /** 提单路径 */
    submit_address: "",
    /** 查单路径 */
    query_address: "",
    /** 余额路径 */
    balance_address: "",
    /** 备注 */
    remark: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "模板ID",
      prop: "id"
    },
    {
      label: "模板名称",
      prop: "name"
    },
    {
      label: "回调地址(全局)",
      prop: "global_cb_address"
    },
    {
      label: "需要商品编码",
      prop: "is_need_product_id",
      cellRenderer: ({ row }) => (
        <span>{row.is_need_product_id === 1 ? "是" : "否"}</span>
      )
    },
    {
      label: "需要绑定回调地址",
      prop: "is_bind_callback_address",
      cellRenderer: ({ row }) => (
        <span>{row.is_bind_callback_address === 1 ? "是" : "否"}</span>
      )
    },
    {
      label: "支持处理供货金额不一致",
      prop: "is_support_inconsistent",
      cellRenderer: ({ row }) => (
        <span>{row.is_support_inconsistent === 1 ? "是" : "否"}</span>
      )
    },
    {
      label: "是否支持撤单",
      prop: "is_support_cancel",
      cellRenderer: ({ row }) => (
        <span>{row.is_support_cancel === 1 ? "是" : "否"}</span>
      )
    },
    {
      label: "提单路径",
      prop: "submit_address"
    },
    {
      label: "查单路径",
      prop: "query_address"
    },
    {
      label: "余额路径",
      prop: "balance_address"
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

  async function onSearch() {
    loading.value = true;
    var params = { ...form, ...pagination };
    console.log("params", params);
    const { data } = await getSupplierTemplateList(params);
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
  });

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
