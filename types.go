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

type Date interface {
	Year() int
	Month() time.Month
	Day() int
}

type DateSlice []Date

func (v DateSlice) Less() {

}

func (v DateSlice) Sort() {
	sort.Sort(v)
}
