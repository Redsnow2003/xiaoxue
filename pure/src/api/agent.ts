import { http } from "@/utils/http";
import { baseUrlApi, type ResultTable, type Result } from "./utils";

/** 代理商管理-获取全部代理商简单信息 */
export const getAgentSimpleList = () => {
  return http.request<Result>("post", baseUrlApi("/get-agent-simple-list"));
};

/** 代理商管理-代理商-获取列表 */
export const getAgentList = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/agent-list"), {
    data
  });
};

/** 代理商管理-代理商-新增 */
export const addAgent = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent"), {
    data
  });
};

/** 代理商管理-代理商-修改 */
export const updateAgent = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/agent"), {
    data
  });
};

/** 代理商管理-代理商-删除 */
export const deleteAgent = (data: object) => {
  return http.request<Result>("delete", baseUrlApi("/agent"), {
    data
  });
};

/** 代理商管理-批量修改代理商状态 */
export const batchUpdateAgentStatus = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-update-agent-status"), {
    data
  });
};

/** 代理商管理-代理商资金操作 */
export const changeAgentFund = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-fund"), {
    data
  });
};

/** 代理商管理-获取代理商资金操作日志 */
export const getAgentFundLog = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/get-agent-fund-log"), {
    data
  });
};

/** 代理商管理-获取代理商IP白名单 */
export const getAgentIpWhiteList = (data: object) => {
  return http.request<ResultTable>(
    "post",
    baseUrlApi("/get-agent-ip-white-list"),
    {
      data
    }
  );
};

/** 代理商管理-新增代理商IP白名单 */
export const addAgentIpWhiteList = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-ip-white-list"), {
    data
  });
};

/** 代理商管理-删除代理商IP白名单 */
export const deleteAgentIpWhiteList = (data: object) => {
  return http.request<Result>("delete", baseUrlApi("/agent-ip-white-list"), {
    data
  });
};

/** 代理商管理-更新代理商IP白名单 */
export const updateAgentIpWhiteList = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/agent-ip-white-list"), {
    data
  });
};

/** 代理商管理-获取代理商产品 */
export const getAgentProductList = (data: object) => {
  return http.request<ResultTable>(
    "post",
    baseUrlApi("/get-agent-product-list"),
    {
      data
    }
  );
};

/** 代理商管理-获取所有代理商产品 */
export const getAllAgentProductList = (data: object) => {
  return http.request<Result>(
    "post",
    baseUrlApi("/get-all-agent-product-list"),
    {
      data
    }
  );
};

/** 代理商管理-添加代理商产品 */
export const addAgentProduct = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-product"), {
    data
  });
};

/** 代理商管理-修改代理商产品 */
export const updateAgentProduct = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/agent-product"), {
    data
  });
};

/** 代理商管理-删除代理商产品 */
export const deleteAgentProduct = (data: object) => {
  return http.request<Result>("delete", baseUrlApi("/agent-product"), {
    data
  });
};

/** 代理商管理-批量修改代理商产品 */
export const batchUpdateAgentProduct = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/batch-update-agent-product"),
    {
      data
    }
  );
};

/** 代理商管理-批量修改代理商产品折扣 */
export const batchUpdateAgentProductDiscount = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/batch-update-agent-product-discount"),
    {
      data
    }
  );
};
