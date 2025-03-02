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

/** 代理商订单提交日志 */
export const agentOrderSubmitLog = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-order-submit-log"), {
    data
  });
};

/** 代理商订单查单日志 */
export const agentOrderQueryLog = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-order-query-log"), {
    data
  });
};

/** 代理商订单通知日志 */
export const agentOrderNoticeLog = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/agent-order-notice-log"), {
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

/** 订单查找电话号码 */
export const orderFindPhone = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/order-find-phone"), {
    data
  });
};

/** 备用通道提交日志 */
export const backupSubmitLog = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/backup-submit-log"), {
    data
  });
};

/** 删除备用通道提交日志 */
export const deleteBackupSubmitLog = (data: object) => {
  return http.request<Result>(
    "delete",
    baseUrlApi("/delete-backup-submit-log"),
    {
      data
    }
  );
};

/** 查询上游订单信息 */
export const queryUpOrderInfo = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/query-up-order-info"), {
    data
  });
};

/** 供货单提交日志 */
export const supplierOrderSubmitLog = (data: object) => {
  return http.request<Result>(
    "post",
    baseUrlApi("/supplier-order-submit-log"),
    {
      data
    }
  );
};

/** 供货单查询日志 */
export const supplierOrderQueryLog = (data: object) => {
  return http.request<Result>("post", baseUrlApi("/supplier-order-query-log"), {
    data
  });
};

/** 供货单撤单日志 */
export const supplierOrderCancelLog = (data: object) => {
  return http.request<Result>(
    "post",
    baseUrlApi("/supplier-order-cancel-log"),
    {
      data
    }
  );
};

/** 供货单回调日志 */
export const supplierOrderCallbackLog = (data: object) => {
  return http.request<Result>(
    "post",
    baseUrlApi("/supplier-order-callback-log"),
    {
      data
    }
  );
};

/** 获取拦截单信息 */
export const getInterceptOrderInfo = (data: object) => {
  return http.request<ResultTable>(
    "post",
    baseUrlApi("/intercept-order-info"),
    {
      data
    }
  );
};

/** 获取缓存单信息列表 */
export const getCacheOrderList = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/cache-order-list"), {
    data
  });
};

/** 号码黑名单查询 */
export const getNumberBlackList = (data: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/number-black-list"), {
    data
  });
};

/** 删除号码黑名单 */
export const deleteNumberBlackList = (data: object) => {
  return http.request<Result>(
    "delete",
    baseUrlApi("/delete-number-black-list"),
    {
      data
    }
  );
};
