package new_timeslot

import "time"

type Command struct {
	TeacherID string
	StartTime time.Time
	EndTime   time.Time
}
