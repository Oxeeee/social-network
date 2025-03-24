import { useState } from "react";
import { Sidebar } from "./Sidebar";
import { AppBar, IconButton, Toolbar, Typography } from "@mui/material";
import { Menu, Send } from "@mui/icons-material";

export const Header = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  

  const handleToggleSidebar = () => {
    setIsSidebarOpen((isOpen) => !isOpen);
  };

  return (
    <>
      <AppBar
        sx={{
          mb: 3,
        }}
        position="relative"
      >
        <Toolbar
          sx={{
            display: "flex",
            justifyContent: "space-between",
          }}
        >
          <Typography
            display={"flex"}
            alignItems={"center"}
            gap={1}
            variant="h5"
            noWrap
            component="div"
          >
            <Send />
            Connectify
          </Typography>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleToggleSidebar}
            edge="start"
            sx={[isSidebarOpen && { display: "none" }]}
          >
            <Menu />
          </IconButton>
        </Toolbar>
      </AppBar>
      <Sidebar isOpen={isSidebarOpen} handleToggle={handleToggleSidebar} />
    </>
  );
};
