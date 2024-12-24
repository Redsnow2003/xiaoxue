import { http } from "@/utils/http";
import { baseUrlApi, type ResultTable, type Result } from "./utils";

/** 获取系统管理-用户管理列表 */
export const getUserList = (data?: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/get-user-list"), {
    data
  });
};

/** 系统管理-用户管理-新增用户 */
export const addUser = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/user"), { data });
};

/** 系统管理-用户管理-编辑用户 */
export const updateUser = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/user"), { data });
};

/** 系统管理-用户管理-删除用户 */
export const deleteUser = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/user"), { data });
};

/** 系统管理-用户管理-批量删除用户 */
export const deleteBatchUser = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/batch-user"), { data });
};

/** 系统管理-用户管理-修改用户状态 */
export const updateUserStatus = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/user-status"), { data });
};

/** 系统管理-用户管理-重置用户密码 */
export const resetUserPassword = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/user-password"), { data });
};

/** 系统管理-用户管理-用户分配角色 */
export const updateUserRole = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/user-role"), { data });
};

/** 系统管理-用户管理-获取所有角色列表 */
export const getAllRoleList = () => {
  return http.request<Result>("get", baseUrlApi("/list-all-role"));
};

/** 系统管理-用户管理-根据userId，获取对应角色id列表（userId：用户id） */
export const getRoleIds = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/list-role-ids"), { data });
};

/** 获取系统管理-角色管理列表 */
export const getRoleList = (data?: object) => {
  return http.request<ResultTable>("get", baseUrlApi("/role"), { data });
};

/** 系统管理-角色管理-新增角色 */
export const addRole = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/role"), { data });
};

/** 系统管理-角色管理-编辑角色 */
export const updateRole = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/role"), { data });
};

/** 系统管理-角色管理-删除角色 */
export const deleteRole = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/role"), { data });
};

/** 系统管理-角色管理-修改角色状态 */
export const updateRoleStatus = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/role-status"), { data });
};

/** 获取系统管理-菜单管理列表 */
export const getMenuList = (data?: object) => {
  return http.request<Result>("get", baseUrlApi("/menu"), { data });
};

/**系统管理-菜单管理-新增菜单 */
export const addMenu = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/menu"), { data });
};

/**系统管理-菜单管理-编辑菜单 */
export const updateMenu = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/menu"), { data });
};

/**系统管理-菜单管理-删除菜单 */
export const deleteMenu = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/menu"), { data });
};

/** 获取系统管理-部门管理列表 */
export const getDeptList = (data?: object) => {
  return http.request<Result>("get", baseUrlApi("/dept"), { data });
};
/**系统管理-部门管理-新增部门 */
export const addDept = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/dept"), { data });
};
/**系统管理-部门管理-编辑部门 */
export const updateDept = (data?: object) => {
  return http.request<Result>("put", baseUrlApi("/dept"), { data });
};
/**系统管理-部门管理-删除部门 */
export const deleteDept = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/dept"), { data });
};

/** 获取系统监控-在线用户列表 */
export const getOnlineLogsList = (data?: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/online-logs"), {
    data
  });
};

/** 获取系统监控-登录日志列表 */
export const getLoginLogsList = (data?: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/login-logs"), { data });
};

/** 删除系统监控-登录日志列表 */
export const deleteLoginLogs = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/login-logs"), { data });
};

/** 获取系统监控-操作日志列表 */
export const getOperationLogsList = (data?: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/operation-logs"), {
    data
  });
};

/** 删除操作日志 */
export const deleteOperationLogs = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/operation-logs"), {
    data
  });
};

/** 获取系统监控-系统日志列表 */
export const getSystemLogsList = (data?: object) => {
  return http.request<ResultTable>("post", baseUrlApi("/system-logs"), {
    data
  });
};

/** 删除系统监控-系统日志列表 */
export const deleteSystemLogs = (data?: object) => {
  return http.request<Result>("delete", baseUrlApi("/system-logs"), { data });
};

/** 获取系统监控-系统日志-根据 id 查日志详情 */
export const getSystemLogsDetail = (data?: object) => {
  return http.request<Result>("post", "/system-logs-detail", { data });
};

/** 获取角色管理-权限-菜单权限 */
export const getRoleMenu = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/role-menu"), { data });
};

/** 获取角色管理-权限-菜单权限-根据角色 id 查对应菜单 */
export const getRoleMenuIds = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/role-menu-ids"), { data });
};

/**更新角色管理-权限-菜单权限-根据角色id更新对应菜单 */
export const updateRoleMenuIds = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/update-role-menu"), {
    data
  });
};

/** 系统管理-用户管理-上传用户头像 */
export const uploadAvatar = (data?: object) => {
  return http.request<Result>("post", baseUrlApi("/upload-avatar"), { data });
};
