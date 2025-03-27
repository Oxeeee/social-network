import { Container } from "@mui/material";
import { Header } from "@/shared/ui/Header";
import { Navigate, Outlet } from "react-router-dom";
import { useSelector } from "react-redux";
import { userSelectors } from "@/entites/user/store/userSlice";
import { ROUTER_PATHS } from "@/shared/routes";

export const BaseLayout = () => {
  const isAuthorized = useSelector(userSelectors.isAuthorized);

  if (!isAuthorized) {
    return <Navigate to={ROUTER_PATHS.REGISTER.pathname} />;
  }

  return (
    <>
      <Header />
      <Container
        sx={{
          position: "relative",
        }}
      >
        <Outlet />
      </Container>
    </>
  );
};
