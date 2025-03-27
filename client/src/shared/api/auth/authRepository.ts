import { api } from "@/shared/api/axiosInstance";
import { AxiosRequestConfig } from "@/shared/api/types";
import {
  LoginPayload,
  LoginResponse,
  RegisterPayload,
  RegisterResponse,
} from "@/shared/api/auth/types";

class AuthRepository {
  register({ params, config }: AxiosRequestConfig<RegisterPayload>) {
    return api.post<RegisterResponse>("/register", params, config);
  }

  login({
    params,
    config,
  }: AxiosRequestConfig<Omit<LoginPayload, "username">>) {
    return api.post<LoginResponse>("/auth", params, config);
  }
}

export const authApi = new AuthRepository();
