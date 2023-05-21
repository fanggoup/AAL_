package modle

import (
	// "time"

	"gorm.io/gorm"
)

type TimeConsumption struct {
	gorm.Model
	UserID     uint       // 用户ID（外键，关联用户表）
	StartTime  int64 // 时间消费开始时间
	EndTime    int64 // 时间消费结束时间
	Content    string    // 时间消费内容
	WastedTime bool      // 是否是浪费的时间
	TagID      uint       // 标签ID（外键，关联标签表）
	// 其他时间消费字段...
}
