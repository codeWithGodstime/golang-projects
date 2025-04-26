import { UtensilsCrossed } from "lucide-react"

interface LogoProps {
  className?: string
}

export function Logo({ className }: LogoProps) {
  return <UtensilsCrossed className={className} />
}
