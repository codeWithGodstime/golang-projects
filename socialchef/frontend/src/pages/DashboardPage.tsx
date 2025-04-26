import type React from "react"
import { Routes, Route } from "react-router-dom"
import DashboardLayout from "../components/DashboardLayout"
import OrdersTable from "../components/OrdersTable"
import OrdersStats from "../components/OrdersStats"

const Dashboard: React.FC = () => {
  return (
    <div className="space-y-6">
      <div className="flex flex-col justify-between gap-4 md:flex-row md:items-center">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
          <p className="text-muted-foreground">Manage your restaurant orders from social media platforms</p>
        </div>
      </div>
      <OrdersStats />
      <OrdersTable />
    </div>
  )
}

const DashboardPage: React.FC = () => {
  return (
    <DashboardLayout>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        {/* Add more routes as needed */}
        <Route path="*" element={<Dashboard />} />
      </Routes>
    </DashboardLayout>
  )
}

export default DashboardPage
