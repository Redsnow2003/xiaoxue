// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了
interface FormItemProps {
  /** 代理商ID */
  id: number;
  /** 代理商名称 */
  name: string;
  /** 代理商简称 */
  nickname: string;
  /** 代理商部门 */
  dept: number;
  /** 代理商电话 */
  phone: string;
  /** 代理商邮箱 */
  email: string;
  /** 密钥 */
  secret_key: string;
  /** 通知地址 */
  notification_address: string;
  /** 通知方式 */
  notification_method: number;
  /** 客户 */
  customer: string;
  /** 状态 */
  status: number;
  /** 资金余额 */
  fund_balance: number;
  /** 授信余额 */
  credit_balance: number;
  /** 冻结金额 */
  frozen_amount: number;
  /** 缓存可用金额 */
  cache_amount: number;
  /** 备注 */
  remark: string;
}
interface FormProps {
  formInline: FormItemProps;
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
  FormItemProps,
  FormProps,
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
