package usecase

const weekdayRoomValue float32 = 120.00
const weekendRoomValue float32 = 150.00

func RoomTotalValue(DateStart string, DateEnd string) float32 {
	DateStartTime, DateEndTime := ConvertDBDateToDateTime(DateStart, DateEnd)
	totalWeekEnd, totalWeekDay := CountWeekDayAndWeekEnd(DateStartTime, DateEndTime)
	return CalculateRoomTotalValue(totalWeekDay, totalWeekEnd)
}

func CalculateRoomTotalValue(weekday int, weekend int) float32 {
	return float32(weekday)*weekdayRoomValue + float32(weekend)*weekendRoomValue
}

func CalculateExtraRoomHour(extraHour bool) float32 {
	if extraHour == true {
		return weekdayRoomValue
	}
	return 0
}
