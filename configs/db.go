package configs

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB
var MysqlDbErr error

func GetDbConfig() map[string]string {
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

	return dbConfig
}

func InitDB() {
	// get db config
	// db err instance
	dbConfig := GetDbConfig()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
	)

	// connect and open db connection
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)

	if MysqlDbErr != nil {
		panic("database data source name error: " + MysqlDbErr.Error())
	}

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(dbConfig["DB_MAX_OPEN_CONNS"])
	MysqlDb.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(dbConfig["DB_MAX_IDLE_CONNS"])
	MysqlDb.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig["DB_MAX_LIFETIME_CONNS"])
	MysqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	// check db connection at once avoid connect failed
	// else error will be reported until db first sql operate
	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("database connect failed: " + MysqlDbErr.Error())
	}
}

func GetDB() *sql.DB {
	return MysqlDb
}
