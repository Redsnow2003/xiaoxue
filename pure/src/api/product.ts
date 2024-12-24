import { http } from "@/utils/http";
import { baseUrlApi, type ResultTable, type Result } from "./utils";

/** 获取产品管理-产品类别表 */
export const getProductCategoryList = (data?: object) => {
  return http.request<ResultTable>(
    "post",
    baseUrlApi("/get-product-category"),
    {
      data
    }
  );
};

/** 添加产品管理-产品类别 */
export const addProductCategory = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/product-category"), {
    data
  });
};

/** 修改产品管理-产品类别 */
export const updateProductCategory = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/product-category"), { data });
};

/** 删除产品管理-产品类别 */
export const deleteProductCategory = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/product-category"), {
    data
  });
};

/** 获取产品管理-产品信息列表 */
export const getProductInformationList = (data?: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/get-product-list"), {
    data
  });
};

/** 添加产品管理-产品信息 */
export const addProductInformation = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/product-info"), { data });
};

/** 修改产品管理-产品信息 */
export const updateProductInformation = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/product-info"), { data });
};

/** 删除产品管理-产品信息 */
export const deleteProductInformation = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/product-info"), { data });
};

/** 导入产品信息 */
export const importProductInformation = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/product-info/import"), {
    data
  });
};

/** 导出产品信息 */
export const exportProductInformation = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/product-info/export"), {
    data
  });
};
