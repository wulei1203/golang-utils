package utils

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var wdatabase *gorm.DB
var rdatabase *gorm.DB

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		Log.WithFields(log.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
	}
	if v[0] == "log" {
		Log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

func GetWriteDB() *gorm.DB {
	return wdatabase
}

func GetReadDB() *gorm.DB {
	return rdatabase
}

func CloseWriteDB() {
	if wdatabase != nil {
		wdatabase.DB().Close()
	}
}

func CloseReadDB() {
	if rdatabase != nil {
		rdatabase.DB().Close()
	}
}

//新增前的回调
func createSelfCall(scope *gorm.Scope) {
	if scope.HasColumn("create_time") {
		scope.SetColumn("create_time", time.Now())
	}
	if scope.HasColumn("active") {
		scope.SetColumn("active", 1)
	}
}

//删除前的回调
func deleteSelfCall(scope *gorm.Scope) {
	if scope.HasColumn("deletedate") {
		scope.SetColumn("deletedate", time.Now())
	}

	if scope.HasColumn("isdelete") {
		scope.SetColumn("isdelete", true)
	}
}

func InitDB(data map[string]string) *gorm.DB {
	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		data["dbhost"], data["dbport"], data["dbuser"], data["dbpwd"], data["dbName"])

	log.Printf("Creating a new connection: %v", connstr)

	db, err := gorm.Open("postgres", connstr)
	if err != nil {
		panic(err.Error())
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err.Error())
	}

	// 全局禁用表名复数
	db.SingularTable(true)

	//开启日志模式
	db.LogMode(true)

	//给所有表增加默认前缀
	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return data["prefix"] + defaultTableName
	}

	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)
	rdatabase = db

	//增加新增时的回调函数
	db.Callback().Create().Replace("gorm:before_create", createSelfCall)
	wdatabase = db

	return db
}
