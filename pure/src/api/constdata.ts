//运营商列表4-全部
const OperatorListAll = [
  { label: "全网", value: 0 },
  { label: "移动", value: 1 },
  { label: "联通", value: 2 },
  { label: "电信", value: 3 },
  { label: "中石油", value: 4 },
  { label: "权益任意数量", value: 5 },
  { label: "中石化", value: 6 },
  { label: "广东石化", value: 7 },
  { label: "南网", value: 8 },
  { label: "国网", value: 9 },
  { label: "广电", value: 10 },
  { label: "权益通用运营商", value: 11 },
  { label: "京东礼品卡(直充)", value: 12 },
  { label: "Q币充值", value: 13 },
  { label: "前向流量", value: 14 },
  { label: "卡券通用运营商", value: 15 }
];
// 运营商列表1-话费流量
const OperatorListTelecom = [
  { label: "全网", value: 0 },
  { label: "移动", value: 1 },
  { label: "电信", value: 3 },
  { label: "联通", value: 2 },
  { label: "广电", value: 10 }
];
//运营商列表2-话费流量-限定运营商
const OperatorListTelecom2 = [
  { label: "移动", value: 1 },
  { label: "电信", value: 3 },
  { label: "联通", value: 2 },
  { label: "广电", value: 10 }
];
//运营商列表3-权益
const OperatorListRights = [
  { label: "权益任意数量", value: 5 },
  { label: "权益通用运营商", value: 11 },
  { label: "京东礼品卡(直充)", value: 12 },
  { label: "Q币充值", value: 13 },
  { label: "前向流量", value: 14 }
];
//运营商列表5-油卡
const OperatorListOilcard = [
  { label: "中石油", value: 4 },
  { label: "中石化", value: 6 },
  { label: "广东石化", value: 7 }
];
//运营商列表6-电费
const OperatorListPowerbill = [
  { label: "南网", value: 8 },
  { label: "国网", value: 9 }
];

//业务类型列表
const BusinessTypeList = [
  { label: "话费", value: 0 },
  { label: "权益", value: 1 },
  { label: "卡券", value: 2 },
  { label: "油卡", value: 3 },
  { label: "流量", value: 4 },
  { label: "电费", value: 5 }
];

//单位列表
const UnitList = [
  { label: "元", value: 0 },
  { label: "M", value: 1 },
  { label: "月", value: 2 },
  { label: "个", value: 3 },
  { label: "周", value: 4 },
  { label: "张", value: 5 },
  { label: "半年", value: 6 },
  { label: "季", value: 7 },
  { label: "年", value: 8 }
];

//全国省份列表
const ProvinceList = [
  { label: "全国", value: 0 },
  { label: "北京", value: 1 },
  { label: "天津", value: 2 },
  { label: "河北", value: 3 },
  { label: "山西", value: 4 },
  { label: "内蒙古", value: 5 },
  { label: "辽宁", value: 6 },
  { label: "吉林", value: 7 },
  { label: "黑龙江", value: 8 },
  { label: "上海", value: 9 },
  { label: "江苏", value: 10 },
  { label: "浙江", value: 11 },
  { label: "安徽", value: 12 },
  { label: "福建", value: 13 },
  { label: "江西", value: 14 },
  { label: "山东", value: 15 },
  { label: "河南", value: 16 },
  { label: "湖北", value: 17 },
  { label: "湖南", value: 18 },
  { label: "广东", value: 19 },
  { label: "广西", value: 20 },
  { label: "海南", value: 21 },
  { label: "重庆", value: 22 },
  { label: "四川", value: 23 },
  { label: "贵州", value: 24 },
  { label: "云南", value: 25 },
  { label: "西藏", value: 26 },
  { label: "陕西", value: 27 },
  { label: "甘肃", value: 28 },
  { label: "青海", value: 29 },
  { label: "宁夏", value: 30 },
  { label: "新疆", value: 31 },
  { label: "台湾", value: 32 },
  { label: "香港", value: 33 },
  { label: "澳门", value: 34 }
];

//供货策略
const SupplyStrategyList = [
  { label: "权重", value: 0 },
  { label: "优先级", value: 1 },
  { label: "全国转分省", value: 2 },
  { label: "成功率", value: 3 }
];

//资金操作类型
const FundOperationTypeList = [
  { label: "余额加款", value: 0 },
  { label: "余额减款", value: 1 },
  { label: "余额校正", value: 2 },
  { label: "授信加款", value: 3 },
  { label: "授信减款", value: 4 },
  { label: "余额支出", value: 5 },
  { label: "余额退款", value: 6 }
];

const FundOperationTypeList2 = [
  { label: "余额加款", value: 0 },
  { label: "余额减款", value: 1 },
  { label: "余额校正", value: 2 },
  { label: "授信加款", value: 3 },
  { label: "授信减款", value: 4 }
];

const FundOperationTypeList3 = [
  { label: "余额加款", value: 0 },
  { label: "余额减款", value: 1 },
  { label: "余额校正", value: 2 }
];

// 通知状态列表
const NotifyStatusList = [
  { label: "未发送", value: 0 },
  { label: "通知成功", value: 1 },
  { label: "已发送", value: 2 }
];

//订单状态
const OrderStatusList = [
  { label: "执行中", value: 0 },
  { label: "成功", value: 1 },
  { label: "失败", value: 2 },
  { label: "缓存", value: 3 },
  { label: "已返销", value: 4 },
  { label: "备用通道重试中", value: 5 }
];

//供货单状态
const SupplierOrderStatusList = [
  { label: "执行中", value: 0 },
  { label: "成功", value: 1 },
  { label: "失败", value: 2 },
  { label: "上游未确认", value: 3 },
  { label: "已返销", value: 4 }
];

//批量超时操作类型
const BatchTimeoutType = [
  { label: "立即超时", value: 0 },
  { label: "多久后超时", value: 1 },
  { label: "什么时候超时", value: 2 },
  { label: "订单创建多久超时", value: 3 }
];

export {
  OperatorListTelecom,
  OperatorListTelecom2,
  OperatorListRights,
  OperatorListAll,
  OperatorListOilcard,
  OperatorListPowerbill,
  ProvinceList,
  UnitList,
  BusinessTypeList,
  SupplyStrategyList,
  FundOperationTypeList,
  FundOperationTypeList2,
  FundOperationTypeList3,
  NotifyStatusList,
  OrderStatusList,
  BatchTimeoutType,
  SupplierOrderStatusList
};
