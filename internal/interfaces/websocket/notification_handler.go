// package websocket

// import (
// 	"encoding/json"
// 	"log"

// 	"github.com/gofiber/websocket/v2"
// 	"github.com/huuloc2026/go-social/internal/domain"
// )

// type NotificationHandler struct {
// 	hub                 *Hub
// 	notificationService domain.NotificationService
// }

// func NewNotificationHandler(hub *Hub, notificationService domain.NotificationService) *NotificationHandler {
// 	return &NotificationHandler{
// 		hub:                 hub,
// 		notificationService: notificationService,
// 	}
// }

// func (h *NotificationHandler) HandleConnection(c *websocket.Conn) {
// 	userID := c.Locals("userID").(uint)
// 	client := NewClient(userID, c, h.hub)
// 	h.hub.register <- client

// 	defer func() {
// 		h.hub.unregister <- client
// 		c.Close()
// 	}()

// 	for {
// 		_, message, err := c.ReadMessage()
// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("WebSocket error: %v", err)
// 			}
// 			break
// 		}

// 		var msg struct {
// 			Type string          `json:"type"`
// 			Data json.RawMessage `json:"data"`
// 		}

// 		if err := json.Unmarshal(message, &msg); err != nil {
// 			log.Printf("Error parsing message: %v", err)
// 			continue
// 		}

// 		switch msg.Type {
// 		case "mark_as_read":
// 			var data struct {
// 				NotificationID uint `json:"notification_id"`
// 			}
// 			if err := json.Unmarshal(msg.Data, &data); err != nil {
// 				log.Printf("Error parsing mark_as_read data: %v", err)
// 				continue
// 			}
// 			_ = h.notificationService.MarkAsRead(userID, data.NotificationID)
// 		}
// 	}
// }
