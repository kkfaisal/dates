package dates

import (
	"time"
)

var (
	utcLocation *time.Location
)

func init() {
	utcLocation, _ = time.LoadLocation("Local")
}
