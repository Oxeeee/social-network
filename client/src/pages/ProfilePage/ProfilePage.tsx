import { useAppSelector } from "@/shared/lib/redux";
import { Avatar, Box, IconButton, Typography } from "@mui/material";

import styles from "./ProfilePage.module.css";
import { ChangeCircle } from "@mui/icons-material";

export const ProfilePage = () => {
  const username = useAppSelector((state) => state.user.username);

  return (
    <Box
      display="flex"
      flexDirection="column"
      justifyContent="center"
      alignItems="center"
    >
      <Box position="relative">
        <Avatar
          sx={{
            width: 96,
            height: 96,
          }}
          alt="Аватарка"
        />
        <IconButton className={styles.addProfileIcon} aria-label="add-circle">
          <ChangeCircle />
        </IconButton>
      </Box>
      <Typography variant="h2">{username}</Typography>
    </Box>
  );
};
