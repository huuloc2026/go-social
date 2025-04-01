import { cookies } from "next/headers"
import { NextResponse } from "next/server"

export async function POST(request: Request) {
  try {
    const body = await request.json()
    const { email, password } = body

    // Call your Golang backend API
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    })

    const data = await response.json()
   
    if (!response.ok) {
      return NextResponse.json({ error: data.message || "Login failed" }, { status: response.status })
    }
      
    console.log( data.data.token.access_token)
    // Set the JWT token in an HTTP-only cookie
    const cookieStore = await cookies()
    await cookieStore.set({
      name: "auth_token",
      value: data.data.token.access_token,
      httpOnly: true,
      path: "/",
      secure: process.env.NODE_ENV === "production",
      maxAge: 60 * 60 * 24 * 7, 
    })
    // Return user data without the token (since it's in the cookie)
    return NextResponse.json({ user: data.user })
  } catch (error) {
    console.error("Login error:", error)
    return NextResponse.json({ error: "Internal server error" }, { status: 500 })
  }
}

