package service

import (
	"AAL_time/modle"
	"AAL_time/package/e"
	"AAL_time/package/utils"
	"AAL_time/serializer"
	"fmt"
	"time"
)

type CreateTimeConsumption struct {
	StartTime  int64  `json:"starttime" form:"starttime"`
	Content    string `json:"content" form:"content"`
	WastedTime bool   `json:"wastedtime" form:"wastedtime"`
	TagID      uint   `json:"tagid" form:"tagid"`
}

type ShowTimeConsumption struct{}

type UpdateTimeConsumption struct {
	// EndTime    int64  `json:"endtime" form:"endtime"`
	Content    string `json:"content" form:"content"`
	WastedTime bool   `json:"wastedtime" form:"wastedtime"`
	TagID      uint   `json:"tagid" form:"tagid"`
}

type DeleteTimeConsumption struct{}

type ListTasksService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	Pagesize int `json:"page_size" form:"page_size"`
}


func (service *CreateTimeConsumption) Create(id uint) serializer.Response {
	var user modle.User
	code := e.SUCCESS
	// 检查用户是否存在
	modle.DB.First(&user, id)
	timeCunsumption := modle.TimeConsumption{
		UserID:     user.ID,
		StartTime:  service.StartTime,
		EndTime:    time.Now().Unix(),
		Content:    service.Content,
		WastedTime: service.WastedTime,
		TagID:      service.TagID,
	}
	err := modle.DB.Create(&timeCunsumption).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTimeConsumption(timeCunsumption),
		Msg:    e.GetMsg(code),
	}
}

func (service *ShowTimeConsumption) Show(id string) serializer.Response {
	var timeCon modle.TimeConsumption
	code := e.SUCCESS
	err := modle.DB.First(&timeCon, id).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTimeConsumption(timeCon),
		Msg:    e.GetMsg(code),
	}
}

func (service UpdateTimeConsumption) Update(id string) serializer.Response{
	var timeCon modle.TimeConsumption
	code := e.SUCCESS
	modle.DB.First(&timeCon, id)
	timeCon.EndTime = time.Now().Unix()
	timeCon.Content = service.Content
	timeCon.WastedTime = service.WastedTime
	timeCon.TagID = service.TagID
	err := modle.DB.Model(&modle.TimeConsumption{}).Save(&timeCon).Error
	if err != nil{
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   "修改成功",
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteTimeConsumption)Delete(id string)serializer.Response{
	var timeCon modle.TimeConsumption
	code := e.SUCCESS
	err := modle.DB.First(&timeCon,id).Error
	if err != nil{
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = modle.DB.Delete(&timeCon,id).Error
	if err != nil{
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   "删除成功",
		Msg:    e.GetMsg(code),
	}
}


func (service *ListTasksService) List(uid uint,start,end int64) serializer.Response {
	var timeconsumptions []modle.TimeConsumption
	var total int64

	// 有分页
	if service.Pagesize == 0 {
		service.Pagesize = 15
	}
	modle.DB.Model(&modle.TimeConsumption{}).Preload("time_consumption").Where("tag_id=?", uid).Where("end_time BETWEEN ? AND ?", start, end).Count(&total).Limit(service.Pagesize).Offset((service.PageNum - 1) * service.Pagesize).Find(&timeconsumptions)
	return serializer.BuildListResponse(serializer.BuildTimeConsumptions(timeconsumptions), uint(total))
}


// 获取时间消耗表里的实际用时
func CountTime(id uint,timeStart,timeEnd int64) (returnTime string ){
	listService := ListTasksService{}
	res := listService.List(id,timeStart,timeEnd)
	counthour := 0
	countminute := 0
	if list, ok := res.Data.(serializer.DataList);ok{
		if list1,ok := list.Item.([]serializer.TimeConsumption);ok{
			len := len(list1)
			if len != 0{
				for i :=0;i<len;i++{
					// 计算时间差
					// 如果有个任务跨越星期天和星期一，则把这个任务截开，周一的算在这周时间内
					if list1[i].StartTime<timeStart{
						list1[i].StartTime = timeStart
					}
					time1 := time.Unix(list1[i].StartTime, 0)
					time2 := time.Unix(list1[i].EndTime, 0)

					// 计算时间差
					duration := time2.Sub(time1)

					// 将差转换为小时和分钟
					hours := int(duration.Hours())
					minutes := int(duration.Minutes()) % 60
					counthour += hours
					countminute += minutes
					if countminute >= 60 {
						counthour += countminute / 60
						countminute %= 60
					}
					// 构造返回时间字符串
					returnTime = fmt.Sprintf("%v小时%v分", counthour,countminute)
				}
			}
		}
	}
	return returnTime
}


