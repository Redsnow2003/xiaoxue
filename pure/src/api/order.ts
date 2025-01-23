import { http } from "@/utils/http";
import { baseUrlApi, type ResultTable, type Result } from "./utils";

/** 订单管理-获取订单列表 */
export const getOrderList = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/order-list"), {
    data
  });
};

/** 批量修改订单状态备注 */
export const batchUpdateOrderStatusRemark = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/batch-update-order-status-remark"),
    {
      data
    }
  );
};

/** 批量通知 */
export const batchOrderNotice = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/batch-order_notice"), {
    data
  });
};

/** 批量备用通道重新提交 */
export const batchBackupSubmit = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-backup-submit"), {
    data
  });
};

/** 批量备用通道取消 */
export const batchBackupCancel = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-backup-cancel"), {
    data
  });
};

/** 批量订单超时 */
export const batchOrderTimeout = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-order-timeout"), {
    data
  });
};

/** 批量订单取消 */
export const batchOrderCancel = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-order-cancel"), {
    data
  });
};

/** 批量订单转人工处理 */
export const batchOrderManual = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-order-manual"), {
    data
  });
};

/** 获取供货单列表 */
export const getSupplierOrderList = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/supplier-order-list"), {
    data
  });
};

/** 修改订单备注 */
export const updateOrderRemark = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/update-order-remark"), {
    data
  });
};

/** 代理商订单提交信息 */
export const agentOrderSubmit = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-order-submit"), {
    data
  });
};

/** 代理商订单查单信息 */
export const agentOrderQuery = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-order-query"), {
    data
  });
};

/** 代理商订单通知信息 */
export const agentOrderNotice = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-order-notice"), {
    data
  });
};

/** 供货商订单批量自动核单 */
export const batchSupplierAutoCheck = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-supplier-auto-check"), {
    data
  });
};

/** 供货商订单批量手动核单 */
export const batchSupplierManualCheck = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/batch-supplier-manual-check"),
    {
      data
    }
  );
};

/** 供货商订单批量备用通道提交 */
export const batchSupplierBackupSubmit = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/batch-supplier-backup-submit"),
    {
      data
    }
  );
};

/** 供货商订单批量取消 */
export const batchSupplierCancel = (data: object) => {
  return http.request<Result>("put", baseUrlApi("/batch-supplier-cancel"), {
    data
  });
};

/** 供货商订单失败转成功 */
export const supplierOrderFailToSuccess = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/supplier-order-fail-to-success"),
    {
      data
    }
  );
};

/** 供货商订单修改备注信息 */
export const updateSupplierOrderRemark = (data: object) => {
  return http.request<Result>(
    "put",
    baseUrlApi("/update-supplier-order-remark"),
    {
      data
    }
  );
};
