import editForm from "../form.vue";
import { getSupplierProductList, getSupplierSimpleList } from "@/api/supplier";
import {
  getProductCategoryList,
  getProductInformationIdAndName
} from "@/api/product";
import { addDialog } from "@/components/ReDialog";
import { reactive, ref, onMounted, h } from "vue";
import type {
  FormItemProps,
  CategoryProps,
  ProductBaseInfoArray
} from "../utils/types";
import type { SupplierIdName } from "@/api/types";
import { deviceDetection } from "@pureadmin/utils";
import type { PaginationProps } from "@pureadmin/table";
import { useRoute } from "vue-router";

export function useDept() {
  const form = reactive({
    business_type: "",
    supplier_id: "" as any,
    product_category: "",
    product_id: "",
    product_name: "",
    operator: "",
    up_product_id: "",
    status: ""
  });

  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });

  const supplierItemLists = ref([] as SupplierIdName[]);
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
      label: "供货商",
      minWidth: 200,
      prop: "supplier_id",
      cellRenderer: ({ row }) => (
        <span>{row.supplier_id + "\n " + row.supplier_name}</span>
      )
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
        <span>{row.product_id + "\n " + row.product_name}</span>
      )
    },
    {
      label: "基础价",
      prop: "base_price",
      minWidth: 100
    },
    {
      label: "折扣类型",
      prop: "discount_type",
      cellRenderer: ({ row }) => (
        <span>{row.discount_type === 0 ? "按折扣" : "按金额"}</span>
      )
    },
    {
      label: "产品折扣",
      prop: "discount",
      minWidth: 100
    },
    {
      label: "供货价",
      prop: "discount",
      cellRenderer: ({ row }) => (
        <span>
          {row.discount_type === 0
            ? row.discount * row.base_price
            : row.discount}
        </span>
      )
    },
    {
      label: "上游产品ID",
      prop: "up_product_id",
      minWidth: 100
    },
    {
      label: "分省",
      prop: "province",
      minWidth: 100
    },
    {
      label: "同配置序号",
      prop: "serial_number",
      minWidth: 100
    },
    {
      label: "状态",
      prop: "status",
      cellRenderer: ({ row }) => (
        <span>{row.status === 0 ? "维护" : "上架"}</span>
      )
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
    const { data } = await getSupplierProductList(params);
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
    //将form.supplier_id类型改为number
    form.supplier_id = Number(route.query.supplier_id);
    onSearch();
    const response = await getSupplierSimpleList();
    supplierItemLists.value = response.data;
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
    supplierItemLists,
    productCategoryList,
    productBaseInfoList,
    onSearch,
    resetForm,
    openDialog,
    handleSelectionChange
  };
}
