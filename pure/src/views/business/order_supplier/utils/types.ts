// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了
interface OrederSupplierItemProps {
  /** 订单ID */
  id: number;
  /** 订单号 */
  order_id: number;
  /** 业务类型 */
  business_type: number;
  /** 上游ID */
  up_id: string;
  /** 代理商ID */
  agent_id: number;
  /** 代理商名称 */
  agent_name: string;
  /** 代理商折扣 */
  agent_discount: number;
  /** 数量 */
  count: number;
  /** 供应商ID */
  supplier_id: number;
  /** 供应商名称 */
  supplier_name: string;
  /** 产品ID */
  product_id: number;
  /** 产品名称 */
  product_name: string;
  /** 基础价格 */
  base_price: number;
  /** 运营商 */
  operator: number;
  /** 供应商折扣 */
  supplier_discount: number;
  /** 充值号码 */
  recharge_number: string;
  /** 归属地 */
  location: string;
  /** 订单时间 */
  order_time: string;
  /** 订单状态 */
  status: string;
  /** 创建时间 */
  create_time: string;
  /** 完成时间 */
  finish_time: string;
  /** 上游信息 */
  up_information: string;
  /** 更新时间 */
  update_time: string;
  /** 备注 */
  remark: string;
  /** 是否备份 */
  is_backup: boolean;
}
interface OrederSupplierFormProps {
  formInline: OrederSupplierItemProps;
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
  OrederSupplierItemProps,
  OrederSupplierFormProps,
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
