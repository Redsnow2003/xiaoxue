// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了

interface FormItemProps {
  /** ID */
  id: number;
  /** 名称 */
  name: string;
  /** 全局回调地址 */
  global_cb_address: string;
  /** 是否需要商品编码 */
  is_need_product_id: number;
  /** 是否需要绑定回调地址 */
  is_bind_callback_address: number;
  /** 是否支持处理供货金额不一致 */
  is_support_inconsistent: number;
  /** 是否支持撤单 */
  is_support_cancel: number;
  /** 提单地址 */
  submit_address: string;
  /** 查单地址 */
  query_address: string;
  /** 查询余额地址 */
  balance_address: string;
  /** 模板表单json数据 */
  template_json: string;
  /** 备注 */
  remark: string;
}
interface FormProps {
  formInline: FormItemProps;
}

export type { FormItemProps, FormProps };
