import {
  Card as UICard,
  CardProps as UICardProps,
  CardContent,
  Typography,
  CardActionArea,
  CardActions,
  CardMedia,
} from "@mui/material";
import { ReactNode } from "react";

export type CardProps = {
  title: string;
  description: string;
  imgSrc?: string;
  alt?: string;
  actionSlot?: ReactNode;
} & UICardProps;

export default function Card({
  title,
  description,
  actionSlot,
  imgSrc,
  alt,
  ...props
}: CardProps) {
  return (
    <UICard {...props}>
      <CardActionArea>
        <CardMedia component="img" height="140" image={imgSrc} alt={alt} />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            {title}
          </Typography>
          <Typography variant="body2" sx={{ color: "text.secondary" }}>
            {description}
          </Typography>
        </CardContent>
      </CardActionArea>
      {actionSlot && <CardActions>{actionSlot}</CardActions>}
    </UICard>
  );
}
