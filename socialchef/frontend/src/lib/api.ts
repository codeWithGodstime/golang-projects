import type { Order } from "../types"

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

// Mock function to get orders
export const getOrders = async (): Promise<Order[]> => {
  // Simulate API call
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(mockOrders)
    }, 1000)
  })
}

// Mock function to update order status
export const updateOrderStatus = async (orderId: string, status: string): Promise<{ success: boolean }> => {
  // Simulate API call
  return new Promise((resolve) => {
    setTimeout(() => {
      // In a real app, you would call your API
      // await fetch(`/api/orders/${orderId}/status`, {
      //   method: "PUT",
      //   headers: { "Content-Type": "application/json" },
      //   body: JSON.stringify({ status }),
      // });

      // Update the mock data
      const orderIndex = mockOrders.findIndex((order) => order.id === orderId)
      if (orderIndex !== -1) {
        mockOrders[orderIndex].status = status as "New" | "Preparing" | "Done"
      }

      resolve({ success: true })
    }, 500)
  })
}

// Mock login function
export const loginUser = async (credentials: { email: string; password: string }) => {
  // Simulate API call
  return new Promise((resolve) => {
    setTimeout(() => {
      // For demo purposes, accept any credentials
      resolve({
        success: true,
        user: {
          name: "Restaurant Staff",
          email: credentials.email,
        },
      })
    }, 1000)
  })
}

// Mock signup function
export const signupUser = async (userData: { name: string; email: string; password: string }) => {
  // Simulate API call
  return new Promise((resolve) => {
    setTimeout(() => {
      // For demo purposes, accept any credentials
      resolve({
        success: true,
        user: {
          name: userData.name,
          email: userData.email,
        },
      })
    }, 1000)
  })
}
