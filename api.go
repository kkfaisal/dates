package dates

import (
	"time"
)

func GetDateFromTime(time.Time) Date

//Random dates.Range of date is from 0 to now.(see time packege documentation for 0 time in Go)
func GetRandomDates(Count int)

//Random dates from from to To dates
func GetRandomDatesInRange(Count int, From, To Date)
