"use client"

import type React from "react"

import { useState, useRef } from "react"
import Image from "next/image"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Textarea } from "@/components/ui/textarea"
import { useToast } from "@/components/ui/use-toast"
import { ImageIcon, X, Loader2 } from "lucide-react"
import { createPost } from "@/lib/api/posts"
import { useAuth } from "@/hooks/use-auth"
import type { Post } from "@/types"

interface CreatePostFormProps {
  onPostCreated: (post: Post) => void
}

export function CreatePostForm({ onPostCreated }: CreatePostFormProps) {
  const { user } = useAuth()
  const { toast } = useToast()
  const [content, setContent] = useState("")
  const [image, setImage] = useState<File | null>(null)
  const [imagePreview, setImagePreview] = useState<string | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)

  const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (!file) return

    if (file.size > 5 * 1024 * 1024) {
      toast({
        title: "File too large",
        description: "Image must be less than 5MB",
        variant: "destructive",
      })
      return
    }

    if (!file.type.startsWith("image/")) {
      toast({
        title: "Invalid file type",
        description: "Only image files are allowed",
        variant: "destructive",
      })
      return
    }

    setImage(file)
    const reader = new FileReader()
    reader.onload = () => {
      setImagePreview(reader.result as string)
    }
    reader.readAsDataURL(file)
  }

  const removeImage = () => {
    setImage(null)
    setImagePreview(null)
    if (fileInputRef.current) {
      fileInputRef.current.value = ""
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!content.trim() && !image) {
      toast({
        title: "Empty post",
        description: "Please add some text or an image to your post",
        variant: "destructive",
      })
      return
    }

    setIsSubmitting(true)
    try {
      const post = await createPost(content, image)
      onPostCreated(post)
      setContent("")
      setImage(null)
      setImagePreview(null)
      if (fileInputRef.current) {
        fileInputRef.current.value = ""
      }
      toast({
        title: "Success",
        description: "Your post has been published",
      })
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to create post. Please try again.",
        variant: "destructive",
      })
    } finally {
      setIsSubmitting(false)
    }
  }

  if (!user) return null

  return (
    <Card>
      <form onSubmit={handleSubmit}>
        <CardContent className="p-4">
          <div className="flex gap-3">
            <Avatar>
              <AvatarImage src={user.avatar || `/placeholder.svg?height=40&width=40`} alt={user.name} />
              <AvatarFallback>{user.name.charAt(0)}</AvatarFallback>
            </Avatar>
            <div className="flex-1">
              <Textarea
                placeholder={`What's on your mind, ${user.name.split(" ")[0]}?`}
                value={content}
                onChange={(e) => setContent(e.target.value)}
                className="min-h-[100px] resize-none border-none focus-visible:ring-0 p-0 shadow-none"
              />
              {imagePreview && (
                <div className="relative mt-2 rounded-md overflow-hidden">
                  <Image
                    src={imagePreview || "/placeholder.svg"}
                    alt="Post preview"
                    width={500}
                    height={300}
                    className="object-cover max-h-[300px] w-full"
                  />
                  <Button
                    type="button"
                    variant="destructive"
                    size="icon"
                    className="absolute top-2 right-2 h-8 w-8 rounded-full"
                    onClick={removeImage}
                  >
                    <X className="h-4 w-4" />
                  </Button>
                </div>
              )}
            </div>
          </div>
        </CardContent>
        <CardFooter className="flex justify-between p-4 pt-0">
          <div>
            <input type="file" accept="image/*" className="hidden" ref={fileInputRef} onChange={handleImageChange} />
            <Button type="button" variant="outline" size="sm" onClick={() => fileInputRef.current?.click()}>
              <ImageIcon className="h-4 w-4 mr-2" />
              Add Image
            </Button>
          </div>
          <Button type="submit" size="sm" disabled={isSubmitting || (!content.trim() && !image)}>
            {isSubmitting ? (
              <>
                <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                Posting...
              </>
            ) : (
              "Post"
            )}
          </Button>
        </CardFooter>
      </form>
    </Card>
  )
}

