import { SignupForm } from "@/components/signup-form"
import { Logo } from "@/components/logo"
import Link from "next/link"

export default function SignupPage() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center p-4">
      <div className="w-full max-w-md space-y-8">
        <div className="flex flex-col items-center space-y-2">
          <Logo className="h-12 w-12" />
          <h1 className="text-3xl font-bold">socialChef</h1>
          <p className="text-muted-foreground">Restaurant Order Management</p>
        </div>
        <div className="rounded-lg border bg-card p-8 shadow-sm">
          <div className="mb-6 space-y-2 text-center">
            <h2 className="text-2xl font-semibold">Create an account</h2>
            <p className="text-sm text-muted-foreground">Enter your information to create an account</p>
          </div>
          <SignupForm />
          <div className="mt-6 text-center text-sm">
            Already have an account?{" "}
            <Link href="/login" className="font-medium text-primary underline">
              Login
            </Link>
          </div>
        </div>
      </div>
    </div>
  )
}
