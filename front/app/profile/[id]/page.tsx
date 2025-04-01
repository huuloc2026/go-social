"use client"

import { useEffect, useState } from "react"
import { useParams } from "next/navigation"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { UserPosts } from "@/components/user-posts"
import { UserFriends } from "@/components/user-friends"
import { getUserProfile, sendFriendRequest, acceptFriendRequest, unfriend } from "@/lib/api/users"
import { useAuth } from "@/hooks/use-auth"
import { useToast } from "@/components/ui/use-toast"
import type { User, FriendshipStatus } from "@/types"
import { Loader2, UserPlus, UserMinus, UserCheck, Edit } from "lucide-react"

export default function ProfilePage() {
  const { id } = useParams()
  const { user: currentUser } = useAuth()
  const { toast } = useToast()
  const [user, setUser] = useState<User | null>(null)
  const [friendshipStatus, setFriendshipStatus] = useState<FriendshipStatus | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const [isActionLoading, setIsActionLoading] = useState(false)

  const isOwnProfile = currentUser?.id === id

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const data = await getUserProfile(id as string)
        setUser(data.user)
        setFriendshipStatus(data.friendshipStatus)
      } catch (error) {
        toast({
          title: "Error",
          description: "Failed to load user profile",
          variant: "destructive",
        })
      } finally {
        setIsLoading(false)
      }
    }

    if (id) {
      fetchProfile()
    }
  }, [id, toast])

  const handleFriendAction = async (action: "send" | "accept" | "unfriend") => {
    if (!user) return

    setIsActionLoading(true)

    try {
      if (action === "send") {
        await sendFriendRequest(user.id)
        setFriendshipStatus("pending_sent")
        toast({
          title: "Friend request sent",
          description: `Friend request sent to ${user.name}`,
        })
      } else if (action === "accept") {
        await acceptFriendRequest(user.id)
        setFriendshipStatus("friends")
        toast({
          title: "Friend request accepted",
          description: `You are now friends with ${user.name}`,
        })
      } else if (action === "unfriend") {
        await unfriend(user.id)
        setFriendshipStatus("none")
        toast({
          title: "Unfriended",
          description: `You are no longer friends with ${user.name}`,
        })
      }
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to perform friend action",
        variant: "destructive",
      })
    } finally {
      setIsActionLoading(false)
    }
  }

  if (isLoading) {
    return (
      <div className="container flex items-center justify-center min-h-[50vh]">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    )
  }

  if (!user) {
    return (
      <div className="container py-8">
        <Card>
          <CardContent className="py-8">
            <div className="text-center">
              <h2 className="text-2xl font-bold">User not found</h2>
              <p className="text-muted-foreground">The user you're looking for doesn't exist or has been removed.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    )
  }

  return (
    <div className="container py-8">
      <Card>
        <CardHeader className="relative pb-0">
          <div className="h-32 bg-gradient-to-r from-primary/20 to-primary/40 rounded-t-lg"></div>
          <div className="flex flex-col md:flex-row gap-4 items-center md:items-end -mt-12 md:-mt-16 px-4 pb-4">
            <Avatar className="h-24 w-24 md:h-32 md:w-32 border-4 border-background">
              <AvatarImage src={user.avatar || `/placeholder.svg?height=128&width=128`} alt={user.name} />
              <AvatarFallback>{user.name.charAt(0)}</AvatarFallback>
            </Avatar>
            <div className="flex-1 text-center md:text-left">
              <h1 className="text-2xl md:text-3xl font-bold">{user.name}</h1>
              <p className="text-muted-foreground">@{user.username}</p>
            </div>
            <div className="flex gap-2">
              {isOwnProfile ? (
                <Button variant="outline" size="sm">
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Profile
                </Button>
              ) : (
                <>
                  {friendshipStatus === "none" && (
                    <Button onClick={() => handleFriendAction("send")} disabled={isActionLoading} size="sm">
                      {isActionLoading ? (
                        <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                      ) : (
                        <UserPlus className="h-4 w-4 mr-2" />
                      )}
                      Add Friend
                    </Button>
                  )}
                  {friendshipStatus === "pending_sent" && (
                    <Button variant="outline" size="sm" disabled>
                      <UserCheck className="h-4 w-4 mr-2" />
                      Request Sent
                    </Button>
                  )}
                  {friendshipStatus === "pending_received" && (
                    <Button onClick={() => handleFriendAction("accept")} disabled={isActionLoading} size="sm">
                      {isActionLoading ? (
                        <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                      ) : (
                        <UserCheck className="h-4 w-4 mr-2" />
                      )}
                      Accept Request
                    </Button>
                  )}
                  {friendshipStatus === "friends" && (
                    <Button
                      variant="outline"
                      onClick={() => handleFriendAction("unfriend")}
                      disabled={isActionLoading}
                      size="sm"
                    >
                      {isActionLoading ? (
                        <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                      ) : (
                        <UserMinus className="h-4 w-4 mr-2" />
                      )}
                      Unfriend
                    </Button>
                  )}
                </>
              )}
            </div>
          </div>
        </CardHeader>
        <CardContent className="pt-6">
          {user.bio && (
            <div className="mb-6">
              <h3 className="font-medium mb-2">About</h3>
              <p className="text-muted-foreground">{user.bio}</p>
            </div>
          )}
          <Separator className="my-4" />
          <Tabs defaultValue="posts">
            <TabsList className="mb-4">
              <TabsTrigger value="posts">Posts</TabsTrigger>
              <TabsTrigger value="friends">Friends</TabsTrigger>
            </TabsList>
            <TabsContent value="posts">
              <UserPosts userId={user.id} />
            </TabsContent>
            <TabsContent value="friends">
              <UserFriends userId={user.id} />
            </TabsContent>
          </Tabs>
        </CardContent>
      </Card>
    </div>
  )
}

