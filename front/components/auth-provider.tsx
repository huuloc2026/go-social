"use client"

import type React from "react"

import { createContext, useEffect, useState } from "react"
import { useRouter, usePathname } from "next/navigation"
import type { User } from "@/types"
import { getCurrentUser, logout as apiLogout } from "@/lib/api/auth"
import { getAuthToken, removeAuthToken } from "@/lib/auth"

interface AuthContextType {
  user: User | null
  setUser: (user: User | null) => void
  logout: () => Promise<void>
  isLoading: boolean
}

export const AuthContext = createContext<AuthContextType>({
  user: null,
  setUser: () => {},
  logout: async () => {},
  isLoading: true,
})

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const router = useRouter()
  const pathname = usePathname()

  useEffect(() => {
    const initAuth = async () => {
      const token = getAuthToken()

      if (!token) {
        setIsLoading(false)
        return
      }

      try {
        const userData = await getCurrentUser()
        setUser(userData)
      } catch (error) {
        // Token might be invalid, remove it
        removeAuthToken()
      } finally {
        setIsLoading(false)
      }
    }

    initAuth()
  }, [])

  const logout = async () => {
    try {
      await apiLogout()
    } catch (error) {
      console.error("Logout error:", error)
    } finally {
      setUser(null)
      removeAuthToken()
      router.push("/auth/login")
    }
  }

  // Redirect to login if not authenticated and trying to access protected routes
  useEffect(() => {
    if (!isLoading && !user && !pathname.startsWith("/auth/") && pathname !== "/") {
      router.push("/auth/login")
    }
  }, [user, isLoading, pathname, router])

  return <AuthContext.Provider value={{ user, setUser, logout, isLoading }}>{children}</AuthContext.Provider>
}

