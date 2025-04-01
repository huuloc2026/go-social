import type { User, FriendshipStatus } from "@/types"
import { getAuthToken } from "@/lib/auth"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api"

// Helper function to get headers with auth token
const getHeaders = () => {
  return {
    Authorization: `Bearer ${getAuthToken()}`,
    "Content-Type": "application/json",
  }
}

export async function getUserProfile(userId: string): Promise<{ user: User; friendshipStatus: FriendshipStatus }> {
  const response = await fetch(`${API_URL}/users/${userId}`, {
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to fetch user profile")
  }

  return await response.json()
}

export async function getUserFriends(userId: string): Promise<{ friends: User[] }> {
  const response = await fetch(`${API_URL}/users/${userId}/friends`, {
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to fetch user friends")
  }

  return await response.json()
}

export async function sendFriendRequest(userId: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/friends/request`, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify({ userId }),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to send friend request")
  }

  return await response.json()
}

export async function acceptFriendRequest(userId: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/friends/accept`, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify({ userId }),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to accept friend request")
  }

  return await response.json()
}

export async function unfriend(userId: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/friends/${userId}`, {
    method: "DELETE",
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to unfriend user")
  }

  return await response.json()
}

