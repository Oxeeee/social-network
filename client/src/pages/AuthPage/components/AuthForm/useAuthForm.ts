import { useDidMount } from "@/shared/lib/react/useDidMount";
import { zodResolver } from "@hookform/resolvers/zod/src/zod.js";

import { useForm } from "react-hook-form";
import { z } from "zod";

const containUppercaseRegexp = /[A-Z]/;

const createAuthSchema = (isLogin: boolean) => {
  const username = isLogin
    ? z.string().optional()
    : z.string().nonempty("Не должен быть пустым");
  const name = isLogin
    ? z.string().optional()
    : z.string().min(2, "Имя должно состоять минимум из 2 знаков");

  return z.object({
    username,
    name,
    email: z.string().email("Неверный email"),
    password: z
      .string()
      .min(8, "Должен содержать минимум 5 знаков")
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
      name: "",
    },
  });

  useDidMount(() => {
    if (isLogin) {
      unregister("username");
      unregister("name");
    }
  });

  return { control, handleSubmit, reset };
};
