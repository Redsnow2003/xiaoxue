// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了

interface FormItemProps {
  /** ID */
  id: number;
  /** 代理商id */
  agent_id: number;
  /** 代理商名称 */
  agent_name: string;
  /** 供货商id */
  supplier_id: number;
  /** 供货商名称 */
  supplier_name: string;
  /** 是否禁用 */
  disabled: boolean;
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

export type { FormItemProps, FormProps, AgentSimpleItem };
