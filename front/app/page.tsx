import { redirect } from "next/navigation"
import { Newsfeed } from "@/components/newsfeed"
import { getAuthToken } from "@/lib/auth"

export default function Home() {
  // Check if user is authenticated, redirect to login if not
  const token = getAuthToken()
  if (!token) {
    redirect("/auth/login")
  }

  return (
    <main className="container mx-auto px-4 py-6">
      <Newsfeed />
    </main>
  )
}

