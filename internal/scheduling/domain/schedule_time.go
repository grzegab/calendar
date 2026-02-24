package domain

import (
	"errors"
	"time"
)

type ScheduleTime struct {
	start time.Time
	end   time.Time
}

func NewScheduleTime(ts time.Time, te time.Time) (ScheduleTime, error) {
	if ts.After(te) {
		return ScheduleTime{}, errors.New("start time must be before end time")
	}

	return ScheduleTime{ts, te}, nil
}
