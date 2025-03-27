import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import {
  LoginPayload,
  LoginResponse,
  RegisterPayload,
  RegisterResponse,
} from "@/shared/api/auth/types";
import { authApi } from "@/shared/api/auth/authRepository";
import { User } from "../types";

type UserState = {
  user: User;
  isLoading: boolean;
  isError: boolean;
}

const initialState: UserState = {
  user: {
    username: "",
    surname: "",
    name: "",
    accessToken: "",
    photo: "",
  },
  isLoading: false,
  isError: false,
};

const userState: UserState = !localStorage.getItem("userInfo")
  ? initialState
  : JSON.parse(localStorage.getItem("userInfo")!);

const registerUser = createAsyncThunk<RegisterResponse, RegisterPayload>(
  "users/registerUser",
  async (params) => {
    const response = await authApi.register({
      params,
    });

    return response.data;
  },
);

const loginUser = createAsyncThunk<LoginResponse, LoginPayload>(
  "users/loginUser",
  async (params) => {
    const response = await authApi.login({
      params,
    });

    return response.data;
  },
);

export const userSlice = createSlice({
  name: "user",
  initialState: userState,
  reducers: {
    logOut: (state) => {
      state.user.username = "";
      state.user.accessToken = "";
      localStorage.removeItem("userInfo");
    },
  },
  selectors: {
    isAuthorized: (state) => !!state.user.accessToken,
    getFullname: (state) => state.user.name + state.user.surname,
  },
  extraReducers(builder) {
    builder.addCase(registerUser.pending, (state) => {
      state.isLoading = true;
    });
    builder.addCase(registerUser.fulfilled, (state) => {
      state.isLoading = false;
    });
    builder.addCase(loginUser.fulfilled, (state, { payload: { data } }) => {
      state.user = data;
      state.isLoading = false;
      localStorage.setItem("userInfo", JSON.stringify(state));
    });
    builder.addCase(loginUser.pending, (state) => {
      state.isLoading = true;
    });
  },
});

export const {
  actions: userActions,
  reducer: userReducer,
  selectors: userSelectors,
} = userSlice;

export const userThunks = {
  registerUser,
  loginUser,
};
