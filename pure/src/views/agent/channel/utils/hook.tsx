import type { PaginationProps } from "@pureadmin/table";
import editForm from "../form.vue";
import { addDialog } from "@/components/ReDialog";
import {
  getAgentChannelList,
  getAgentSimpleList,
  addAgentChannel,
  deleteAgentChannel
} from "@/api/agent";
import { getSupplierSimpleList } from "@/api/supplier";
import { reactive, ref, onMounted, h, type Ref } from "vue";
import type { AgentSimpleItem, FormItemProps } from "../utils/types";
import type { SupplierIdName } from "@/api/types";
import { deviceDetection } from "@pureadmin/utils";
import { useRoute } from "vue-router";
import { message } from "@/utils/message";
export function useCategory(tableRef: Ref) {
  const form = reactive({
    /** 代理商 */
    agent_id: "" as any,
    agent_name: "",
    /** 供货商 */
    supplier_id: "",
    supplier_name: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const formRef = ref();
  const loading = ref(true);
  const selectedNum = ref(0);
  const agentItemLists = ref([] as AgentSimpleItem[]);
  const supplierItemLists = ref([] as SupplierIdName[]);
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
      label: "ID",
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
      label: "供货商ID",
      prop: "supplier_id"
    },
    {
      label: "供货商名称",
      prop: "supplier_name"
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
    const { data } = await getAgentChannelList(params);
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

  function openDialog() {
    addDialog({
      title: `新增通道`,
      props: {
        formInline: {
          agent_id: form.agent_id,
          agent_name: form.agent_name,
          disabled: disabled.value
        }
      },
      width: "40%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(editForm, {
          ref: formRef,
          formInline: null
        }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;
        function chores() {
          done(); // 关闭弹框
          console.log("chores");
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            // 表单规则校验通过
            await addAgentChannel(curData).then(res => {
              if (res.success) {
                message("新增白名单成功", { type: "success" });
                chores();
              } else {
                message(res.message, { type: "error" });
              }
            });
            chores();
          }
        });
      }
    });
  }

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
    console.log("onMounted");
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
