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
