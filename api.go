package dates

import (
	"errors"
	"sort"
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

//All dates from from to to date .Will be in order
func AllDates(from, to time.Time) (out []time.Time, err error) {
	if from, to, err = checkParms(from, to); err != nil {
		return
	}
	out = allDays(from, to)
	return
}

//Random dates between from and to.If count is greater than maximum possible ,error.
func RandomDate(from, to time.Time, count int) (out []time.Time, err error) {
	if from, to, err = checkParms(from, to); err != nil {
		return
	}
	out, err = randomDays(from, to, count)
	return
}

//Same as RandomDate but sorted
func RandomDateSorted(from, to time.Time, count int) (out []time.Time, err error) {

	if out, err = RandomDate(from, to, count); err != nil {
		return
	}
	sort.Sort(DateSlice(out))
	return
}

//Week days in a range(from-to).week day convention is same as in time package(Sunday=0).Out is in order.
func AllWeekDays(from, to time.Time, weekDay time.Weekday) (out []time.Time, err error) {

	if from, to, err = checkParms(from, to); err != nil {
		return
	}
	out = allWeekDays(from, to, weekDay)
	return
}

//Number of random week days in a range.
func RandomWeekDays(from, to time.Time, weekDay time.Weekday, count int) (out []time.Time, err error) {
	if from, to, err = checkParms(from, to); err != nil {
		return
	}
	out, err = randomWeekDays(from, to, weekDay, count)
	return
}

//like RandomWeekDays but sorted.
func RandomWeekDaysSorted(from, to time.Time, weekDay time.Weekday, count int) (out []time.Time, err error) {

	if out, err = RandomWeekDays(from, to, weekDay, count); err != nil {
		return
	}
	sort.Sort(DateSlice(out))
	return
}

// //Random dates from from to To dates
// func GetRandomDatesInRange(count int, from, to Date) []Date {

// }

// //Random Dates from a specific day
// func GetRandomDatesSince(count int, from Date) []Date {

// }

func checkParms(infrom, into time.Time) (from, to time.Time, err error) {
	from = dateFromTime(infrom)
	to = dateFromTime(into)
	if from.After(to) {
		err = errors.New("From date is after to date")
		return
	}
	return

}
