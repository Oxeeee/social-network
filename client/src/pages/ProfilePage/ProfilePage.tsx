import { useAppSelector } from "@/shared/lib/redux";
import { Avatar, Box, Button, IconButton, Typography } from "@mui/material";
import { Upload } from "@mui/icons-material";
import styles from "./ProfilePage.module.css";
import { userSelectors } from "@/entites/user/store/userSlice";

export const ProfilePage = () => {
  const username = useAppSelector((state) => state.user.username);
  const fullname = useAppSelector(userSelectors.getFullname);

  const handleFileUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files as FileList;
    if (files[0]) {
      const reader = new FileReader();
      reader.readAsDataURL(files[0]);
      reader.onload = () => {
        console.log(reader.result);
      };
    }
  };

  return (
    <Box className={styles.profileContainer}>
      <Box className={styles.profileCard}>
        <Box className={styles.avatarContainer}>
          <Avatar className={styles.profileAvatar} alt="Profile picture" />
          <IconButton className={styles.uploadButton} component="label">
            <Upload fontSize="small" />
            <input
              className={styles.hiddenFileInput}
              accept="image/*"
              type="file"
              onChange={handleFileUpload}
            />
          </IconButton>
        </Box>

        <Typography variant="h4" className={styles.username}>
          {username}
        </Typography>

        <Typography variant="h5" className={styles.fullName}>
          {fullname || "Matvey"}
        </Typography>

        <Button color="inherit" variant="contained">
          Редактировать
        </Button>
      </Box>
    </Box>
  );
};
