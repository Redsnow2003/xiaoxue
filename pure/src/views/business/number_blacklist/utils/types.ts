// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了
interface BacklistItemProps {
  id: number;
  recharge_number: string;
  remark: string;
}
interface Backlistrops {
  formInline: BacklistItemProps;
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
  BacklistItemProps,
  Backlistrops,
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
