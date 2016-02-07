Random dates /week days.Can be used in testing data transformation/ETL functions.

Installation is simple as

    go get  github.com/kkfaisal/dates

[Go doc](https://godoc.org/github.com/kkfaisal/dates)


#### Features
- All Dates/Week days in range
- Random Dates/Week days in range
- Sorted output is possible



#### Example   1 - Generate a random date in a range
```go
  loc, _ := time.LoadLocation("")
  from := time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
  out, err := dates.RandomDate(from, time.Now(), 1)
  if err == nil {
    fmt.Println(out[0])//A random date in range
  }
```
#### Example   1 - Generate a random weekday
```go
    loc, _ := time.LoadLocation("")
  from := time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
  out, err := dates.RandomWeekDays(from, time.Now(), 1, 1)
  if err == nil {
    fmt.Println(out[0].Weekday()) //Monday
  }

```

#### TODO
- Month first date and last dates in range