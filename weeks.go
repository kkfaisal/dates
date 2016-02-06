package dates

import (
	"errors"
	"time"
)

func numberOfWeekDaysBetween(from, to time.Time, weekDay time.Weekday) (count int) {
	return numOfDaysBetween(nearWeekDay(from, weekDay), nearWeekDay(to, weekDay)) / 7

}

func nearWeekDay(day time.Time, weekDay time.Weekday) (out time.Time) {
	out = day //default case where day itself is  weekday
	fromWeekDay := day.Weekday()
	switch {
	case fromWeekDay < weekDay:
		out = day.AddDate(0, 0, int(weekDay-fromWeekDay))
	case fromWeekDay > weekDay:
		out = day.AddDate(0, 0, int(7-fromWeekDay+weekDay))
	}
	return
}

func randomWeekDays(from, to time.Time, weekDay time.Weekday, count int) (out []time.Time, err error) {
	max := numberOfWeekDaysBetween(from, to, weekDay)

	if count > max {
		err = errors.New("count should be less than maximum possible value")
		return
	}
	randNums, err := uniqueRandomNumbers(count, max)
	if err != nil {
		return
	}
	from = nearWeekDay(from, weekDay)
	out = weekDays(from, randNums)
	return

}
func weekDays(from time.Time, numbers []int) (out []time.Time) {
	for _, v := range numbers {
		out = append(out, from.AddDate(0, 0, v*7))
	}
	return
}
func allWeekDays(from, to time.Time, weeKday time.Weekday) (out []time.Time) {
	newDate := nearWeekDay(from, weeKday)
	out = append(out, newDate)
	for {
		if newDate = newDate.AddDate(0, 0, 7); newDate.After(to) || newDate.Equal(to) {
			break
		}
		out = append(out, newDate)
	}
	return
}
