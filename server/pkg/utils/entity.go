package utils

import (
	"os/user"
	"time"
)

type DeletedDT = time.Time
type DeletedBy string

type CreatedBy string
type CreatedDT = time.Time

type UpdatedBy string
type UpdatedDT = time.Time


func NewCreatedDT() CreatedDT {
	return CreatedDT(time.Now().UTC())
}

func NewUpdatedDT() UpdatedDT {
	return UpdatedDT(time.Now().UTC())
}

func NewDeletedDT() UpdatedDT {
	return UpdatedDT(time.Now().UTC())
}

func NewDeletedBy() DeletedBy {
	user, err := user.Current()

	if err != nil {
		return DeletedBy("defualt - error")
	}

	return DeletedBy(user.Username)
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
