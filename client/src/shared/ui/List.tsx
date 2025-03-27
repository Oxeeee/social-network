import { ComponentProps, ReactNode } from "react";

type ListProps<Item> = {
  renderItem: (item: Item) => ReactNode;
  items: Item[];
} & ComponentProps<"ul">;

export const List = <Item,>({
  className,
  renderItem,
  items,
}: ListProps<Item>) => {
  return <ul className={className}>{items.map((item) => renderItem(item))}</ul>;
};
