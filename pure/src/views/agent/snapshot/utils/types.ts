// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了

interface FormItemProps {
  /** 代理商产品ID */
  id: number;
  /** 业务类型 */
  business_type: number;
  /** 产品ID */
  product_id: number;
  /** 产品名称 */
  product_name: string;
  /** 产品类别 */
  product_category: number;
  /** 运营商 */
  operator: number;
  /** 基础价格 */
  base_price: number;
  /** 代理商ID */
  agent_id: number;
  /** 代理商名称 */
  agent_name: string;
  /** 代理商产品ID */
  agent_product_id: number;
  /** 代理商折扣 */
  agent_discount: number;
  /** 供货策略 */
  supply_strategy: number;
  /** 供应商产品ID */
  supplier_product_id: number;
  /** 供应商ID */
  supplier_id: number;
  /** 供应商名称 */
  supplier_name: string;
  /** 供应商折扣 */
  supplier_discount: number;
  /** 上游产品ID */
  up_product_id: number;
  /** 执行次数 */
  execute_count: number;
  /** 权重 */
  weight: number;
  /** 优先级 */
  priority: number;
  /** 省份 */
  province: number;
  /** 状态 */
  status: number;
  /** 是否连接 */
  is_connect: number;
  /** 备注 */
  remark: string;
}
interface FormProps {
  formInline: FormItemProps;
}

interface CategoryItemProps {
  /** ID */
  id: number;
  /** 标签 */
  category_name: string;
}

interface CategoryProps extends Array<CategoryItemProps> {}

export type { FormItemProps, FormProps, CategoryProps };
