import { Feed } from "@/components/feed"
import { Header } from "@/components/header"

export default function Home() {
  return (
    <main className="min-h-screen bg-background">
      <Header />
      <div className="container max-w-4xl py-8">
        <Feed />
      </div>
    </main>
  )
}

