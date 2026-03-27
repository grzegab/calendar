package domain

import (
	"errors"
	"time"
)

var (
	ErrSlotAlreadyBooked = errors.New("slot already booked")
	ErrCannotBookPast    = errors.New("cannot book past lesson")
)

type StudentID string
type SlotID string

type Booking struct {
	ID        string
	StudentID StudentID
	SlotID    SlotID
	StartsAt  time.Time
	EndsAt    time.Time
	Status    string
}

func NewBooking(
	id string,
	studentID string,
	slotID string,
	startsAt time.Time,
	slotAlreadyBooked bool,
) (*Booking, error) {

	if startsAt.Before(time.Now()) {
		return nil, ErrCannotBookPast
	}

	//if slotAlreadyBooked {
	//	return nil, ErrSlotAlreadyBooked
	//}

	return &Booking{
		ID:        id,
		StudentID: StudentID(studentID),
		SlotID:    SlotID(slotID),
		StartsAt:  startsAt,
		Status:    "pending",
	}, nil
}

func (b *Booking) Cancel() {
	b.Status = "cancelled"
}
