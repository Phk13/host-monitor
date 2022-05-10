package domain

import "time"

type Notification interface {
	SendNotification(hostIP string, timestamp time.Time)
}
