export const baseUrlApi = (url: string) =>
  process.env.NODE_ENV === "development"
    ? `/api${url}`
    : `http://127.0.0.1:8080${url}`;

export type Result = {
  success: boolean;
  message?: string;
  data?: Array<any>;
};

export type ResultTable = {
  success: boolean;
  data?: {
    /** 列表数据 */
    list: Array<any>;
    /** 总条目数 */
    total?: number;
    /** 每页显示条目个数 */
    pageSize?: number;
    /** 当前页数 */
    currentPage?: number;
  };
};

/** 日期、时间选择器快捷选项，常搭配 [DatePicker](https://element-plus.org/zh-CN/component/date-picker.html) 和 [DateTimePicker](https://element-plus.org/zh-CN/component/datetime-picker.html) 的`shortcuts`属性使用 */
export const getPickerShortcuts = (): Array<{
  text: string;
  value: Date | Function;
}> => {
  return [
    {
      text: "今天",
      value: () => {
        const today = new Date();
        today.setHours(0, 0, 0, 0);
        const todayEnd = new Date();
        todayEnd.setHours(23, 59, 59, 999);
        return [today, todayEnd];
      }
    },
    {
      text: "昨天",
      value: () => {
        const yesterday = new Date();
        yesterday.setDate(yesterday.getDate() - 1);
        yesterday.setHours(0, 0, 0, 0);
        const yesterdayEnd = new Date();
        yesterdayEnd.setDate(yesterdayEnd.getDate() - 1);
        yesterdayEnd.setHours(23, 59, 59, 999);
        return [yesterday, yesterdayEnd];
      }
    },
    {
      text: "前天",
      value: () => {
        const beforeYesterday = new Date();
        beforeYesterday.setDate(beforeYesterday.getDate() - 2);
        beforeYesterday.setHours(0, 0, 0, 0);
        const beforeYesterdayEnd = new Date();
        beforeYesterdayEnd.setDate(beforeYesterdayEnd.getDate() - 2);
        beforeYesterdayEnd.setHours(23, 59, 59, 999);
        return [beforeYesterday, beforeYesterdayEnd];
      }
    },
    {
      text: "本周",
      value: () => {
        const today = new Date();
        const startOfWeek = new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate() - today.getDay() + (today.getDay() === 0 ? -6 : 1)
        );
        startOfWeek.setHours(0, 0, 0, 0);
        const endOfWeek = new Date(
          startOfWeek.getTime() +
            6 * 24 * 60 * 60 * 1000 +
            23 * 60 * 60 * 1000 +
            59 * 60 * 1000 +
            59 * 1000 +
            999
        );
        return [startOfWeek, endOfWeek];
      }
    },
    {
      text: "上周",
      value: () => {
        const today = new Date();
        const startOfLastWeek = new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate() - today.getDay() - 7 + (today.getDay() === 0 ? -6 : 1)
        );
        startOfLastWeek.setHours(0, 0, 0, 0);
        const endOfLastWeek = new Date(
          startOfLastWeek.getTime() +
            6 * 24 * 60 * 60 * 1000 +
            23 * 60 * 60 * 1000 +
            59 * 60 * 1000 +
            59 * 1000 +
            999
        );
        return [startOfLastWeek, endOfLastWeek];
      }
    },
    {
      text: "本月",
      value: () => {
        const today = new Date();
        const startOfMonth = new Date(today.getFullYear(), today.getMonth(), 1);
        startOfMonth.setHours(0, 0, 0, 0);
        const endOfMonth = new Date(
          today.getFullYear(),
          today.getMonth() + 1,
          0
        );
        endOfMonth.setHours(23, 59, 59, 999);
        return [startOfMonth, endOfMonth];
      }
    },
    {
      text: "上个月",
      value: () => {
        const today = new Date();
        const startOfLastMonth = new Date(
          today.getFullYear(),
          today.getMonth() - 1,
          1
        );
        startOfLastMonth.setHours(0, 0, 0, 0);
        const endOfLastMonth = new Date(
          today.getFullYear(),
          today.getMonth(),
          0
        );
        endOfLastMonth.setHours(23, 59, 59, 999);
        return [startOfLastMonth, endOfLastMonth];
      }
    },
    {
      text: "本年",
      value: () => {
        const today = new Date();
        const startOfYear = new Date(today.getFullYear(), 0, 1);
        startOfYear.setHours(0, 0, 0, 0);
        const endOfYear = new Date(today.getFullYear(), 11, 31);
        endOfYear.setHours(23, 59, 59, 999);
        return [startOfYear, endOfYear];
      }
    }
  ];
};

//两个时间之间的秒数,格式化为天时分秒
export const timeDiff = (start1: string, end1: string) => {
  //将字符串转成日期
  let start = new Date(start1);
  let end = new Date(end1);
  let time = Math.floor((end.getTime() - start.getTime()) / 1000);
  if (time < 0) {
    return "0秒";
  }
  let day = Math.floor(time / 86400);
  let hour = Math.floor((time % 86400) / 3600);
  let minute = Math.floor((time % 3600) / 60);
  let second = Math.floor(time % 60);
  if (day > 0) {
    return `${day}天${hour}小时${minute}分钟${second}秒`;
  }
  if (hour > 0) {
    return `${hour}小时${minute}分钟${second}秒`;
  }
  if (minute > 0) {
    return `${minute}分钟${second}秒`;
  }
  return `${second}秒`;
};
