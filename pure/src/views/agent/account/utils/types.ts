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

interface AgentSimpleItem {
  /** ID */
  id: number;
  /** 名称 */
  name: string;
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

export type {
  FormItemProps,
  FormProps,
  AgentSimpleItem,
  ChangeFundFormProps,
  TemplateInfoItem,
  ChangeFundFormItemProps
};
