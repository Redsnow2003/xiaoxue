import { http } from "@/utils/http";
import { baseUrlApi, type ResultTable, type Result } from "./utils";

/** 供应商管理-模板-获取模板列表 */
export const getSupplierTemplateList = (data?: object) => {
  return http.request<ResultTable>(
    "post",
    baseUrlApi("/get-supplier-template-list"),
    {
      data
    }
  );
};

/** 供应商管理-模板-获取模板JSON */
export const getSupplierTemplateJson = (data?: object) => {
  return http.request<Result>(
    "post",
    baseUrlApi("/get-supplier-template-json"),
    {
      data
    }
  );
};

/** 供应商管理-获取所有模板名称 */
export const getSupplierTemplateNameList = () => {
  return http.request<Result>(
    "post",
    baseUrlApi("/get-supplier-template-name-list")
  );
};

/** 供应商管理-获取全部供应商简单信息 */
export const getSupplierSimpleList = () => {
  return http.request<Result>("post", baseUrlApi("/get-supplier-simple-list"));
};

/** 供应商管理-供应商-获取列表 */
export const getSupplierList = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/supplier-list"), {
    data
  });
};

/** 供应商管理-供应商-新增 */
export const addSupplier = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/supplier"), {
    data
  });
};

/** 供应商管理-供应商-修改 */
export const updateSupplier = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/supplier"), {
    data
  });
};

/** 供应商管理-供应商-删除 */
export const deleteSupplier = (data: object) => {
  return http.request<Result>("delete", baseUrlApi("/supplier"), {
    data
  });
};

/** 供应商管理-批量修改供应商状态 */
export const batchUpdateSupplierStatus = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/batch-update-supplier-status"),
    {
      data
    }
  );
};

/** 供应商管理-批量更新供应商余额 */
export const batchUpdateSupplierBalance = () => {
  return http.request<Result>(
    "put",
    baseUrlApi("/batch-update-supplier-balance")
  );
};

/** 供应商管理-资金操作 */
export const supplierChangeFund = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/supplier-change-fund"), {
    data
  });
};
