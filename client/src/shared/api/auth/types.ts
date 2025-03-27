export type ApiResponse<Data> = {
  data: Data;
  message: string;
  error: string;
}

export type RegisterPayload = {
  email: string;
  password: string;
  name: string;
  surname: string;
  username: string;
}

export type RegisterResponse = ApiResponse<null>;

export type LoginPayload = Pick<RegisterPayload, "email" | "password">;

export type LoginResponse = ApiResponse<{
  accessToken: string;
  username: string;
  surname: string;
  photo: string;
  name: string;
  userId: string;
}>;
