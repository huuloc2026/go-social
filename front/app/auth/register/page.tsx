import { RegisterForm } from "@/components/auth/register-form"
import { Header } from "@/components/header"

export default function RegisterPage() {
  return (
    <main className="min-h-screen bg-background">
      <Header />
      <div className="container flex items-center justify-center py-12">
        <div className="w-full max-w-md">
          <RegisterForm />
        </div>
      </div>
    </main>
  )
}

