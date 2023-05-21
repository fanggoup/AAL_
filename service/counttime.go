package service

import (
	"AAL_time/serializer"
)

type SelectTime struct {
	StartTime int64 `json:"starttime" form:"starttime"`
	EndTime   int64 `json:"endtime" form:"endtime"`
	ShowAllCategory
}

func Start(res serializer.Response, starttime,endtime int64) serializer.Response {
	var lens int
	var count []serializer.TagTime

	if list, ok := res.Data.(serializer.DataList); ok {
		if list1, ok := list.Item.([]serializer.Category); ok {
			lens = len(list1)
			count = make([]serializer.TagTime, lens) // 初始化 count 切片
			if lens != 0 {
				for i := 0; i < lens; i++ {
					counttime := CountTime(list1[i].ID, starttime,endtime)
					count[i].TagID = list1[i].ID
					count[i].TagName = list1[i].Name
					count[i].CountTime = counttime
				}
			}
		} else {
			return serializer.BuildListResponse(count, uint(lens))
		}
	}
	return serializer.BuildListResponse(count, uint(lens))
}
