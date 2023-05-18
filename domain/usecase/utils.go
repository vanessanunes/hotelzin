package usecase

import (
	"fmt"
	"time"
)

func ConvertStringToDate(DateStart string, DateEnd string) (time.Time, time.Time) {
	layout := "2006-01-02 15:04"
	DateStartTime, err := time.Parse(layout, DateStart)
	if err != nil {
		fmt.Println(err)
	}
	DateEndTime, err := time.Parse(layout, DateEnd)
	if err != nil {
		fmt.Println(err)
	}
	return DateStartTime, DateEndTime
}

func ConvertDBDateToDateTime(DateStart string, DateEnd string) (time.Time, time.Time) {
	layout := "2006-01-02T15:04:05Z"
	DateStartTime, err := time.Parse(layout, DateStart)
	if err != nil {
		fmt.Println(err)
	}
	DateEndTime, err := time.Parse(layout, DateEnd)
	if err != nil {
		fmt.Println(err)
	}
	return DateStartTime, DateEndTime
}

func CountWeekDayAndWeekEnd(DateStart time.Time, DateEnd time.Time) (weekend int, weekday int) {
	diff := DateEnd.Sub(DateStart)
	diff_days := diff.Hours() / 24

	if int(diff_days) == 0 {
		diff_days = 1
	}

	for i := 0; i < int(diff_days); i++ {
		compareDate := DateStart
		compareDate = compareDate.AddDate(0, 0, i)
		fmt.Println(compareDate)
		if compareDate.Weekday() != time.Saturday && compareDate.Weekday() != time.Sunday {
			weekday += 1
		} else {
			weekend += 1
		}
	}
	return weekend, weekday
}
