import { Box } from "@mui/material";
import { ReactNode } from "react";

export const ScreenCenteredBox = ({ children }: { children: ReactNode }) => {
  return (
    <Box
      sx={{
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      {children}
    </Box>
  );
};
