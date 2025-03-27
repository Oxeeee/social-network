import { Container, Paper, Typography, Button } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import { colors } from "../shared/lib/muiTheme";
import { ROUTER_PATHS } from "../shared/routes";

export const ErrorPage = () => {
  return (
    <Container
      maxWidth="sm"
      sx={{
        display: "flex",
        justifyContent: "center",
        flexDirection: "column",
        alignItems: "center",
        height: "100vh",
        textAlign: "center",
      }}
    >
      <Paper
        elevation={8}
        sx={{
          padding: 4,
          backgroundColor: "background.paper",
          borderRadius: 3,
          mb: 4,
          borderLeft: `6px solid ${colors.error.main}`,
          transform: "scale(1)",
          transition: "transform 0.3s ease",
          "&:hover": {
            transform: "scale(1.02)",
          },
        }}
      >
        <Typography
          variant="h3"
          component="h1"
          sx={{
            mb: 2,
            fontWeight: 700,
            color: colors.error.main,
            letterSpacing: "0.05em",
          }}
        >
          404 Error
        </Typography>
        <Typography variant="h5" component="h2" sx={{ mb: 2 }}>
          Такой страницы не существует :(
        </Typography>
        <Typography variant="body1" sx={{ color: "text.secondary" }}>
          Данная страница была удалена, либо её никогда не существовало.
        </Typography>
      </Paper>

      <Button
        variant="contained"
        color="primary"
        component={RouterLink}
        to={ROUTER_PATHS.ROOT.pathname}
        size="large"
        sx={{
          px: 4,
          py: 1.5,
          borderRadius: 2,
          fontWeight: 600,
          textTransform: "none",
          boxShadow: 3,
          "&:hover": {
            boxShadow: 6,
          },
        }}
      >
        Вернуться на главную
      </Button>
    </Container>
  );
};
