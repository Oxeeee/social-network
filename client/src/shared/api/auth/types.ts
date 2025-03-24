export interface RegisterUserPayload {
  email: string
  username: string
  password: string
}

export type AuthUserPayload = Omit<
  RegisterUserPayload,
  'username'
>

export interface AuthUserResponse {
  token: string
  data: {
    email: string
    username: string
    password: string
    id: number
  }
}
