import { useAppSelector } from "@/shared/lib/redux";
import { Avatar, Box, Button, IconButton, Typography } from "@mui/material";
import { Upload } from "@mui/icons-material";

import styles from "./ProfilePage.module.css";

export const ProfilePage = () => {
  const username = useAppSelector((state) => state.user.username);

  const handleFileUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files as FileList;

    // TODO: делать запрос на добавление аватарки
    if (files[0]) {
      const reader = new FileReader();
      reader.readAsDataURL(files[0]);

      reader.onload = () => {
        console.log(reader.result);
      };
    }
  };

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
        <div>
          <IconButton className={styles.addProfileIcon} aria-label="add-circle">
            <Upload />
          </IconButton>
          <Button
            className={styles.hiddenInput}
            component="label"
            role={undefined}
            tabIndex={-1}
          >
            <input type="file" onChange={handleFileUpload} />
          </Button>
        </div>
      </Box>
      <Typography variant="h2">{username}</Typography>
    </Box>
  );
};
