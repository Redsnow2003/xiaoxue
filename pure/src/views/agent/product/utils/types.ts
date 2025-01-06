//供应商产品定义
interface AgentProcuctItem {
  //产品id
  product_id: number;
  //折扣
  discount: number;
}

interface NewProductFormItemProps {
  //供应商id
  agent_id: number;
  //产品列表
  product_list: AgentProcuctItem[];
}

interface NewProdtuctFormProps {
  formInline: NewProductFormItemProps;
}
//供应商简单信息
interface SupplierSimpleItem {
  /** ID */
  id: number;
  /** 名称 */
  name: string;
}

interface AgentProcuctItem {
  /** 代理商产品ID */
  id: number;
  /** 业务类型 */
  business_type: number;
  /** 代理商ID */
  agent_id: number;
  /** 代理商名称 */
  agent_name: string;
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
  /** 供货策略 */
  supply_strategy: number;
  /** 备用通道供货策略 */
  backup_channel_strategy: number;
  /** 折扣类型 */
  discount_type: number;
  /** 折扣 */
  discount: number;
  /** 超时时间 */
  timeout: number;
  /** 超时不缓存 */
  timeout_not_cache: number;
  /** 自动提交备份 */
  auto_submit_backup: number;
  /** 内部时间 */
  interal_time: number;
  /** 支持缓存 */
  support_cache: number;
  /** 转网检测 */
  transfer_check: number;
  /** 空号检测 */
  empty_check: number;
  /** 禁用地区 */
  disabled_area: number[];
  /** 可用地区 */
  scope: number;
  enabled_area: number[];
  /** 限定运营商 */
  limit_operator: number[];
  /** 状态 */
  status: number;
  /** 备注 */
  remark: string;
  /** 定时更改时间 */
  timing_change_time: string;
}

interface AgentProductProps {
  formInline: AgentProcuctItem;
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
  NewProductFormItemProps,
  NewProdtuctFormProps,
  SupplierSimpleItem,
  AgentProcuctItem,
  AgentProductProps,
  CategoryProps,
  CategoryItemProps,
  ProductBaseInfo,
  ProductBaseInfoArray
};
