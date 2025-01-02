// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了

interface FormItemProps {
  /** ID */
  id: number;
  /** 代理商id */
  agent_id: number;
  /** 代理商名称 */
  agent_name: string;
  /** Ip地址 */
  ip: string;
  /** Ip归属地 */
  ip_location: string;
  /** 备注 */
  remark: string;
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
