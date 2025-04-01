"use client"

import type React from "react"

import { useState } from "react"
import Link from "next/link"
import { formatDistanceToNow } from "date-fns"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Separator } from "@/components/ui/separator"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { useNotifications } from "@/hooks/use-notifications"
import { markAllNotificationsAsRead } from "@/lib/api/notifications"
import type { Notification } from "@/types"

interface NotificationDropdownProps {
  children: React.ReactNode
}

export function NotificationDropdown({ children }: NotificationDropdownProps) {
  const { notifications, markAsRead, clearUnread } = useNotifications()
  const [open, setOpen] = useState(false)

  const handleOpenChange = (isOpen: boolean) => {
    setOpen(isOpen)
    if (isOpen) {
      // Mark all as read when opening the dropdown
      markAllNotificationsAsRead().then(() => {
        clearUnread()
      })
    }
  }

  const getNotificationContent = (notification: Notification) => {
    switch (notification.type) {
      case "friend_request":
        return {
          text: "sent you a friend request",
          link: `/profile/${notification.sender.id}`,
        }
      case "friend_accept":
        return {
          text: "accepted your friend request",
          link: `/profile/${notification.sender.id}`,
        }
      case "post_like":
        return {
          text: "liked your post",
          link: `/post/${notification.postId}`,
        }
      case "post_comment":
        return {
          text: "commented on your post",
          link: `/post/${notification.postId}`,
        }
      default:
        return {
          text: "sent you a notification",
          link: "/",
        }
    }
  }

  return (
    <Popover open={open} onOpenChange={handleOpenChange}>
      <PopoverTrigger asChild>{children}</PopoverTrigger>
      <PopoverContent className="w-80 p-0" align="end">
        <div className="flex items-center justify-between p-4">
          <h4 className="font-semibold">Notifications</h4>
          <Button variant="ghost" size="sm" className="text-xs">
            Mark all as read
          </Button>
        </div>
        <Separator />
        <ScrollArea className="h-[300px]">
          {notifications.length > 0 ? (
            <div>
              {notifications.map((notification) => {
                const content = getNotificationContent(notification)
                return (
                  <div
                    key={notification.id}
                    className={`flex items-start gap-3 p-4 hover:bg-muted transition-colors ${
                      !notification.read ? "bg-muted/50" : ""
                    }`}
                    onClick={() => markAsRead(notification.id)}
                  >
                    <Avatar className="h-9 w-9">
                      <AvatarImage
                        src={notification.sender.avatar || `/placeholder.svg?height=36&width=36`}
                        alt={notification.sender.name}
                      />
                      <AvatarFallback>{notification.sender.name.charAt(0)}</AvatarFallback>
                    </Avatar>
                    <div className="flex-1 space-y-1">
                      <p className="text-sm">
                        <Link href={`/profile/${notification.sender.id}`} className="font-medium hover:underline">
                          {notification.sender.name}
                        </Link>{" "}
                        <Link href={content.link} className="text-muted-foreground">
                          {content.text}
                        </Link>
                      </p>
                      <p className="text-xs text-muted-foreground">
                        {formatDistanceToNow(new Date(notification.createdAt), {
                          addSuffix: true,
                        })}
                      </p>
                    </div>
                    {!notification.read && <div className="h-2 w-2 rounded-full bg-primary"></div>}
                  </div>
                )
              })}
            </div>
          ) : (
            <div className="flex items-center justify-center h-full">
              <p className="text-sm text-muted-foreground">No notifications yet</p>
            </div>
          )}
        </ScrollArea>
      </PopoverContent>
    </Popover>
  )
}

