package main

import (
	"fmt"
	"time"
)

func main() {
	chovva := 2
	wk, err := time.Parse("Jan 2, 2006", "Jan 1, 1998")
	fmt.Println(wk.Weekday(), err)
	if chovva >= int(wk.Weekday()) {
		fmt.Println(chovva - int(wk.Weekday()) + 1)
	} else {
		fmt.Println(7 - int(wk.Weekday()) + chovva + 1)
	}

}

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
