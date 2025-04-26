import { OrdersTable } from "@/components/orders-table"
import { OrdersStats } from "@/components/orders-stats"

export default function DashboardPage() {
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
