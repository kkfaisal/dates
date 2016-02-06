package dates

import (
	"time"
)

//trim off time components after day
func dateFromTime(in time.Time) tim.Time {
	return time.Date(in.Year(), int(in.Month()), in.Day(), 0, 0, 0, 0, utcLocation)
}
