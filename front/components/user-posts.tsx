"use client"

import { useState, useEffect } from "react"
import { getUserPosts } from "@/lib/api/posts"
import { PostCard } from "@/components/post-card"
import { Loader2 } from "lucide-react"
import type { Post } from "@/types"
import { useToast } from "@/components/ui/use-toast"

interface UserPostsProps {
  userId: string
}

export function UserPosts({ userId }: UserPostsProps) {
  const { toast } = useToast()
  const [posts, setPosts] = useState<Post[]>([])
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        setIsLoading(true)
        const data = await getUserPosts(userId)
        setPosts(data.posts)
      } catch (error) {
        toast({
          title: "Error",
          description: "Failed to load user posts",
          variant: "destructive",
        })
      } finally {
        setIsLoading(false)
      }
    }

    fetchPosts()
  }, [userId, toast])

  const handlePostUpdated = (updatedPost: Post) => {
    setPosts((prevPosts) => prevPosts.map((post) => (post.id === updatedPost.id ? updatedPost : post)))
  }

  const handlePostDeleted = (postId: string) => {
    setPosts((prevPosts) => prevPosts.filter((post) => post.id !== postId))
  }

  if (isLoading) {
    return (
      <div className="flex justify-center py-8">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    )
  }

  if (posts.length === 0) {
    return (
      <div className="text-center py-8">
        <p className="text-muted-foreground">No posts to display</p>
      </div>
    )
  }

  return (
    <div className="space-y-4">
      {posts.map((post) => (
        <PostCard key={post.id} post={post} onPostUpdated={handlePostUpdated} onPostDeleted={handlePostDeleted} />
      ))}
    </div>
  )
}

