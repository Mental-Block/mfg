package entity

import (
	"time"
)

type Id = int
type CreatedBy = string
type CreatedDT = time.Time
type UpdatedBy = *string
type UpdatedDT = *time.Time

func NewCreatedDT() CreatedDT {
	return CreatedDT(time.Now().UTC())
}

func NewUpdatedDT() UpdatedDT {
	utc := time.Now().UTC()
	return UpdatedDT(&utc)
}

func NewUpdatedBy(username string) UpdatedBy {
	return UpdatedBy(&username)
}

func NewCreatedBy(username string) CreatedBy {
	return CreatedBy(username)
}
