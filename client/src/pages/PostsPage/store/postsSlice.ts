import { createSlice } from '@reduxjs/toolkit'

export const postsSlice = createSlice({
  name: 'posts',
  initialState: {
    posts: [
      {
        id: 1,
        title: 'Заголовок 1',
        description: 'Описание первого поста'
      },
      {
        id: 2,
        title: 'Заголовок 2',
        description: 'Описание второго поста'
      },
      {
        id: 3,
        title: 'Заголовок 3',
        description: 'Описание третьего поста'
      },
      {
        id: 4,
        title: 'Заголовок 4',
        description: 'Описание четвертого поста'
      },
      {
        id: 5,
        title: 'Заголовок 5',
        description: 'Описание пятого поста'
      }
    ]
  },
  selectors: {
    getPosts: (state) => state.posts,
    getPostById: (state, id: number) =>
      state.posts.find((post) => post.id === id)
  },
  reducers: {}
})

export const {
  reducer: postsReducer,
  actions: postsActions,
  selectors: postsSelectors
} = postsSlice
