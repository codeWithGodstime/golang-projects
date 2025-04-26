import type React from "react"
import DashboardSidebar from "./DashboardSidebar"

interface DashboardLayoutProps {
  children: React.ReactNode
}

const DashboardLayout: React.FC<DashboardLayoutProps> = ({ children }) => {
  return (
    <div className="flex h-screen flex-col md:flex-row">
      <DashboardSidebar />
      <main className="flex-1 overflow-y-auto p-4 md:p-6">{children}</main>
    </div>
  )
}

export default DashboardLayout
