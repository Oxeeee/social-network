import { Container, Link, Paper, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import { colors } from "../shared/lib/muiTheme";
import { ROUTER_PATHS } from "../shared/routes";

export const ErrorPage = () => {
  return (
    <Container
      sx={{
        display: "flex",
        justifyContent: "center",
        flexDirection: "column",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <Paper
        square={false}
        elevation={3}
        sx={{
          padding: 4,
          backgroundColor: colors.error.main,
          mb: 3,
        }}
      >
        <Typography variant="h3">Страница не найдена</Typography>
      </Paper>
      <Link
        variant="h4"
        color="primary"
        replace
        component={RouterLink}
        to={ROUTER_PATHS.ROOT.pathname}
      >
        На главную
      </Link>
    </Container>
  );
};
