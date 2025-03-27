import { AccountCircle, Email, Key } from "@mui/icons-material";
import { Box, Button, InputAdornment, Link, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import { colors } from "@/shared/lib/muiTheme";
import { AuthFormFields, useAuthForm } from "./useAuthForm";
import { Input } from "@/shared/ui/Input";
import { SubmitHandler } from "react-hook-form";
import { ROUTER_PATHS } from "@/shared/routes";
import { useAppSelector } from "@/shared/lib/redux";

type AuthFormProps = {
  onSubmit: (values: AuthFormFields) => void;
  isLogin: boolean;
}

export const AuthForm = ({ onSubmit, isLogin }: AuthFormProps) => {
  const isSubmitLoading = useAppSelector((state) => state.user.isLoading);

  const { control, handleSubmit, reset } = useAuthForm(isLogin);

  const handleFormSubmit: SubmitHandler<AuthFormFields> = (formFields) => {
    onSubmit(formFields);
    reset();
  };

  return (
    <div>
      <Typography variant="h4" mb={1}>
        {isLogin ? "Войти" : "Регистрация"}
      </Typography>
      <Box
        onSubmit={handleSubmit(handleFormSubmit)}
        component="form"
        bgcolor={colors.divider}
        borderRadius={3}
        display="flex"
        flexDirection="column"
        gap={3}
        width="fit-content"
        py={3}
        px={8}
        boxShadow={3}
      >
        {!isLogin && (
          <>
            <Box display="flex" gap={3}>
              <Input
                label="Введите имя"
                controller={{
                  control,
                  name: "name",
                }}
                autoComplete="off"
              />
              <Input
                label="Введите фамилию"
                controller={{
                  control,
                  name: "surname",
                }}
                autoComplete="off"
              />
            </Box>
            <Input
              label="Введите логин"
              controller={{
                control,
                name: "username",
              }}
              slotProps={{
                input: {
                  startAdornment: (
                    <InputAdornment position="start">
                      <AccountCircle />
                    </InputAdornment>
                  ),
                },
              }}
              autoComplete="off"
            />
          </>
        )}

        <Input
          label="Введите почту"
          controller={{
            control,
            name: "email",
          }}
          slotProps={{
            input: {
              startAdornment: (
                <InputAdornment position="start">
                  <Email />
                </InputAdornment>
              ),
            },
          }}
          autoComplete="off"
        />

        <Input
          label="Введите пароль"
          controller={{
            control,
            name: "password",
          }}
          slotProps={{
            input: {
              startAdornment: (
                <InputAdornment position="start">
                  <Key />
                </InputAdornment>
              ),
            },
          }}
          autoComplete="off"
        />
        <Button
          loading={isSubmitLoading}
          disabled={isSubmitLoading}
          type="submit"
          variant="outlined"
        >
          Подтвердить
        </Button>
        <Link
          variant="subtitle2"
          color="primary"
          component={RouterLink}
          sx={{
            marginLeft: "auto",
          }}
          to={
            isLogin
              ? ROUTER_PATHS.REGISTER.pathname
              : ROUTER_PATHS.LOGIN.pathname
          }
        >
          {isLogin ? "Не зарегистрирован?" : "Уже зарегистрирован?"}
        </Link>
      </Box>
    </div>
  );
};
