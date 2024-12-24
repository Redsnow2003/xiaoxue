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
  /** 使用范围 */
  scope: number;
  /** 面额 */
  price: number;
  /** 单位 */
  unit: number;
  /** 基础价格 */
  base_price: number;
  /** 禁用地区 */
  disable_area: number[];
  /** 限定运营商 */
  limit_operator: number;
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
