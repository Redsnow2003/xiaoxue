// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了
interface OrderItemProps {
  id: number; // 订单ID
  business_type: string; // 业务类型
  down_id: string; // 下游单号
  agent_id: number; // 代理商ID
  agent_name: string; // 代理商名称
  product_category: string; // 产品分类
  product_id: number; // 产品ID
  product_name: string; // 产品名称
  base_price: number; // 基础价格
  operator: string; // 运营商
  agent_discount: number; // 代理商折扣
  count: number; // 购买数量
  recharge_number: string; // 充值号码
  location: string; // 归属地
  status: number; // 订单状态
  is_timeout: boolean; // 是否超时
  timeout: number; // 超时时长
  is_cancel: boolean; // 是否已取消
  create_time: string; // 创建时间
  finish_time: string; // 完成时间
  notify_status: number; // 通知状态
  special_params: Record<string, any>; // 特殊参数
  remark: string; // 备注
}
interface OrderProps {
  formInline: OrderItemProps;
}

interface BatchStatusProps {
  /** 状态 */
  status: number;
  /** 备注 */
  remark: string;
}

interface BatchStatusFormProps {
  formInline: BatchStatusProps;
}

interface BatchBackupProps {
  /** 次数 */
  count: number;
  /** 时间间隔 */
  interval: number;
}

interface BatchBackupFormProps {
  formInline: BatchBackupProps;
}

interface BatchTimeoutProps {
  type: number;
  seconds: number;
  time: string;
  create_after: number;
}

interface BatchTimeoutFormProps {
  formInline: BatchTimeoutProps;
}

interface ChangeRemarkProps {
  /** 订单ID */
  order_id: number;
  /** 备注 */
  remark: string;
}

interface ChangeRemarkFormProps {
  formInline: ChangeRemarkProps;
}

interface ChangeFundFormItemProps {
  agent_id: string; // 代理商ID
  agent_name: string; // 代理商名称
  fund_action: number; // 余额操作
  amount: number; // 金额
  confirm_amount: number; // 确认金额
  file: string; // 附件
  remark: string; // 备注
}

interface ChangeFundFormProps {
  formInline: ChangeFundFormItemProps;
}

interface TemplateInfoItem {
  index: number;
  label: string;
  prop: string;
  type: string;
  value: string;
  options: any[];
}

interface CategoryItemProps {
  /** ID */
  id: number;
  /** 标签 */
  category_name: string;
}

interface CategoryProps extends Array<CategoryItemProps> {}

export type {
  OrderItemProps,
  OrderProps,
  ChangeFundFormProps,
  TemplateInfoItem,
  ChangeFundFormItemProps,
  CategoryProps,
  BatchStatusProps,
  BatchStatusFormProps,
  BatchBackupProps,
  BatchBackupFormProps,
  BatchTimeoutProps,
  BatchTimeoutFormProps,
  ChangeRemarkProps,
  ChangeRemarkFormProps
};
