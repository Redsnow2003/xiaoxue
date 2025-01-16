import newForm from "../form/newproduct.vue";
import editForm from "../form/editproduct.vue";
import batchForm from "../form/batchchangdiscount.vue";
import {
  getAgentSimpleList,
  getAgentProductList,
  addAgentProduct,
  updateAgentProduct,
  deleteAgentProduct,
  batchUpdateAgentProduct,
  batchUpdateAgentProductDiscount
} from "@/api/agent";
import {
  getProductCategoryList,
  getProductInformationIdAndName
} from "@/api/product";
import {
  SupplyStrategyList,
  BusinessTypeList,
  OperatorListTelecom,
  ProvinceList,
  OperatorListAll
} from "@/api/constdata";
import { addDialog } from "@/components/ReDialog";
import { reactive, ref, onMounted, h, type Ref } from "vue";
import type {
  NewProductFormItemProps,
  SupplierSimpleItem,
  AgentProcuctItem,
  CategoryProps,
  ProductBaseInfoArray
} from "../utils/types";
import { deviceDetection, getKeyList } from "@pureadmin/utils";
import type { PaginationProps } from "@pureadmin/table";
import { useRoute, useRouter } from "vue-router";

export function useDept(tableRef: Ref) {
  const form = reactive({
    //业务类型
    business_type: "",
    //供货商
    agent_id: "" as any,
    agent_name: "" as any,
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
  const selectedNum = ref(0);
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
      label: "编号",
      prop: "id"
    },
    {
      label: "代理商",
      minWidth: 150,
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
      label: "产品类别",
      prop: "product_category",
      minWidth: 100,
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
      label: "产品",
      prop: "product_id",
      cellRenderer: ({ row }) => (
        <span>
          {row.product_id}
          <br />
          {row.product_name}
          <br />
          {row.base_price}
        </span>
      )
    },
    {
      label: "产品运营商",
      prop: "operator",
      minWidth: 100,
      cellRenderer: ({ row }) => (
        <span>
          {OperatorListTelecom.find(item => item.value === row.operator)?.label}
        </span>
      )
    },
    {
      label: "产品折扣价",
      prop: "discount",
      cellRenderer: ({ row }) => (
        <span>
          {row.discount}
          <br />
          {(row.discount * row.base_price).toFixed(2)}
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
      label: "超时时间(秒)",
      prop: "timeout"
    },
    {
      label: "超时不缓存",
      prop: "timeout_not_cache",
      cellRenderer: ({ row }) => (
        <span>{row.timeout_not_cache === 0 ? "启动" : "不启动"}</span>
      )
    },
    {
      label: "策略",
      prop: "auto_submit_backup",
      cellRenderer: ({ row }) => (
        <span>{row.timeout_not_cache === 0 ? "启动" : "不启动"}</span>
      )
    },
    {
      label: "支持缓存",
      prop: "support_cache",
      cellRenderer: ({ row }) => (
        <span>{row.support_cache === 0 ? "支持" : "不支持"}</span>
      )
    },
    {
      label: "转网检测",
      prop: "transfer_check",
      cellRenderer: ({ row }) => (
        <span>{row.transfer_check === 0 ? "不启用" : "启用"}</span>
      )
    },
    {
      label: "空号检测",
      prop: "empty_check",
      cellRenderer: ({ row }) => (
        <span>{row.empty_check === 0 ? "不启用" : "启用"}</span>
      )
    },
    {
      label: "禁用地区",
      prop: "disabled_area",
      cellRenderer: ({ row, props }) => (
        <span>
          {row.disabled_area
            .split(",")
            .filter(item => item)
            .map((item: number) => (
              <el-tag size={props.size} key={item}>
                {ProvinceList.find(province => province.value == item)?.label}
              </el-tag>
            ))}
        </span>
      )
    },
    {
      label: "可用地区",
      prop: "enabled_area",
      cellRenderer: ({ row, props }) => (
        <span>
          {row.enabled_area
            .split(",")
            .filter(item => item)
            .map((item: number) => (
              <el-tag size={props.size} key={item}>
                {ProvinceList.find(province => province.value == item)?.label}
              </el-tag>
            ))}
        </span>
      )
    },
    {
      label: "限定运营商",
      prop: "limit_operator",
      cellRenderer: ({ row, props }) => (
        <span>
          {row.limit_operator
            .split(",")
            .filter(item => item)
            .map((item: number) => (
              <el-tag size={props.size} key={item}>
                {
                  OperatorListAll.find(operator => operator.value == item)
                    ?.label
                }
              </el-tag>
            ))}
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
      label: "备注",
      prop: "remark",
      minWidth: 100
    },
    {
      label: "操作",
      fixed: "right",
      width: 250,
      slot: "operation"
    }
  ];

  /** 取消选择 */
  function onSelectionCancel() {
    selectedNum.value = 0;
    // 用于多选表格，清空用户的选择
    tableRef.value.getTableRef().clearSelection();
  }
  /** 当CheckBox选择项发生变化时会触发该事件 */
  function handleSelectionChange(val) {
    selectedNum.value = val.length;
    // 重置表格高度
    tableRef.value.setAdaptive();
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
    console.log("data", data);
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;
    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  function batchChangeDiscount() {
    addDialog({
      title: `修改产品`,
      props: {
        formInline: {
          agent_id: form.agent_id
        }
      },
      width: "75%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(batchForm, { formInline: null }),
      beforeSure: async (done, { options }) => {
        const curData = options.props.formInline as AgentProcuctItem;
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        console.log("curData", curData);
        await batchUpdateAgentProductDiscount(curData);
        chores();
      }
    });
  }

  function batchChangeProduct() {
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    if (ids.length === 0) {
      return;
    }
    addDialog({
      title: `修改产品`,
      props: {
        formInline: {}
      },
      width: "75%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as AgentProcuctItem;
        //如果curData的字段只有一个，说明没有填写任何信息，直接关闭弹框
        if (Object.keys(curData).length === 1) {
          done();
          return;
        }
        // 将curData中的字段product_name删除
        delete curData.product_name;
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async (valid: boolean) => {
          if (valid) {
            var data = { ...curData, ids };
            console.log("data", data);
            await batchUpdateAgentProduct(data);
            chores();
          }
        });
      }
    });
  }

  function changeProduct(row: AgentProcuctItem) {
    var disabled_area = row.disabled_area as unknown as string;
    var enabled_area = row.enabled_area as unknown as string;
    var limit_operator = row.limit_operator as unknown as string;
    addDialog({
      title: `修改产品`,
      props: {
        formInline: {
          ...row,
          disabled_area: disabled_area.split(",").map(Number),
          enabled_area: enabled_area.split(",").map(Number),
          limit_operator: limit_operator.split(",").map(Number)
        }
      },
      width: "75%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as AgentProcuctItem;
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            await updateAgentProduct(curData);
            chores();
          }
        });
      }
    });
  }

  function newProcuct() {
    addDialog({
      title: `新增产品`,
      props: {
        formInline: {
          agent_id: form.agent_id,
          agent_name: form.agent_name,
          product_list: []
        }
      },
      width: "60%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(newForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as NewProductFormItemProps;
        curData.product_list.forEach(item => {
          item.enabled_area = String(item.scope).split(",").map(Number);
          item.limit_operator = String(item.limit_operator)
            .split(",")
            .map(Number);
        });
        console.log("curData", curData);
        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            await addAgentProduct(curData);
            chores();
          }
        });
      }
    });
  }

  function handleProductChannel(row: AgentProcuctItem) {
    console.log("row", row);
  }

  async function handleBatchDelete() {
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    var ids = getKeyList(curSelected, "id");
    await deleteAgentProduct(ids).then(() => {
      onSearch();
      tableRef.value.getTableRef().clearSelection();
    });
  }

  async function handleDelete(row) {
    var ids = [row.id];
    await deleteAgentProduct(ids).then(() => {
      onSearch();
    });
  }

  onMounted(async () => {
    if (route.query.agent_id) {
      form.agent_id = Number(route.query.agent_id);
      form.agent_name = route.query.agent_name as string;
    }
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
    selectedNum,
    pagination,
    agentItemLists,
    productCategoryList,
    productBaseInfoList,
    onSearch,
    resetForm,
    newProcuct,
    changeProduct,
    batchChangeProduct,
    batchChangeDiscount,
    handleDelete,
    handleBatchDelete,
    handleSelectionChange,
    handleProductChannel,
    onSelectionCancel
  };
}

function useProductHandlers() {
  const router = useRouter();
  function handleProductChannel(row?: AgentProcuctItem) {
    console.log("handleWhitelist", row);
    if (row && row.id) {
      // 携带参数 row.id 跳转到白名单配置页面
      router.push({
        path: "/agent/channel/index",
        query: {
          agent_id: row.agent_id,
          agent_name: row.agent_name
        }
      });
    } else {
      console.error("Row or Row ID is missing");
    }
  }

  return { handleProductChannel };
}
export { useProductHandlers };
