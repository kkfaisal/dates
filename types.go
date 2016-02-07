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

type DateSlice []time.Time

func (v DateSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v DateSlice) Len() int {
	return len(v)
}

func (v DateSlice) Less(i, j int) bool {
	return v[i].Before(v[j])
}
