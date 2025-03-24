import { CircularProgress } from "@mui/material";
import styles from "./Loader.module.css";

export const Loader = () => {
  return <CircularProgress className={styles.loader} disableShrink />;
};
