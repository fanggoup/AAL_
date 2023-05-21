package serializer

import "AAL_time/modle"

type TimeConsumption struct {
	ID         uint   `json:"id" example:"1"`
	Content    string `json:"content" example:"微信读书 《叫魂》 541页"`
	StartTime  int64  `json:"starttime"`
	EndTime    int64  `json:"endtime"`
	WastedTime bool   `json:"wastedtime"`
	TagID      uint   `json:"tagid" example:"1"`
}

type TagTime struct {
	TagID     uint   `json:"tagid"`
	TagName   string `json:"tagname"`
	CountTime string `json:"counttime"`
}

func BuildTimeConsumption(item modle.TimeConsumption) TimeConsumption {
	return TimeConsumption{
		ID:         item.ID,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		WastedTime: item.WastedTime,
		TagID:      item.TagID,
	}
}

func BuildTimeConsumptions(items []modle.TimeConsumption) (tasks []TimeConsumption) {
	for _, item := range items {
		task := BuildTimeConsumption(item)
		tasks = append(tasks, task)
	}
	return tasks
}
