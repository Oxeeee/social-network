import { TextField, TextFieldProps } from "@mui/material";
import {
  FieldValues,
  useController,
  UseControllerProps,
} from "react-hook-form";

type InputProps<T extends FieldValues> = {
  controller: UseControllerProps<T>;
  label: string;
} & Omit<TextFieldProps, "label" | "id">;

export const Input = <T extends FieldValues>({
  controller,
  label,
  ...props
}: InputProps<T>) => {
  const { field, fieldState } = useController(controller);

  const errorText = fieldState.error ? fieldState.error.message : "";

  return (
    <TextField
      id={controller.name}
      label={label}
      error={!!fieldState.error}
      helperText={errorText}
      variant="standard"
      {...field}
      {...props}
    />
  );
};
