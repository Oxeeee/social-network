import { api } from '../axiosInstance'
import { AxiosRequestConfig } from '../types'
import { AuthUserPayload, AuthUserResponse, RegisterUserPayload } from './types'

class AuthApi {
  register({ params, config }: AxiosRequestConfig<RegisterUserPayload>) {
    return api.post<AuthUserResponse>('/register', params, config)
  }

  login({
    params,
    config
  }: AxiosRequestConfig<Omit<AuthUserPayload, 'username'>>) {
    return api.post<AuthUserResponse>('/auth', params, config)
  }
}

export const authApi = new AuthApi()
