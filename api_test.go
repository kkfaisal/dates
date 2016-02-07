package dates

import (
	"testing"
	"time"
)

func TestDateFromTime(t *testing.T) {
	tm := time.Now()
	tm = dateFromTime(tm)
	if tm.Hour() != 0 || tm.Minute() != 0 || tm.Nanosecond() != 0 || tm.Second() != 0 {
		t.Fail()
	}
	t.Log("DateFromTime is passed")
}

func TestRandomDate(t *testing.T) {
	loc, _ := time.LoadLocation("")
	from := time.Date(2000, 1, 1, 10, 2, 33, 22, loc)
	to := time.Date(2000, 1, 10, 10, 2, 33, 22, loc)

	_, err := RandomDate(from, to, 10)

	if err == nil {
		t.Fail()
	}

	out, err := RandomDate(from, to, 8)

	if err != nil {
		t.Fail()
	}

	if len(out) != 8 {
		t.Fail()
	}
}

func TestRandomWeekDay(t *testing.T) {
	loc, _ := time.LoadLocation("")
	from := time.Date(2000, 1, 1, 10, 2, 33, 22, loc)
	to := time.Date(2000, 1, 30, 10, 2, 33, 22, loc)

	_, err := RandomWeekDays(from, to, 1, 5)
	if err == nil {
		t.Fail()
	}
	out, err := RandomWeekDays(from, to, 1, 4)

	if err != nil {
		t.Fail()
	}

	if len(out) != 4 {
		t.Fail()
	}

	for _, v := range out {
		if v.Weekday() != 1 {
			t.Fail()
		}
	}
}
