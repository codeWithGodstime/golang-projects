import type React from "react"
import { UtensilsCrossed } from "lucide-react"

interface LogoProps {
  className?: string
}

const Logo: React.FC<LogoProps> = ({ className }) => {
  return <UtensilsCrossed className={className} />
}

export default Logo
