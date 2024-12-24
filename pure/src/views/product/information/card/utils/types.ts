// 虽然字段很少 但是抽离出来 后续有扩展字段需求就很方便了

interface FormItemProps {
  /** 类别ID */
  id: number;
  /** 业务类型 */
  type: number;
  /** 产品类别 */
  category: number;
  /** 产品名称 */
  name: string;
  /** 产品运营商 */
  operator: number;
  /** 面额 */
  price: number;
  /** 单位 */
  unit: number;
  /** 基础价格 */
  base_price: number;
  /** 待售库存 */
  sale_inventory: number;
  /** 待售库存总金额 */
  sale_inventory_amount: number;
  /** API提单数量限制 */
  api_limit: number;
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

export type { FormItemProps, FormProps, CategoryProps, CategoryItemProps };
