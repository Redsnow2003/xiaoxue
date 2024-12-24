import { http } from "@/utils/http";
import { baseUrlApi, type Result } from "./utils";

export const getAsyncRoutes = () => {
  return http.request<Result>("get", baseUrlApi("/get-async-routes"));
};
