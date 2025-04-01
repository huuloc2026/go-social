"use client"

import type React from "react"

import { createContext, useEffect, useState } from "react"
import { getNotifications } from "@/lib/api/notifications"
import { useAuth } from "@/hooks/use-auth"
import { setupWebSocket, closeWebSocket } from "@/lib/websocket"
import type { Notification } from "@/types"

interface NotificationContextType {
  notifications: Notification[]
  unreadCount: number
  markAsRead: (id: string) => void
  clearUnread: () => void
}

export const NotificationContext = createContext<NotificationContextType>({
  notifications: [],
  unreadCount: 0,
  markAsRead: () => {},
  clearUnread: () => {},
})

export function NotificationProvider({ children }: { children: React.ReactNode }) {
  const { user, isLoading } = useAuth()
  const [notifications, setNotifications] = useState<Notification[]>([])
  const [unreadCount, setUnreadCount] = useState(0)

  useEffect(() => {
    if (!isLoading && user) {
      // Fetch initial notifications
      const fetchNotifications = async () => {
        try {
          const data = await getNotifications()
          setNotifications(data.notifications)
          setUnreadCount(data.notifications.filter((n) => !n.read).length)
        } catch (error) {
          console.error("Failed to fetch notifications:", error)
        }
      }

      fetchNotifications()

      // Setup WebSocket connection for real-time notifications
      const handleNewNotification = (notification: Notification) => {
        setNotifications((prev) => [notification, ...prev])
        setUnreadCount((prev) => prev + 1)
      }

      setupWebSocket(handleNewNotification)

      return () => {
        closeWebSocket()
      }
    }
  }, [user, isLoading])

  const markAsRead = (id: string) => {
    setNotifications((prev) =>
      prev.map((notification) => (notification.id === id ? { ...notification, read: true } : notification)),
    )
    setUnreadCount((prev) => Math.max(0, prev - 1))
  }

  const clearUnread = () => {
    setNotifications((prev) => prev.map((notification) => ({ ...notification, read: true })))
    setUnreadCount(0)
  }

  return (
    <NotificationContext.Provider value={{ notifications, unreadCount, markAsRead, clearUnread }}>
      {children}
    </NotificationContext.Provider>
  )
}

