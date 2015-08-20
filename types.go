package dates

import (
	"sort"
	"time"
)

var (
	utcLocation *time.Location
)

func init() {
	utcLocation, _ = time.LoadLocation("Local")
}

type Date struct {
	Year, Month, Day int
}

type DateSlice []time.Time

func (v DateSlice) Less() {

}

func (v DateSlice) Sort() {
	sort.Sort(v)
}
