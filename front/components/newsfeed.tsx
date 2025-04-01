"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Separator } from "@/components/ui/separator"
import { PostCard } from "@/components/post-card"
import { CreatePostForm } from "@/components/create-post-form"
import { getPosts } from "@/lib/api/posts"
import { useAuth } from "@/hooks/use-auth"
import type { Post } from "@/types"
import { Loader2 } from "lucide-react"

export function Newsfeed() {
  const { user } = useAuth()
  const [posts, setPosts] = useState<Post[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [activeTab, setActiveTab] = useState("all")

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        setIsLoading(true)
        const data = await getPosts(activeTab)
        setPosts(data.posts)
      } catch (error) {
        console.error("Failed to fetch posts:", error)
      } finally {
        setIsLoading(false)
      }
    }

    fetchPosts()
  }, [activeTab])

  const handlePostCreated = (newPost: Post) => {
    setPosts((prevPosts) => [newPost, ...prevPosts])
  }

  const handlePostUpdated = (updatedPost: Post) => {
    setPosts((prevPosts) => prevPosts.map((post) => (post.id === updatedPost.id ? updatedPost : post)))
  }

  const handlePostDeleted = (postId: string) => {
    setPosts((prevPosts) => prevPosts.filter((post) => post.id !== postId))
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div className="md:col-span-2 space-y-6">
        {user && <CreatePostForm onPostCreated={handlePostCreated} />}

        <Tabs defaultValue="all" className="w-full" onValueChange={setActiveTab}>
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="all">All Posts</TabsTrigger>
            <TabsTrigger value="friends">Friends</TabsTrigger>
          </TabsList>
          <TabsContent value="all" className="space-y-4 mt-4">
            {isLoading ? (
              <div className="flex justify-center py-8">
                <Loader2 className="h-8 w-8 animate-spin text-primary" />
              </div>
            ) : posts.length > 0 ? (
              posts.map((post) => (
                <PostCard
                  key={post.id}
                  post={post}
                  onPostUpdated={handlePostUpdated}
                  onPostDeleted={handlePostDeleted}
                />
              ))
            ) : (
              <div className="text-center py-8">
                <p className="text-muted-foreground">No posts to display</p>
                <Button variant="outline" className="mt-4">
                  Find friends to follow
                </Button>
              </div>
            )}
          </TabsContent>
          <TabsContent value="friends" className="space-y-4 mt-4">
            {isLoading ? (
              <div className="flex justify-center py-8">
                <Loader2 className="h-8 w-8 animate-spin text-primary" />
              </div>
            ) : posts.length > 0 ? (
              posts.map((post) => (
                <PostCard
                  key={post.id}
                  post={post}
                  onPostUpdated={handlePostUpdated}
                  onPostDeleted={handlePostDeleted}
                />
              ))
            ) : (
              <div className="text-center py-8">
                <p className="text-muted-foreground">No posts from friends yet</p>
                <Button variant="outline" className="mt-4">
                  Find friends to follow
                </Button>
              </div>
            )}
          </TabsContent>
        </Tabs>
      </div>

      <div className="hidden md:block space-y-6">
        <div className="rounded-lg border bg-card text-card-foreground shadow-sm">
          <div className="p-4 flex flex-row items-center justify-between space-y-0 pb-2">
            <h3 className="font-semibold leading-none tracking-tight">Friend Suggestions</h3>
          </div>
          <Separator />
          <div className="p-4">
            <div className="space-y-4">
              {/* Friend suggestions would go here */}
              <p className="text-sm text-muted-foreground">Find people you may know based on your connections.</p>
              <Button variant="outline" className="w-full">
                Find Friends
              </Button>
            </div>
          </div>
        </div>

        <div className="rounded-lg border bg-card text-card-foreground shadow-sm">
          <div className="p-4 flex flex-row items-center justify-between space-y-0 pb-2">
            <h3 className="font-semibold leading-none tracking-tight">Trending Topics</h3>
          </div>
          <Separator />
          <div className="p-4">
            <div className="space-y-4">
              {/* Trending topics would go here */}
              <p className="text-sm text-muted-foreground">No trending topics at the moment.</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

