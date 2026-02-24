package domain

import (
	"time"
)

type SchedulingEvent interface {
	EventName() string
	OccurredAt() time.Time
}
