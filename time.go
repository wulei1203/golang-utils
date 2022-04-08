package utils

import (
	"strconv"
	"time"
)

func TimeToDateString(t *time.Time) string {
	return t.Format("2006-01-02")
}

func DateStringtoTime(timeStr string) *time.Time {
	value, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return &value
}

func ExcelDateToDate(excelDate string) *time.Time {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(excelDate)
	value := excelTime.Add(time.Second * time.Duration(days*86400))
	return &value
}

func GetNowTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowTime() *time.Time {
	now := time.Now()
	return &now
}
