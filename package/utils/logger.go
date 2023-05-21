package utils

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// 日志文件路径生成
var LogrusObj *logrus.Logger

func init(){
	if LogrusObj != nil {
		src, err := setOutputFile()
		if err != nil {
			log.Fatal(err)
		}
		//设置输出
		LogrusObj.Out = src
		return
	}
	logger := logrus.New()
	src, err := setOutputFile()
	if err != nil {
		log.Fatal("文件问题",err)
	}

	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutputFile()(*os.File,error){
	now := time.Now()
	logFilePath := ""
	if dir,err := os.Getwd();err == nil{
		logFilePath = filepath.Join(dir,"logs")
	}

	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	logFileName := now.Format("2006-01-02") + ".log"
	fileName := filepath.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("文件打不开",err)
		return nil, err
	}
	
	return src, nil
}