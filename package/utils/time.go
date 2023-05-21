package utils

import "time"

func WeekStart() (time.Time,time.Time) {
	// 获取当前时间
	now := time.Now()

	// 获取本周一的日期
	weekStart := now.AddDate(0, 0, -int(now.Weekday())+1)

	// 设置起始时间为当天的 00:00:00
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())

	// fmt.Println("本周起始时间:", weekStart)
	return weekStart,now
}

func MonthStart() (time.Time,time.Time){
	// 获取当前时间
	now := time.Now()

	// 获取当前月份的第一天
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return firstDay,now
}

func YearStart()(time.Time,time.Time) {
	// 获取当前时间
	now := time.Now()

	// 获取当前月份的第一天
	firstDay := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())

	return firstDay,now
}