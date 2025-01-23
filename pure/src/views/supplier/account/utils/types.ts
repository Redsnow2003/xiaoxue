// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了

interface FormItemProps {
  /** 供应商ID */
  id: number;
  /** 供应商名称 */
  name: string;
  /** 供应商简称 */
  nickname: string;
  /** 供应商部门 */
  dept: number;
  /** 供应商电话 */
  phone: string;
  /** 供应商邮箱 */
  email: string;
  /** 我方平台账户余额(元) */
  our_balance: number;
  /** 供应商平台账户余额(元) */
  up_balance: number;
  /** 上游平台账户余额更新时间 */
  up_balance_update_time: string;
  /** 供应商模板 */
  up_template: string;
  /** 供应商模板JSON */
  template_json: string;
  /** 状态 */
  status: number;
  /** 状态信息 */
  status_info: string;
  /** 备注 */
  remark: string;
}
interface FormProps {
  formInline: FormItemProps;
}

interface TemplateNameItem {
  /** ID */
  id: number;
  /** 名称 */
  name: string;
}

interface ChangeFundFormItemProps {
  supplier_id: string; // 供应商ID
  supplier_name: string; // 供应商名称
  fund_action: string; // 余额操作
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
  TemplateNameItem,
  ChangeFundFormProps,
  TemplateInfoItem,
  ChangeFundFormItemProps
};
