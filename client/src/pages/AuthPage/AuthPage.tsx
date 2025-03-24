import { AuthForm } from './components/AuthForm/AuthForm'
import { AuthFormFields } from './components/AuthForm/useAuthForm'
import { isAxiosError } from 'axios'
import { useNotifications } from '@toolpad/core/useNotifications'
import { useAppDispatch } from '@/shared/lib/redux'
import { userThunks } from '@/shared/store/userSlice/userSlice'
import { useNavigate } from 'react-router-dom'
import { ROUTER_PATHS } from '@/shared/routes'

export const AuthPage = ({ isLogin }: { isLogin: boolean }) => {
  const navigate = useNavigate()
  const dispatch = useAppDispatch()

  const { show } = useNotifications()

  const handleRegisterSubmit = async ({
    username,
    password,
    email
  }: AuthFormFields) => {
    if (!username) return

    try {
      await dispatch(
        userThunks.registerUser({
          username,
          password,
          email,
        })
      ).unwrap()
      navigate(ROUTER_PATHS.LOGIN.pathname)
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        show('Ошибка при регистрации пользователя', {
          severity: 'error',
          autoHideDuration: 3000
        })
      }
    }
  }

  const handleLoginSubmit = async ({ password, email }: AuthFormFields) => {
    try {
      await dispatch(
        userThunks.loginUser({
          password,
          email
        })
      ).unwrap()
      navigate(ROUTER_PATHS.POSTS.pathname)
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        show('Пользователь не авторизован', {
          severity: 'error',
          autoHideDuration: 1000
        })
      }
    }
  }

  const handleSubmit = isLogin ? handleLoginSubmit : handleRegisterSubmit

  return (
    <AuthForm key={String(isLogin)} onSubmit={handleSubmit} isLogin={isLogin} />
  )
}
