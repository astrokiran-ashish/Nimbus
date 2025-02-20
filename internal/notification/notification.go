package notification

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type INotification interface {
	SendNotificationToConsultant(consultantID uuid.UUID, message string) error
	SendNotificationToUser(userID uuid.UUID, message string) error
}

type Notification struct {
	db                    *database.Database
	logger                *zap.Logger
	userConnections       map[uuid.UUID]http.ResponseWriter
	consultantConnections map[uuid.UUID]http.ResponseWriter
	mu                    sync.Mutex // Mutex to protect access to the maps
}

func NewNotification(db *database.Database, logger *zap.Logger) *Notification {
	return &Notification{
		db:                    db,
		logger:                logger,
		userConnections:       make(map[uuid.UUID]http.ResponseWriter),
		consultantConnections: make(map[uuid.UUID]http.ResponseWriter),
	}
}

func (n *Notification) SendNotificationToConsultant(consultantID uuid.UUID, message string) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if conn, ok := n.consultantConnections[consultantID]; ok {
		fmt.Fprintf(conn, "data: %s\n\n", message)
		n.logger.Info("Sent notification to consultant", zap.String("message", message))
		return nil
	}
	n.logger.Warn("No connection found for consultant", zap.String("consultantID", consultantID.String()))
	return fmt.Errorf("no connection found for consultant %s", consultantID)
}

func (n *Notification) SendNotificationToUser(userID uuid.UUID, message string) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if conn, ok := n.userConnections[userID]; ok {
		fmt.Fprintf(conn, "data: %s\n\n", message)
		n.logger.Info("Sent notification to user", zap.String("message", message))
		return nil
	}
	n.logger.Warn("No connection found for user", zap.String("userID", userID.String()))
	return fmt.Errorf("no connection found for user %s", userID)
}
