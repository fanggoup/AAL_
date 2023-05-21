package service

import (
	"AAL_time/modle"
	"AAL_time/package/e"
	"AAL_time/package/utils"
	"AAL_time/serializer"
)

type CreateCategory struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

type ShowAllCategory struct{}

type UpdateCategory struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func (service *CreateCategory) Create(id uint) serializer.Response {
	var tag modle.Category
	code := e.SUCCESS
	countUser := modle.DB.Find(&tag, id).RowsAffected
	if countUser != 0{
		code = e.ErrorExistCategory
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	category := modle.Category{
		UserID:      id,
		Name:        service.Name,
		Description: service.Description,
	}
	err := modle.DB.Create(&category).Error
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
		Data:   serializer.BuildCategory(category),
		Msg:    e.GetMsg(code),
	}
}

// 查询标签（返回用户所有标签）
func (service *ShowAllCategory) ShowAll(id uint) serializer.Response {
	var tags []modle.Category
	var total int64
	code := e.SUCCESS

	err := modle.DB.Model(&modle.Category{}).Where("user_id=?", id).Count(&total).Find(&tags).Error

	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCategorys(tags), uint(total))

}

// 修改标签
func (service *UpdateCategory) Update(id string)serializer.Response{
	var tags modle.Category
	code := e.SUCCESS
	modle.DB.Model(&modle.Category{}).Where("id=?",id).First(&tags)
	tags.Name = service.Name
	tags.Description = service.Description
	err := modle.DB.Save(&tags).Error
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