import {
  Box,
  Button,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import { NAV_ROUTES, ROUTER_PATHS } from "../routes";
import { Link, useNavigate } from "react-router-dom";
import { useAppDispatch } from "../lib/redux";
import { userActions } from "@/entites/user/store/userSlice";

interface SidebarProps {
  handleToggle: () => void;
  isOpen: boolean;
}

export const Sidebar = ({ handleToggle, isOpen }: SidebarProps) => {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();

  const handleSignOutClick = async (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  ) => {
    e.stopPropagation();
    dispatch(userActions.logOut());
    navigate(ROUTER_PATHS.AUTH.pathname);
  };

  const DrawerList = (
    <Box
      sx={{
        width: 250,
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "100%",
      }}
      role="presentation"
      onClick={handleToggle}
    >
      <List>
        {NAV_ROUTES.map((route) => (
          <Link key={route.name} className="navlink" to={route.path}>
            <ListItem disablePadding>
              <ListItemButton>
                <ListItemIcon>{route.icon}</ListItemIcon>
                <ListItemText primary={route.name} />
              </ListItemButton>
            </ListItem>
          </Link>
        ))}
      </List>
      <Button
        variant="outlined"
        color="error"
        onClick={handleSignOutClick}
        sx={{ marginInline: "auto", mb: 2 }}
      >
        Выйти
      </Button>
    </Box>
  );

  return (
    <div>
      <Drawer anchor="right" open={isOpen} onClose={handleToggle}>
        {DrawerList}
      </Drawer>
    </div>
  );
};
