import { createBrowserRouter, Navigate, Outlet } from 'react-router-dom'
import { BaseLayout } from './layout/BaseLayout'
import { ROUTER_PATHS } from '../shared/routes'
import { lazy } from 'react'
import { withSuspense } from '../shared/lib/react/withSuspense'
import { Loader } from '../shared/ui/Loader/Loader'
import { AuthPage } from '@/pages/AuthPage'
import { ErrorPage } from '../pages/ErrorPage'
import { ScreenCenteredBox } from '../shared/ui/ScreenCenteredBox'
import { PostPage } from '@/pages/PostPage'
import { PostsPage } from '@/pages/PostsPage'

const LazyProfilePage = lazy(() =>
  import('../pages/ProfilePage/ProfilePage').then((module) => ({
    default: module.ProfilePage
  }))
)

export const router = createBrowserRouter([
  {
    path: ROUTER_PATHS.ROOT.pathname,
    element: <BaseLayout />,
    children: [
      {
        index: true,
        element: <Navigate replace to={ROUTER_PATHS.POSTS.pathname} />
      },
      {
        path: ROUTER_PATHS.POSTS.pathname,
        element: <PostsPage />
      },
      {
        path: ROUTER_PATHS.POST.pathname,
        element: <PostPage />
      },
      {
        path: ROUTER_PATHS.PROFILE.pathname,
        element: withSuspense(<LazyProfilePage />, <Loader />)
      }
    ]
  },
  {
    path: ROUTER_PATHS.AUTH.pathname,
    element: (
      <ScreenCenteredBox>
        <Outlet />
      </ScreenCenteredBox>
    ),
    children: [
      {
        index: true,
        element: <Navigate to={ROUTER_PATHS.LOGIN.pathname} />
      },
      {
        path: ROUTER_PATHS.LOGIN.pathname,
        element: <AuthPage isLogin={true} />
      },
      {
        path: ROUTER_PATHS.REGISTER.pathname,
        element: <AuthPage isLogin={false} />
      }
    ]
  },
  {
    path: '*',
    element: <ErrorPage />
  }
])
