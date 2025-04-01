import type { User } from "@/types"
import { setAuthToken, removeAuthToken } from "@/lib/auth"

const API_URL = process.env.NEXT_PUBLIC_API_URL 

export async function login(email: string, password: string): Promise<{ user: User; token: string }> {
  const response = await fetch(`${API_URL}/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to login")
  }

  const data = await response.json()
  setAuthToken(data.token)

  return data
}

export async function register(name: string, email: string, password: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/auth/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ name, email, password }),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to register")
  }

  return await response.json()
}

export async function logout(): Promise<void> {
  const token = localStorage.getItem("auth_token")

  if (!token) return

  try {
    const response = await fetch(`${API_URL}/auth/logout`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    if (!response.ok) {
      console.error("Logout failed on server")
    }
  } catch (error) {
    console.error("Logout error:", error)
  } finally {
    removeAuthToken()
  }
}

export async function getCurrentUser(): Promise<User> {
  const token = localStorage.getItem("auth_token")

  if (!token) {
    throw new Error("No authentication token found")
  }

  const response = await fetch(`${API_URL}/auth/me`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to get current user")
  }

  return (await response.json()).user
}

