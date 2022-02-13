package configs

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义一个全局对象db
var DB *gorm.DB

func GetDbConfig() string {
	// 初始化数据库配置map
	dbConfig := make(map[string]string)

	dbConfig["DB_HOST"] = "180.184.74.86"
	dbConfig["DB_PORT"] = "3306"
	dbConfig["DB_NAME"] = "course"
	dbConfig["DB_USER"] = "root"
	dbConfig["DB_PWD"] = "bytedancecamp"
	dbConfig["DB_CHARSET"] = "utf8"

	dbConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
	)

	return dbDSN
}

// 定义一个初始化数据库的函数
func InitDB() {
	dsn := GetDbConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db.Debug()
	if err != nil {
		panic(fmt.Sprintf("open mysql failed, err is %s", err))
	}
}

// 声明一个全局的rdb变量
var Rdb *redis.Client

// 初始化连接
func InitClient() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "180.184.74.86:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("open mysql failed, err is %s", err))
	}
}
