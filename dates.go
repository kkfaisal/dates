package dates

import (
	"errors"
	"math/rand"
	"time"
)

//trim off time components after day
func dateFromTime(in time.Time) time.Time {
	return time.Date(in.Year(), in.Month(), in.Day(), 0, 0, 0, 0, utcLocation)
}

func allDays(from, to time.Time) (out []time.Time) {
	newDate := from
	out = append(out, newDate)
	for {
		if newDate = newDate.AddDate(0, 0, 1); newDate.After(to) || newDate.Equal(to) {
			break
		}
		out = append(out, newDate)
	}
	return
}

func randomDays(from, to time.Time, count int) (out []time.Time, err error) {
	max := numOfDaysBetween(from, to)

	if count > max {
		err = errors.New("count should be less than maximum possible value")
		return
	}
	randNums, err := uniqueRandomNumbers(count, max)
	if err != nil {
		return
	}
	out = days(from, randNums)
	return
}

func days(from time.Time, numbers []int) (out []time.Time) {
	for _, v := range numbers {
		out = append(out, from.AddDate(0, 0, v))
	}
	return
}

func numOfDaysBetween(from, to time.Time) (count int) {
	years := yearsBetween(from, to)
	if len(years) == 0 {
		return to.YearDay() - from.YearDay()
	}
	for _, v := range years {
		count += daysInYear(v)
	}
	currentYearDays := daysInYear(from.Year()) - from.YearDay()
	count += currentYearDays + to.YearDay()
	return
}

//From and to years are not included.
func yearsBetween(from, to time.Time) (years []int) {
	gap := to.Year() - from.Year()
	i := 1
	for i < gap {
		years = append(years, from.Year()+1)
		i++
	}
	return
}

func daysInYear(year int) int {
	loc, _ := time.LoadLocation("")
	return time.Date(year, 12, 31, 0, 0, 0, 0, loc).YearDay()
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

//create random numbers
func uniqueRandomNumbers(count, max int) (outSlice []int, err error) {
	if count < 0 {
		err = errors.New("Count Can't be negetive")
		return
	}

	if max < 0 {
		err = errors.New("Max Can't be negetive")
		return

	}

	if count > max {
		err = errors.New("count can't be greater than max")
		return
	}
	outMap := make(map[int]struct{})
	outSlice = []int{}
	for {
		num := rand.Intn(max)
		addUniqueRandValue(outMap, num, max, count)

		if len(outMap) == count {
			break
		}
	}
	for k, _ := range outMap {
		outSlice = append(outSlice, k)
	}
	return
}

//recursively called to avoid duplication
func addUniqueRandValue(inMap map[int]struct{}, num, max, count int) {
	if _, ok := inMap[num]; !ok {
		inMap[num] = struct{}{}
		return
	}
	if num < max && len(inMap) < count {
		//If a number is generated again to that number and try
		addUniqueRandValue(inMap, num+1, max, count)
	}
	return
}
