package modle

import (
	"AAL_time/package/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

const (
	maxIdleConnections = 20
	maxOpenConnections = 100
	connMaxLifetime    = time.Second * 30
)



func InitializeDatabase(constring string) {
	enableLogging := true
	if gin.Mode() != "release"{
		enableLogging = false
	}
	db,err := batabase(constring,enableLogging)
	if err != nil {
		utils.LogrusObj.Fatalf("Failed to initialize database: %v", err)
	}
	// 声明了全局变量，在这种情况下，不需要在关闭数据库连接的地方使用 defer，因为连接是在程序退出时关闭的。
	// 函数执行完毕后，不需要手动关闭数据库连接，因为连接是在程序退出时关闭的。
	// defer db.Close()
	DB = db

	migration()
}

func batabase(constring string,enableLogging bool)(*gorm.DB, error){
	db,err := gorm.Open("mysql",constring)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL database: %w", err)
	}

	// 不同的运行模式
	if enableLogging {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(maxIdleConnections)
	db.DB().SetMaxOpenConns(maxOpenConnections)
	db.DB().SetConnMaxLifetime(connMaxLifetime)

	return db, nil
}