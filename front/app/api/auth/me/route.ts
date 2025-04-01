import { cookies } from "next/headers"
import { NextResponse } from "next/server"

export async function GET() {
  try {
    const cookieStore = await cookies()
    const token = cookieStore.get("auth_token")?.value

    if (!token) {
      return NextResponse.json({ error: "Not authenticated" }, { status: 401 })
    }

    // Call your Golang backend API to validate the token and get user data
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/me`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    if (!response.ok) {
      // If token is invalid, clear it
      if (response.status === 401) {
        await cookieStore.delete("auth_token")
      }

      return NextResponse.json({ error: "Authentication failed" }, { status: response.status })
    }

    const data = await response.json()
    return NextResponse.json({ user: data.user })
  } catch (error) {
    console.error("Auth verification error:", error)
    return NextResponse.json({ error: "Internal server error" }, { status: 500 })
  }
}

