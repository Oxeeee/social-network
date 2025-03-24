import { AccountBox, Feed } from '@mui/icons-material'
import { ReactNode } from 'react'

export const ROUTER_PATHS = {
  ROOT: {
    pathname: '/'
  },
  POSTS: {
    pathname: '/posts'
  },
  POST: {
    pathname: '/posts/:id'
  },
  PROFILE: {
    pathname: '/profile'
  },
  AUTH: {
    pathname: '/auth/'
  },
  LOGIN: {
    pathname: '/auth/login'
  },
  REGISTER: {
    pathname: '/auth/register'
  }
} as const

type ExtractPathnames<T> = {
  [K in keyof T]: T[K] extends { pathname: string }
    ? T[K]['pathname']
    : T[K] extends object
      ? `${ExtractPathnames<T[K]>}`
      : never
}[keyof T]

export const NAV_ROUTES: {
  name: string
  path: ExtractPathnames<typeof ROUTER_PATHS>
  icon: ReactNode
}[] = [
  {
    name: 'Новости',
    path: ROUTER_PATHS.POSTS.pathname,
    icon: <Feed />
  },
  {
    name: 'Профиль',
    path: ROUTER_PATHS.PROFILE.pathname,
    icon: <AccountBox />
  }
]
