"use client"

import { useState, useEffect } from "react"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Badge } from "@/components/ui/badge"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Search } from "lucide-react"
import { API_URL } from "@/lib/utils"

type Order = {
  id: string
  customer: string
  items: { name: string; quantity: number }[]
  timestamp: string
  platform: "WhatsApp" | "Instagram" | "Facebook" | "Phone"
  status: "New" | "Preparing" | "Done"
}

// Mock data for orders
const mockOrders: Order[] = [
  {
    id: "ORD-001",
    customer: "John Doe",
    items: [
      { name: "Jollof Rice", quantity: 2 },
      { name: "Grilled Chicken", quantity: 1 },
    ],
    timestamp: "2023-04-26T10:30:00",
    platform: "WhatsApp",
    status: "New",
  },
  {
    id: "ORD-002",
    customer: "Jane Smith",
    items: [
      { name: "Fried Rice", quantity: 1 },
      { name: "Turkey", quantity: 2 },
    ],
    timestamp: "2023-04-26T11:15:00",
    platform: "Instagram",
    status: "Preparing",
  },
  {
    id: "ORD-003",
    customer: "Michael Johnson",
    items: [
      { name: "Vegetable Soup", quantity: 1 },
      { name: "Pounded Yam", quantity: 1 },
    ],
    timestamp: "2023-04-26T09:45:00",
    platform: "Facebook",
    status: "Done",
  },
  {
    id: "ORD-004",
    customer: "Sarah Williams",
    items: [
      { name: "Egusi Soup", quantity: 1 },
      { name: "Semovita", quantity: 1 },
    ],
    timestamp: "2023-04-26T12:00:00",
    platform: "WhatsApp",
    status: "New",
  },
  {
    id: "ORD-005",
    customer: "David Brown",
    items: [
      { name: "Pepper Soup", quantity: 2 },
      { name: "White Rice", quantity: 1 },
    ],
    timestamp: "2023-04-26T10:00:00",
    platform: "Phone",
    status: "Preparing",
  },
]

export function OrdersTable() {
  const [orders, setOrders] = useState<Order[]>([])
  const [searchTerm, setSearchTerm] = useState("")
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await fetch(`${API_URL}/orders`)
        const data = await response.json()
        setOrders(data)
      } catch (error) {
        console.error("Failed to fetch orders:", error)
      } finally {
        setLoading(false)
      }
    }

    fetchOrders()
  }, [])

  const handleStatusChange = async (orderId: string, newStatus: string) => {
    try {
      await fetch(`${API_URL}/api/orders/${orderId}/status`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ status: newStatus }),
      })

      setOrders((prevOrders) =>
        prevOrders.map((order) =>
          order.id === orderId ? { ...order, status: newStatus as "New" | "Preparing" | "Done" } : order,
        ),
      )
    } catch (error) {
      console.error("Failed to update order status:", error)
    }
  }

  const filteredOrders = orders.filter(
    (order) =>
      order.customer.toLowerCase().includes(searchTerm.toLowerCase()) ||
      order.id.toLowerCase().includes(searchTerm.toLowerCase()),
  )

  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return new Intl.DateTimeFormat("en-US", {
      month: "short",
      day: "numeric",
      hour: "numeric",
      minute: "numeric",
    }).format(date)
  }

  const getStatusColor = (status: string) => {
    switch (status) {
      case "New":
        return "bg-blue-500"
      case "Preparing":
        return "bg-yellow-500"
      case "Done":
        return "bg-green-500"
      default:
        return "bg-gray-500"
    }
  }

  const getPlatformColor = (platform: string) => {
    switch (platform) {
      case "WhatsApp":
        return "bg-green-100 text-green-800"
      case "Instagram":
        return "bg-purple-100 text-purple-800"
      case "Facebook":
        return "bg-blue-100 text-blue-800"
      case "Phone":
        return "bg-gray-100 text-gray-800"
      default:
        return "bg-gray-100 text-gray-800"
    }
  }

  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle>Recent Orders</CardTitle>
        <div className="relative w-64">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Search orders..."
            className="pl-8"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
      </CardHeader>
      <CardContent>
        {loading ? (
          <div className="flex h-40 items-center justify-center">
            <div className="h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent"></div>
          </div>
        ) : (
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Order ID</TableHead>
                  <TableHead>Customer</TableHead>
                  <TableHead className="hidden md:table-cell">Items</TableHead>
                  <TableHead className="hidden md:table-cell">Time</TableHead>
                  <TableHead>Platform</TableHead>
                  <TableHead>Status</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredOrders.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={6} className="h-24 text-center">
                      No orders found.
                    </TableCell>
                  </TableRow>
                ) : (
                  filteredOrders.map((order) => (
                    <TableRow key={order.id}>
                      <TableCell className="font-medium">{order.id}</TableCell>
                      <TableCell>{order.customer}</TableCell>
                      <TableCell className="hidden md:table-cell">
                        {order.items.map((item, index) => (
                          <div key={index}>
                            {item.quantity}x {item.name}
                          </div>
                        ))}
                      </TableCell>
                      <TableCell className="hidden md:table-cell">{formatDate(order.timestamp)}</TableCell>
                      <TableCell>
                        <Badge variant="outline" className={getPlatformColor(order.platform)}>
                          {order.platform}
                        </Badge>
                      </TableCell>
                      <TableCell>
                        <Select
                          defaultValue={order.status}
                          onValueChange={(value) => handleStatusChange(order.id, value)}
                        >
                          <SelectTrigger className="w-[110px]">
                            <SelectValue />
                          </SelectTrigger>
                          <SelectContent>
                            <SelectItem value="New">
                              <div className="flex items-center gap-2">
                                <div className="h-2 w-2 rounded-full bg-blue-500" aria-hidden="true" />
                                New
                              </div>
                            </SelectItem>
                            <SelectItem value="Preparing">
                              <div className="flex items-center gap-2">
                                <div className="h-2 w-2 rounded-full bg-yellow-500" aria-hidden="true" />
                                Preparing
                              </div>
                            </SelectItem>
                            <SelectItem value="Done">
                              <div className="flex items-center gap-2">
                                <div className="h-2 w-2 rounded-full bg-green-500" aria-hidden="true" />
                                Done
                              </div>
                            </SelectItem>
                          </SelectContent>
                        </Select>
                      </TableCell>
                    </TableRow>
                  ))
                )}
              </TableBody>
            </Table>
          </div>
        )}
      </CardContent>
    </Card>
  )
}
