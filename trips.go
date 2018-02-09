package privatezen

import (
	"time"
)

type TripStatus struct {
	Id    int
	Label string
}

type Trip struct {
	Id          int64
	Name        string
	Arrival     *time.Time
	Departure   *time.Time
	Destination *Whosonfirst // wof_id
	Status      *TripStatus
	Notes       string
}
