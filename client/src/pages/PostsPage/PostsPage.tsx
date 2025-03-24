import Card from '@/shared/ui/Card'
import { Typography } from '@mui/material'
import { useAppSelector } from '@/shared/lib/redux'
import { postsSelectors } from './store/postsSlice'
import { List } from '@/shared/ui/List'

import styles from './PostsPage.module.css'
import { generatePath, useNavigate } from 'react-router-dom'
import { ROUTER_PATHS } from '@/shared/routes'

export const PostsPage = () => {
  const navigate = useNavigate()
  const posts = useAppSelector(postsSelectors.getPosts)

  const handleClick = (id: number) => {
    navigate(generatePath(ROUTER_PATHS.POST.pathname, { id: String(id) }))
  }

  return (
    <div>
      <Typography mb={3} variant="h2">
        Новости
      </Typography>

      <List
        className={styles.postsList}
        items={posts}
        renderItem={(post) => (
          <Card
            onClick={() => handleClick(post.id)}
            key={post.id}
            title={post.title}
            description={post.description}
          />
        )}
      />
    </div>
  )
}
