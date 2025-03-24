import { useAppSelector } from '@/shared/lib/redux'
import { useParams } from 'react-router-dom'
import { postsSelectors } from './PostsPage/store/postsSlice'

export const PostPage = () => {
  const { id } = useParams<{ id: string }>()
  const post = useAppSelector((state) =>
    postsSelectors.getPostById(state, Number(id))
  )

  return <h1>{post?.title ?? 'Пусто'}</h1>
}
