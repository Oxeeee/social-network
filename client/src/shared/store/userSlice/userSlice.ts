import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import {
  AuthUserPayload,
  AuthUserResponse,
  RegisterUserPayload
} from '@/shared/api/auth/types'
import { authApi } from '@/shared/api/auth/AuthApi'

interface UserState {
  username: string
  userId: number | null
  token: string | null
  avatar: string
  isLoading: boolean
  isError: boolean
}
const initialState = {
  username: '',
  userId: null,
  token: null,
  isLoading: false,
  avatar: '',
  isError: false
}

const userState: UserState = !localStorage.getItem('userInfo')
  ? initialState
  : JSON.parse(localStorage.getItem('userInfo')!)

const registerUser = createAsyncThunk<AuthUserResponse, RegisterUserPayload>(
  'users/registerUser',
  async (params) => {
    const response = await authApi.register({
      params
    })

    return response.data
  }
)

const loginUser = createAsyncThunk<AuthUserResponse, AuthUserPayload>(
  'users/loginUser',
  async (params) => {
    const response = await authApi.login({
      params
    })

    return response.data
  }
)

export const userSlice = createSlice({
  name: 'user',
  initialState: userState,
  reducers: {
    logOut: (state) => {
      state.username = ''
      state.token = null
      localStorage.removeItem('userInfo')
    }
  },
  selectors: {
    isAuthorized: (state) => !!state.token
  },
  extraReducers(builder) {
    builder.addCase(registerUser.pending, (state) => {
      state.isLoading = true
    }),
      builder.addCase(registerUser.fulfilled, (state) => {
        state.isLoading = false
      }),
      builder.addCase(loginUser.fulfilled, (state, action) => {
        state.token = action.payload.token
        state.username = action.payload.data.username
        state.avatar = action.payload.data.avatarUrl
        state.isLoading = false
        localStorage.setItem('userInfo', JSON.stringify(state))
      }),
      builder.addCase(loginUser.pending, (state) => {
        state.isLoading = true
      })
  }
})

export const {
  actions: userActions,
  reducer: userReducer,
  selectors: userSelectors
} = userSlice

export const userThunks = {
  registerUser,
  loginUser
}
