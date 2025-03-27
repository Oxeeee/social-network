import { useDidMount } from "@/shared/lib/react/useDidMount";
import { zodResolver } from "@hookform/resolvers/zod/src/zod.js";

import { useForm } from "react-hook-form";
import { z, ZodString } from "zod";

const containUppercaseRegexp = /[A-Z]/;

const createFieldOptional = (field: ZodString, isOptional: boolean) => {
  return isOptional ? z.string().optional() : field;
};

const createAuthSchema = (isLogin: boolean) => {
  const name = createFieldOptional(
    z.string().min(2, "Имя должно состоять минимум из 2 знаков"),
    isLogin,
  );
  const surname = createFieldOptional(
    z.string().min(2, "Фамилия должна состоять минимум из 2 знаков"),
    isLogin,
  );
  const username = createFieldOptional(
    z.string().nonempty("Обязательное поле"),
    isLogin,
  );

  return z.object({
    username,
    name,
    surname,
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
      surname: "",
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
