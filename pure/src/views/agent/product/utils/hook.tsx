import editForm from "../form.vue";
import { getAgentSimpleList, getAgentProductList } from "@/api/agent";
import {
  getProductCategoryList,
  getProductInformationIdAndName
} from "@/api/product";
import { SupplyStrategyList } from "@/api/constdata";
import { addDialog } from "@/components/ReDialog";
import { reactive, ref, onMounted, h } from "vue";
import type {
  FormItemProps,
  SupplierSimpleItem,
  CategoryProps,
  ProductBaseInfoArray
} from "../utils/types";
import { deviceDetection } from "@pureadmin/utils";
import type { PaginationProps } from "@pureadmin/table";
import { useRoute } from "vue-router";

export function useDept() {
  const form = reactive({
    //业务类型
    business_type: "",
    //供货商
    agent_id: "" as any,
    //产品类别
    product_category: "",
    //产品ID
    product_id: "",
    //产品名称
    product_name: "",
    //运营商
    operator: "",
    //状态
    status: "",
    //支持缓存
    support_cache: "",
    //转网检测
    transfer_check: "",
    //空号检测
    empty_check: "",
    //自动提交备用通道
    auto_submit_backup: "",
    //超时不缓存
    timeout_not_cache: ""
  });

  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });

  const agentItemLists = ref([] as SupplierSimpleItem[]);
  const productCategoryList = ref([] as CategoryProps);
  const productBaseInfoList = ref([] as ProductBaseInfoArray);
  const formRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const route = useRoute();
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
      minWidth: 70
    },
    {
      label: "编号",
      prop: "id",
      minWidth: 70
    },
    {
      label: "代理商",
      minWidth: 200,
      prop: "agent_id",
      cellRenderer: ({ row }) => <>{row.agent_id + "\n " + row.agent_name}</>
    },
    {
      label: "产品类别",
      prop: "product_category",
      minWidth: 100
    },
    {
      label: "产品",
      prop: "product_id",
      cellRenderer: ({ row }) => (
        <>
          {row.product_id + "\n " + row.product_name + "\n " + row.base_price}
        </>
      )
    },
    {
      label: "产品运营商",
      prop: "operator",
      minWidth: 100
    },
    {
      label: "产品折扣价",
      prop: "discount",
      cellRenderer: ({ row }) => (
        <>{row.discount + "\n" + row.discount * row.base_price}</>
      )
    },
    {
      label: "供货策略",
      prop: "supply_strategy",
      cellRenderer: ({ row }) => <>{SupplyStrategyList[row.supply_strategy]}</>
    },
    {
      label: "超时时间(秒)",
      prop: "timeout"
    },
    {
      label: "超时不缓存",
      prop: "timeout_not_cache",
      cellRenderer: ({ row }) => (
        <>{row.timeout_not_cache === 0 ? "启动" : "不启动"}</>
      )
    },
    {
      label: "策略",
      prop: "auto_submit_backup",
      cellRenderer: ({ row }) => (
        <>{row.timeout_not_cache === 0 ? "启动" : "不启动"}</>
      )
    },
    {
      label: "支持缓存",
      prop: "support_cache",
      cellRenderer: ({ row }) => (
        <>{row.support_cache === 0 ? "支持" : "不支持"}</>
      )
    },
    {
      label: "转网检测",
      prop: "transfer_check",
      cellRenderer: ({ row }) => (
        <>{row.transfer_check === 0 ? "不启用" : "启用"}</>
      )
    },
    {
      label: "空号检测",
      prop: "empty_check",
      cellRenderer: ({ row }) => (
        <>{row.empty_check === 0 ? "不启用" : "启用"}</>
      )
    },
    {
      label: "禁用地区",
      prop: "disabled_area"
    },
    {
      label: "可用地区",
      prop: "enabled_area"
    },
    {
      label: "限定运营商",
      prop: "limit_operator"
    },
    {
      label: "状态",
      prop: "status",
      cellRenderer: ({ row }) => <>{row.status === 0 ? "启用" : "禁用"}</>
    },
    {
      label: "备注",
      prop: "remark",
      minWidth: 100
    },
    {
      label: "操作",
      fixed: "right",
      width: 210,
      slot: "operation"
    }
  ];

  function handleSelectionChange(val) {
    console.log("handleSelectionChange", val);
  }

  function resetForm(formEl) {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    var params = {
      currentPage: pagination.currentPage,
      pageSize: pagination.pageSize,
      ...form
    };
    const { data } = await getAgentProductList(params);
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;
    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}部门`,
      props: {
        formInline: {
          supplier_id: row?.supplier_id ?? 0
        }
      },
      width: "40%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(valid => {
          if (valid) {
            // 表单规则校验通过
            if (title === "新增") {
              // 实际开发先调用新增接口，再进行下面操作
              console.log("新增", curData);
            } else {
              // 实际开发先调用修改接口，再进行下面操作
            }
            chores();
          }
        });
      }
    });
  }

  onMounted(async () => {
    form.agent_id = Number(route.query.agent_id);
    onSearch();
    const response = await getAgentSimpleList();
    agentItemLists.value = response.data;
    var requestData = { category_name: "", currentPage: 1, pageSize: 100 };
    const response2 = await getProductCategoryList(requestData);
    productCategoryList.value = response2.data.list;
    const response3 = await getProductInformationIdAndName();
    productBaseInfoList.value = response3.data;
  });

  return {
    form,
    loading,
    columns,
    dataList,
    agentItemLists,
    productCategoryList,
    productBaseInfoList,
    onSearch,
    resetForm,
    openDialog,
    handleSelectionChange
  };
}
