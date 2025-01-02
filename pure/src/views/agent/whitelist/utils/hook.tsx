import type { PaginationProps } from "@pureadmin/table";
import editForm from "../form.vue";
import { addDialog } from "@/components/ReDialog";
import {
  getAgentIpWhiteList,
  getAgentSimpleList,
  addAgentIpWhiteList,
  updateAgentIpWhiteList,
  deleteAgentIpWhiteList
} from "@/api/agent";
import { reactive, ref, onMounted, h } from "vue";
import type { AgentSimpleItem, FormItemProps } from "../utils/types";
import { deviceDetection } from "@pureadmin/utils";
import { useRoute } from "vue-router";
import { message } from "@/utils/message";
export function useCategory() {
  const form = reactive({
    /** 代理商 */
    agent_id: "" as any,
    agent_name: "",
    /** ip */
    ip: ""
  });
  const curRow = ref();
  const dataList = ref([]);
  const formRef = ref();
  const loading = ref(true);
  const agentItemLists = ref([] as AgentSimpleItem[]);
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
      label: "Ip地址",
      prop: "ip"
    },
    {
      label: "Ip归属地",
      prop: "ip_location"
    },
    {
      label: "备注",
      prop: "remark"
    },
    {
      label: "操作",
      fixed: "right",
      width: 210,
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

  async function onSearch() {
    loading.value = true;
    var params = { ...form, ...pagination };
    console.log("params", params);
    const { data } = await getAgentIpWhiteList(params);
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

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}IP白名单`,
      props: {
        formInline: {
          id: row?.id ?? 0,
          agent_id: row?.agent_id ?? form.agent_id,
          agent_name: row?.agent_name ?? form.agent_name,
          ip: row?.ip ?? "",
          ip_location: row?.ip_location ?? "",
          remark: row?.remark ?? "",
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
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            // 表单规则校验通过
            if (title === "新增") {
              await addAgentIpWhiteList(curData).then(res => {
                if (res.success) {
                  message("新增白名单成功", { type: "success" });
                  chores();
                } else {
                  message(res.message, { type: "error" });
                }
              });
            } else {
              await updateAgentIpWhiteList(curData).then(res => {
                if (res.success) {
                  message("修改白名单成功", { type: "success" });
                } else {
                  message(res.message, { type: "error" });
                }
              });
              chores();
            }
            chores();
          }
        });
      }
    });
  }

  async function handleDelete(row?: FormItemProps) {
    var ids = [row.id];
    await deleteAgentIpWhiteList(ids).then(res => {
      if (res.success) {
        message("删除供应商成功", { type: "success" });
      } else {
        message(res.message, { type: "error" });
      }
    });
    onSearch();
  }

  onMounted(async () => {
    // 动态获取产品类别列表
    const response = await getAgentSimpleList();
    agentItemLists.value = response.data;
    disabled.value = false;
    if (route.query.agent_id !== undefined) {
      form.agent_id = Number(route.query.agent_id);
      form.agent_name = String(route.query.agent_name);
      disabled.value = true;
    }
    onSearch();
  });

  return {
    form,
    disabled,
    curRow,
    loading,
    columns,
    rowStyle,
    dataList,
    pagination,
    agentItemLists,
    onSearch,
    openDialog,
    resetForm,
    handleSizeChange,
    handleDelete,
    handleCurrentChange
  };
}
