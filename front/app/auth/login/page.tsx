import { LoginForm } from "@/components/auth/login-form"
import { Header } from "@/components/header"

export default function LoginPage() {
  return (
    <main className="min-h-screen bg-background">
      <Header />
      <div className="container flex items-center justify-center py-12">
        <div className="w-full max-w-md">
          <LoginForm />
        </div>
      </div>
    </main>
  )
}

