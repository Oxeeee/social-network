import { ThemeProvider, CssBaseline } from "@mui/material";
import { NotificationsProvider } from "@toolpad/core/useNotifications";
import { RouterProvider } from "react-router-dom";
import { router } from "./router";
import { theme } from "../shared/lib/muiTheme";
import { Provider } from "react-redux";
import { store } from "./store";

export const App = () => {
  return (
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <NotificationsProvider>
          <CssBaseline />
          <RouterProvider router={router} />
        </NotificationsProvider>
      </ThemeProvider>
    </Provider>
  );
};
