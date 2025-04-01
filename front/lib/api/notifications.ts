import type { Notification } from "@/types"
import { getAuthToken } from "@/lib/auth"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api"

// Helper function to get headers with auth token
const getHeaders = () => {
  return {
    Authorization: `Bearer ${getAuthToken()}`,
    "Content-Type": "application/json",
  }
}

export async function getNotifications(): Promise<{ notifications: Notification[] }> {
  const response = await fetch(`${API_URL}/notifications`, {
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to fetch notifications")
  }

  return await response.json()
}

export async function markNotificationAsRead(notificationId: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/notifications/${notificationId}/read`, {
    method: "POST",
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to mark notification as read")
  }

  return await response.json()
}

export async function markAllNotificationsAsRead(): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/notifications/read-all`, {
    method: "POST",
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to mark all notifications as read")
  }

  return await response.json()
}

