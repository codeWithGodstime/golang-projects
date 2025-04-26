export interface OrderItem {
  name: string
  quantity: number
}

export interface Order {
  id: string
  customer: string
  items: OrderItem[]
  timestamp: string
  platform: "WhatsApp" | "Instagram" | "Facebook" | "Phone"
  status: "New" | "Preparing" | "Done"
}
