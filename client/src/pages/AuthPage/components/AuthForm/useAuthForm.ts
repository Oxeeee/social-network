import { useDidMount } from "@/shared/lib/react/useDidMount";
import { zodResolver } from "@hookform/resolvers/zod/src/zod.js";

import { useForm } from "react-hook-form";
import { z } from "zod";

const containUppercaseRegexp = /[A-Z]/;

const createAuthSchema = (isLogin: boolean) => {
  const username = isLogin
    ? z.string().optional()
    : z.string().nonempty("Не должен быть пустым");

  return z.object({
    email: z.string().email("Неверный email"),
    username,
    password: z
      .string()
      .min(5, "Должен содержать минимум 5 знаков")
      .regex(containUppercaseRegexp, "Должен содержать хотя бы одну заглавную"),
  });
};

export type AuthFormFields = z.infer<ReturnType<typeof createAuthSchema>>;

export const useAuthForm = (isLogin: boolean) => {
  const { control, handleSubmit, reset, unregister } = useForm<AuthFormFields>({
    resolver: zodResolver(createAuthSchema(isLogin)),
    defaultValues: {
      username: "",
      password: "",
      email: "",
    },
  });

  useDidMount(() => {
    if (isLogin) {
      unregister("username");
    }
  });

  return { control, handleSubmit, reset };
};
