"use client"

import type React from "react"

import { useState } from "react"
import { formatDistanceToNow } from "date-fns"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"
import { Heart, MessageCircle, Send } from "lucide-react"
import { useAuth } from "@/context/auth-context"
import { cn } from "@/lib/utils"

type Comment = {
  id: string
  content: string
  createdAt: string
  user: {
    id: string
    name: string
  }
}

type PostProps = {
  post: {
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
  onLike: (postId: string, liked: boolean) => void
}

export function Post({ post, onLike }: PostProps) {
  const [comment, setComment] = useState("")
  const [comments, setComments] = useState<Comment[]>([])
  const [showComments, setShowComments] = useState(false)
  const [loadingComments, setLoadingComments] = useState(false)
  const [submittingComment, setSubmittingComment] = useState(false)
  const { user } = useAuth()

  const handleLike = async () => {
    if (!user) return

    try {
      const method = post.isLiked ? "DELETE" : "POST"
      const res = await fetch(`/api/posts/${post.id}/like`, {
        method,
      })

      if (res.ok) {
        onLike(post.id, !post.isLiked)
      }
    } catch (error) {
      console.error("Error liking post:", error)
    }
  }

  const loadComments = async () => {
    if (loadingComments) return

    setLoadingComments(true)
    try {
      const res = await fetch(`/api/posts/${post.id}/comments`)

      if (res.ok) {
        const data = await res.json()
        setComments(data.comments)
        setShowComments(true)
      }
    } catch (error) {
      console.error("Error loading comments:", error)
    } finally {
      setLoadingComments(false)
    }
  }

  const handleToggleComments = () => {
    if (!showComments && comments.length === 0) {
      loadComments()
    } else {
      setShowComments(!showComments)
    }
  }

  const handleSubmitComment = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!user || !comment.trim() || submittingComment) return

    setSubmittingComment(true)
    try {
      const res = await fetch(`/api/posts/${post.id}/comments`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ content: comment }),
      })

      if (res.ok) {
        const data = await res.json()
        setComments((prev) => [data.comment, ...prev])
        setComment("")

        // Update comment count in the post
        onLike(post.id, post.isLiked) // Reusing the onLike function to trigger a re-render
      }
    } catch (error) {
      console.error("Error submitting comment:", error)
    } finally {
      setSubmittingComment(false)
    }
  }

  return (
    <Card>
      <CardHeader className="flex flex-row items-start gap-4 space-y-0 pb-2">
        <Avatar>
          <AvatarImage src="/placeholder-user.jpg" alt={post.user.name} />
          <AvatarFallback>{post.user.name.charAt(0).toUpperCase()}</AvatarFallback>
        </Avatar>
        <div className="space-y-1">
          <div className="font-semibold">{post.user.name}</div>
          <div className="text-xs text-muted-foreground">
            {formatDistanceToNow(new Date(post.createdAt), { addSuffix: true })}
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <p className="whitespace-pre-line">{post.content}</p>
      </CardContent>
      <CardFooter className="flex flex-col space-y-4 pt-0">
        <div className="flex items-center gap-4">
          <Button variant="ghost" size="sm" className="gap-1" onClick={handleLike} disabled={!user}>
            <Heart className={cn("h-4 w-4", post.isLiked ? "fill-red-500 text-red-500" : "")} />
            <span>{post.likesCount}</span>
          </Button>
          <Button variant="ghost" size="sm" className="gap-1" onClick={handleToggleComments}>
            <MessageCircle className="h-4 w-4" />
            <span>{post.commentsCount}</span>
          </Button>
        </div>

        {showComments && (
          <div className="w-full space-y-4">
            {user && (
              <form onSubmit={handleSubmitComment} className="flex gap-2">
                <Textarea
                  placeholder="Write a comment..."
                  value={comment}
                  onChange={(e) => setComment(e.target.value)}
                  className="min-h-[60px]"
                />
                <Button type="submit" size="icon" disabled={!comment.trim() || submittingComment}>
                  <Send className="h-4 w-4" />
                </Button>
              </form>
            )}

            <div className="space-y-4">
              {loadingComments ? (
                <div className="flex justify-center py-2">
                  <MessageCircle className="h-4 w-4 animate-pulse" />
                </div>
              ) : comments.length > 0 ? (
                comments.map((comment) => (
                  <div key={comment.id} className="flex gap-2">
                    <Avatar className="h-6 w-6">
                      <AvatarImage src="/placeholder-user.jpg" alt={comment.user.name} />
                      <AvatarFallback>{comment.user.name.charAt(0).toUpperCase()}</AvatarFallback>
                    </Avatar>
                    <div className="flex-1 rounded-lg bg-muted p-2">
                      <div className="flex items-center gap-2">
                        <span className="text-xs font-semibold">{comment.user.name}</span>
                        <span className="text-xs text-muted-foreground">
                          {formatDistanceToNow(new Date(comment.createdAt), { addSuffix: true })}
                        </span>
                      </div>
                      <p className="text-sm">{comment.content}</p>
                    </div>
                  </div>
                ))
              ) : (
                <p className="text-center text-sm text-muted-foreground">No comments yet</p>
              )}
            </div>
          </div>
        )}
      </CardFooter>
    </Card>
  )
}

