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
