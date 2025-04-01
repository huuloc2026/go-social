"use client"

import { useEffect, useState } from "react"
import { Post } from "@/components/post"
import { CreatePostForm } from "@/components/create-post-form"
import { useAuth } from "@/context/auth-context"
import { Button } from "@/components/ui/button"
import { CloudCog, Loader2 } from "lucide-react"

type PostType = {
  id: string
  content: string
  createdAt: string
  user: {
    id: string
    name: string
  }
  likesCount: number
  commentsCount: number
  isLiked: boolean
}

export function Feed() {
  const [posts, setPosts] = useState<PostType[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [page, setPage] = useState(1)
  const [hasMore, setHasMore] = useState(true)
  const [loadingMore, setLoadingMore] = useState(false)
  const { user } = useAuth()

  const fetchPosts = async (pageNum = 1, append = false) => {
    try {
      const res = await fetch(`/posts/`)

      if (!res.ok) {
        throw new Error("Failed to fetch posts")
      }

      const data = await res.json()
      
      if (append) {
        setPosts((prev) => [...prev, ...data.posts])
      } else {
        setPosts(data.posts)
      }

      setHasMore(data.hasMore)
      setError(null)
    } catch (error) {
      console.error("Error fetching posts:", error)
      setError("Failed to load posts. Please try again.")
    } finally {
      setLoading(false)
      setLoadingMore(false)
    }
  }

  useEffect(() => {
    fetchPosts()
  }, [])

  const handleLoadMore = () => {
    if (loadingMore || !hasMore) return

    setLoadingMore(true)
    setPage((prev) => prev + 1)
    fetchPosts(page + 1, true)
  }

  const handlePostCreated = (newPost: PostType) => {
    setPosts((prev) => [newPost, ...prev])
  }

  const handlePostLiked = (postId: string, liked: boolean) => {
    setPosts((prev) =>
      prev.map((post) =>
        post.id === postId
          ? {
              ...post,
              isLiked: liked,
              likesCount: liked ? post.likesCount + 1 : post.likesCount - 1,
            }
          : post,
      ),
    )
  }

  if (loading) {
    return (
      <div className="flex justify-center py-10">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    )
  }

  if (error) {
    return (
      <div className="py-10 text-center">
        <p className="text-red-500">{error}</p>
        <Button
          variant="outline"
          className="mt-4"
          onClick={() => {
            setLoading(true)
            fetchPosts()
          }}
        >
          Try Again
        </Button>
      </div>
    )
  }

  return (
    <div className="space-y-8">
      {user && <CreatePostForm onPostCreated={handlePostCreated} />}

      {posts.length === 0 ? (
        <div className="py-10 text-center">
          <p className="text-muted-foreground">No posts yet</p>
        </div>
      ) : (
        <div className="space-y-6">
          {posts.map((post) => (
            <Post key={post.id} post={post} onLike={handlePostLiked} />
          ))}

          {hasMore && (
            <div className="flex justify-center py-4">
              <Button variant="outline" onClick={handleLoadMore} disabled={loadingMore}>
                {loadingMore ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" /> Loading
                  </>
                ) : (
                  "Load More"
                )}
              </Button>
            </div>
          )}
        </div>
      )}
    </div>
  )
}

