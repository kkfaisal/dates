package dates

import (
	"time"
)

/*
1.AllDays b/w two
2.AllDays since
3.AllWeekday b/w two
4.AllWeekday since
5.Random dates b/w two
6.Random dates Ordered
7.MonthDay b/w two
8.MonthDay since
9.RandomDay
10.RandomDaysince
11.Order dates slice
*/

//Random dates.Range of date is from 0 to now.(see time packege documentation for 0 time in Go)
func GetRandomDates(Count int) []Date {

}

//Random dates from from to To dates
func GetRandomDatesInRange(count int, from, to Date) []Date {

}

//Random Dates from a specific day
func GetRandomDatesSince(count int, from Date) []Date {

}
