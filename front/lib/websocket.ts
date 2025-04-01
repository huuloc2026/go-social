import type { Notification } from "@/types"
import { getAuthToken } from "@/lib/auth"

const WS_URL = process.env.NEXT_PUBLIC_WS_URL || "ws://localhost:8080/ws"

let socket: WebSocket | null = null
let reconnectInterval: NodeJS.Timeout | null = null
let reconnectAttempts = 0
const MAX_RECONNECT_ATTEMPTS = 5

export function setupWebSocket(onNotification: (notification: Notification) => void): void {
  const token = getAuthToken()
  if (!token) return

  try {
    socket = new WebSocket(`${WS_URL}?token=${token}`)

    socket.onopen = () => {
      console.log("WebSocket connection established")
      reconnectAttempts = 0
      if (reconnectInterval) {
        clearInterval(reconnectInterval)
        reconnectInterval = null
      }
    }

    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        if (data.type === "notification") {
          onNotification(data.notification)
        }
      } catch (error) {
        console.error("Error parsing WebSocket message:", error)
      }
    }

    socket.onclose = (event) => {
      console.log("WebSocket connection closed:", event.code, event.reason)

      // Only attempt to reconnect if it wasn't a clean close
      if (!event.wasClean && reconnectAttempts < MAX_RECONNECT_ATTEMPTS) {
        reconnectInterval = setInterval(() => {
          reconnectAttempts++
          console.log(`Attempting to reconnect (${reconnectAttempts}/${MAX_RECONNECT_ATTEMPTS})...`)
          setupWebSocket(onNotification)

          if (reconnectAttempts >= MAX_RECONNECT_ATTEMPTS) {
            if (reconnectInterval) {
              clearInterval(reconnectInterval)
              reconnectInterval = null
            }
          }
        }, 5000)
      }
    }

    socket.onerror = (error) => {
      console.error("WebSocket error:", error)
    }
  } catch (error) {
    console.error("Failed to setup WebSocket:", error)
  }
}

export function closeWebSocket(): void {
  if (socket) {
    socket.close()
    socket = null
  }

  if (reconnectInterval) {
    clearInterval(reconnectInterval)
    reconnectInterval = null
  }
}

export function sendWebSocketMessage(message: any): void {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(message))
  } else {
    console.error("WebSocket is not connected")
  }
}

