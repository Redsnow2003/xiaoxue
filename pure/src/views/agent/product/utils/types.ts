//供应商产品定义
interface SupplierProcuctItem {
  //产品id
  product_id: number;
  //折扣
  discount: number;
  //数量
  count: number;
}

interface FormItemProps {
  //供应商id
  supplier_id: number;
  //产品列表
  product_list: SupplierProcuctItem[];
}

interface FormProps {
  formInline: FormItemProps;
}
//供应商简单信息
interface SupplierSimpleItem {
  /** ID */
  id: number;
  /** 名称 */
  name: string;
}

interface CategoryItemProps {
  /** ID */
  id: number;
  /** 标签 */
  category_name: string;
}

interface CategoryProps extends Array<CategoryItemProps> {}

interface ProductBaseInfo {
  /** ID */
  id: number;
  /** 名称 */
  name: string;
}

interface ProductBaseInfoArray extends Array<ProductBaseInfo> {}

export type {
  FormItemProps,
  FormProps,
  SupplierSimpleItem,
  CategoryProps,
  CategoryItemProps,
  ProductBaseInfo,
  ProductBaseInfoArray
};
