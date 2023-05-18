package usecase

const weekdayParkValue float32 = 15.00
const weekendParkValue float32 = 20.00

func ParkingTotalValue(DateStart string, DateEnd string) float32 {
	DateStartTime, DateEndTime := ConvertDBDateToDateTime(DateStart, DateEnd)
	totalWeekEnd, totalWeekDay := CountWeekDayAndWeekEnd(DateStartTime, DateEndTime)
	return CalculateParkingTotalValue(totalWeekDay, totalWeekEnd)
}

func CalculateParkingTotalValue(weekday int, weekend int) float32 {
	return (float32(weekday) * weekdayParkValue) + (float32(weekend) * weekendParkValue)
}

func CalculateExtraParkingHour(extraHour int32) float32 {
	if extraHour == 1 {
		return weekdayParkValue
	}
	return 0
}
