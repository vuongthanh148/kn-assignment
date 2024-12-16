package port

import "time"

type MiscRepo interface {
	GetCurrentDate() time.Time
	GetCurrentDateString() string
	NewUUID() string
}
