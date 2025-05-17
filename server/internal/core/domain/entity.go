package domain

import (
	"os/user"
	"time"
)

type Id int
type CreatedBy string
type CreatedDT = time.Time
type UpdatedBy string
type UpdatedDT = time.Time

func NewId(id int) Id {
	return Id(id)
}

func NewCreatedDT() CreatedDT {
	return CreatedDT(time.Now().UTC())
}

func NewUpdatedDT() UpdatedDT {
	return UpdatedDT(time.Now().UTC())
}

func NewUpdatedBy() UpdatedBy {
	user, err := user.Current()

	if err != nil {
		return UpdatedBy("defualt - error")
	}

	return UpdatedBy(user.Username)
}

func NewCreatedBy() CreatedBy {
	user, err := user.Current()

	if err != nil {
		return CreatedBy("defualt - error")
	}

	return CreatedBy(user.Username)
}
