import { ComponentProps, ReactNode } from 'react'

interface ListProps<Item> extends ComponentProps<'ul'> {
  renderItem: (item: Item) => ReactNode
  items: Item[]
}

export const List = <Item,>({
  className,
  renderItem,
  items
}: ListProps<Item>) => {
  return <ul className={className}>{items.map((item) => renderItem(item))}</ul>
}
