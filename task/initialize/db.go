package initialize

import (
	"fmt"
	"task/global"
	"task/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	mysqlConfig := global.MySqlConfig
	host := mysqlConfig.Host
	prot := mysqlConfig.Port
	database := mysqlConfig.Database
	username := mysqlConfig.Username
	password := mysqlConfig.Password
	charset := mysqlConfig.Charset

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, prot, database, charset)

	ormLogger := logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  // 禁用 datetime 的精度
		DontSupportRenameIndex:    true,  // 重命名索引的时候采用删除并新建的方式
		DontSupportRenameColumn:   true,  // 用 change 重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 最大打开数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	global.DB = db
	model.InitTask()
}
