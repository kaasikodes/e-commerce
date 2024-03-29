package types

type NotificationService interface {
	SendNotification(email string, message string) error
}