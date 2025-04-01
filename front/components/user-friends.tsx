"use client"

import { useState, useEffect } from "react"
import Link from "next/link"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { getUserFriends } from "@/lib/api/users"
import { useToast } from "@/components/ui/use-toast"
import type { User } from "@/types"
import { Loader2, UserPlus } from "lucide-react"

interface UserFriendsProps {
  userId: string
}

export function UserFriends({ userId }: UserFriendsProps) {
  const { toast } = useToast()
  const [friends, setFriends] = useState<User[]>([])
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const fetchFriends = async () => {
      try {
        setIsLoading(true)
        const data = await getUserFriends(userId)
        setFriends(data.friends)
      } catch (error) {
        toast({
          title: "Error",
          description: "Failed to load user friends",
          variant: "destructive",
        })
      } finally {
        setIsLoading(false)
      }
    }

    fetchFriends()
  }, [userId, toast])

  if (isLoading) {
    return (
      <div className="flex justify-center py-8">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    )
  }

  if (friends.length === 0) {
    return (
      <div className="text-center py-8">
        <p className="text-muted-foreground">No friends to display</p>
        <Button variant="outline" className="mt-4">
          <UserPlus className="h-4 w-4 mr-2" />
          Find Friends
        </Button>
      </div>
    )
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
      {friends.map((friend) => (
        <Card key={friend.id}>
          <CardContent className="p-4">
            <Link href={`/profile/${friend.id}`} className="flex items-center gap-3">
              <Avatar>
                <AvatarImage src={friend.avatar || `/placeholder.svg?height=40&width=40`} alt={friend.name} />
                <AvatarFallback>{friend.name.charAt(0)}</AvatarFallback>
              </Avatar>
              <div>
                <div className="font-medium">{friend.name}</div>
                <div className="text-sm text-muted-foreground">@{friend.username}</div>
              </div>
            </Link>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}

