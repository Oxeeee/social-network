import { postsReducer } from '@/pages/PostsPage/store/postsSlice'
import { userReducer } from '@/shared/store/userSlice/userSlice'
import { combineReducers, configureStore } from '@reduxjs/toolkit'

const reducer = combineReducers({
  user: userReducer,
  posts: postsReducer
})

export const store = configureStore({
  reducer
})
