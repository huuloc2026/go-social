"use client"

import type React from "react"

import { useState } from "react"
import Link from "next/link"
import Image from "next/image"
import { formatDistanceToNow } from "date-fns"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import { useToast } from "@/components/ui/use-toast"
import { Heart, MessageCircle, MoreHorizontal, Pencil, Trash2 } from "lucide-react"
import { likePost, unlikePost, createComment, deletePost, updatePost } from "@/lib/api/posts"
import { useAuth } from "@/hooks/use-auth"
import type { Post, Comment } from "@/types"

interface PostCardProps {
  post: Post
  onPostUpdated: (post: Post) => void
  onPostDeleted: (postId: string) => void
}

export function PostCard({ post, onPostUpdated, onPostDeleted }: PostCardProps) {
  const { user } = useAuth()
  const { toast } = useToast()
  const [isLiked, setIsLiked] = useState(post.isLiked)
  const [likeCount, setLikeCount] = useState(post.likeCount)
  const [showComments, setShowComments] = useState(false)
  const [newComment, setNewComment] = useState("")
  const [comments, setComments] = useState<Comment[]>(post.comments || [])
  const [isSubmittingComment, setIsSubmittingComment] = useState(false)
  const [isEditing, setIsEditing] = useState(false)
  const [editedContent, setEditedContent] = useState(post.content)
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false)
  const [isSubmittingEdit, setIsSubmittingEdit] = useState(false)

  const handleLikeToggle = async () => {
    try {
      if (isLiked) {
        await unlikePost(post.id)
        setLikeCount((prev) => prev - 1)
      } else {
        await likePost(post.id)
        setLikeCount((prev) => prev + 1)
      }
      setIsLiked(!isLiked)
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to update like status",
        variant: "destructive",
      })
    }
  }

  const handleCommentSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newComment.trim()) return

    setIsSubmittingComment(true)
    try {
      const comment = await createComment(post.id, newComment)
      setComments((prev) => [...prev, comment])
      setNewComment("")

      // Update the post with the new comment
      onPostUpdated({
        ...post,
        comments: [...comments, comment],
        commentCount: post.commentCount + 1,
      })
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to post comment",
        variant: "destructive",
      })
    } finally {
      setIsSubmittingComment(false)
    }
  }

  const handleEditSubmit = async () => {
    if (!editedContent.trim() || editedContent === post.content) {
      setIsEditing(false)
      return
    }

    setIsSubmittingEdit(true)
    try {
      const updatedPost = await updatePost(post.id, editedContent)
      onPostUpdated({
        ...post,
        content: editedContent,
        updatedAt: updatedPost.updatedAt,
      })
      setIsEditing(false)
      toast({
        title: "Success",
        description: "Post updated successfully",
      })
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to update post",
        variant: "destructive",
      })
    } finally {
      setIsSubmittingEdit(false)
    }
  }

  const handleDeletePost = async () => {
    try {
      await deletePost(post.id)
      onPostDeleted(post.id)
      toast({
        title: "Success",
        description: "Post deleted successfully",
      })
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to delete post",
        variant: "destructive",
      })
    }
  }

  const isAuthor = user?.id === post.author.id

  return (
    <Card>
      <CardHeader className="flex flex-row items-start space-y-0 p-4">
        <Link href={`/profile/${post.author.id}`} className="flex items-center gap-2">
          <Avatar>
            <AvatarImage src={post.author.avatar || `/placeholder.svg?height=40&width=40`} alt={post.author.name} />
            <AvatarFallback>{post.author.name.charAt(0)}</AvatarFallback>
          </Avatar>
          <div>
            <div className="font-semibold">{post.author.name}</div>
            <div className="text-xs text-muted-foreground">
              {formatDistanceToNow(new Date(post.createdAt), { addSuffix: true })}
              {post.updatedAt !== post.createdAt && " (edited)"}
            </div>
          </div>
        </Link>

        {isAuthor && (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon" className="ml-auto">
                <MoreHorizontal className="h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuItem onClick={() => setIsEditing(true)}>
                <Pencil className="h-4 w-4 mr-2" />
                Edit
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={() => setIsDeleteDialogOpen(true)}
                className="text-destructive focus:text-destructive"
              >
                <Trash2 className="h-4 w-4 mr-2" />
                Delete
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        )}
      </CardHeader>

      <CardContent className="p-4 pt-0">
        {isEditing ? (
          <div className="space-y-2">
            <Textarea
              value={editedContent}
              onChange={(e) => setEditedContent(e.target.value)}
              className="min-h-[100px]"
            />
            <div className="flex justify-end gap-2">
              <Button
                variant="outline"
                size="sm"
                onClick={() => {
                  setIsEditing(false)
                  setEditedContent(post.content)
                }}
              >
                Cancel
              </Button>
              <Button size="sm" onClick={handleEditSubmit} disabled={isSubmittingEdit}>
                Save
              </Button>
            </div>
          </div>
        ) : (
          <div className="space-y-4">
            <p className="whitespace-pre-line">{post.content}</p>
            {post.image && (
              <div className="relative rounded-md overflow-hidden">
                <Image
                  src={post.image || "/placeholder.svg"}
                  alt="Post image"
                  width={600}
                  height={400}
                  className="object-cover max-h-[400px] w-full"
                />
              </div>
            )}
          </div>
        )}
      </CardContent>

      <CardFooter className="flex flex-col p-4 pt-0">
        <div className="flex items-center justify-between w-full">
          <div className="flex items-center gap-2">
            <Button
              variant="ghost"
              size="sm"
              className={`flex items-center gap-1 ${isLiked ? "text-primary" : ""}`}
              onClick={handleLikeToggle}
            >
              <Heart className={`h-4 w-4 ${isLiked ? "fill-primary" : ""}`} />
              <span>{likeCount}</span>
            </Button>
            <Button
              variant="ghost"
              size="sm"
              className="flex items-center gap-1"
              onClick={() => setShowComments(!showComments)}
            >
              <MessageCircle className="h-4 w-4" />
              <span>{post.commentCount}</span>
            </Button>
          </div>
        </div>

        {showComments && (
          <div className="w-full mt-4 space-y-4">
            {comments.length > 0 ? (
              <div className="space-y-3">
                {comments.map((comment) => (
                  <div key={comment.id} className="flex gap-2">
                    <Avatar className="h-8 w-8">
                      <AvatarImage
                        src={comment.author.avatar || `/placeholder.svg?height=32&width=32`}
                        alt={comment.author.name}
                      />
                      <AvatarFallback>{comment.author.name.charAt(0)}</AvatarFallback>
                    </Avatar>
                    <div className="flex-1">
                      <div className="bg-muted p-2 rounded-md">
                        <div className="flex justify-between items-start">
                          <Link href={`/profile/${comment.author.id}`} className="font-semibold text-sm">
                            {comment.author.name}
                          </Link>
                          <span className="text-xs text-muted-foreground">
                            {formatDistanceToNow(new Date(comment.createdAt), {
                              addSuffix: true,
                            })}
                          </span>
                        </div>
                        <p className="text-sm mt-1">{comment.content}</p>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <p className="text-sm text-muted-foreground text-center py-2">
                No comments yet. Be the first to comment!
              </p>
            )}

            <form onSubmit={handleCommentSubmit} className="flex gap-2">
              <Avatar className="h-8 w-8">
                <AvatarImage src={user?.avatar || `/placeholder.svg?height=32&width=32`} alt={user?.name} />
                <AvatarFallback>{user?.name.charAt(0)}</AvatarFallback>
              </Avatar>
              <div className="flex-1 flex gap-2">
                <Textarea
                  placeholder="Write a comment..."
                  value={newComment}
                  onChange={(e) => setNewComment(e.target.value)}
                  className="min-h-[60px] flex-1"
                />
                <Button
                  type="submit"
                  size="sm"
                  className="self-end"
                  disabled={!newComment.trim() || isSubmittingComment}
                >
                  Post
                </Button>
              </div>
            </form>
          </div>
        )}
      </CardFooter>

      <AlertDialog open={isDeleteDialogOpen} onOpenChange={setIsDeleteDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Are you sure?</AlertDialogTitle>
            <AlertDialogDescription>
              This action cannot be undone. This will permanently delete your post and remove it from the server.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancel</AlertDialogCancel>
            <AlertDialogAction onClick={handleDeletePost} className="bg-destructive text-destructive-foreground">
              Delete
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </Card>
  )
}

