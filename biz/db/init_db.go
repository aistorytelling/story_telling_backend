package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var dbClient *gorm.DB
var dbInitLock = sync.Mutex{}

func GetDBClient() *gorm.DB {
	var err error
	if dbClient == nil {
		dbInitLock.Lock()
		if dbClient == nil {
			dbClient, err = connectDB()
			if err != nil {
				panic(fmt.Sprintf("连接数据库失败, err: %v", err))
			}
		}
		dbInitLock.Unlock()
	}
	return dbClient
}

func connectDB() (*gorm.DB, error) {
	dsn := "root:Rewq321.@tcp(43.136.30.125:3306)/ai_audio?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
