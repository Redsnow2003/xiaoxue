import { http } from "@/utils/http";
import { baseUrlApi, type ResultTable, type Result } from "./utils";

/** 代理商管理-获取全部代理商简单信息 */
export const getAgentSimpleList = () => {
  return http.request<Result>("post", baseUrlApi("/get-Agent-simple-list"));
};

/** 代理商管理-代理商-获取列表 */
export const getAgentList = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/Agent-list"), {
    data
  });
};

/** 代理商管理-代理商-新增 */
export const addAgent = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/Agent"), {
    data
  });
};

/** 代理商管理-代理商-修改 */
export const updateAgent = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/Agent"), {
    data
  });
};

/** 代理商管理-代理商-删除 */
export const deleteAgent = (data: object) => {
  return http.request<Result>("delete", baseUrlApi("/Agent"), {
    data
  });
};

/** 代理商管理-批量修改代理商状态 */
export const batchUpdateAgentStatus = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-update-Agent-status"), {
    data
  });
};

/** 代理商管理-代理商资金操作 */
export const changeAgentFund = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/Agent-fund"), {
    data
  });
};
