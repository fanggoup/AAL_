package modle

// 自动迁移模式
func migration(){
	var tableOptions = "charset=utf8mb4"

	DB.Set("gorm:table_options", tableOptions).
		AutoMigrate(&User{}, &TimeConsumption{}, &Category{})
}