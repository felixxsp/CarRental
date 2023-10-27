package Entity

import "time"

type CarMeter struct {
	PlateID    int
	Year       int
	Engine     string
	KMs        int
	LastReturn time.Time
	NextReturn time.Time
}
