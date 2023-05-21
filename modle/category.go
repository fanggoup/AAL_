package modle

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UserID     uint       // 用户ID（外键，关联用户表）
	Name        string // 分类名称
	Description string // 分类描述
}