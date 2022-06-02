package db

import (
	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"time"
)

var Gorm *gorm.DB

// InitGorm 初始化mysqlOrm
func InitGorm(username string, password string, host string, dbName string, maxIdle int, maxOpen int, initFunc func(db *gorm.DB)) {
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		dbName)
	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connInfo, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置, &gorm.Config{})
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxIdle)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(maxOpen)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	initFunc(db)
	Gorm = db
}

type BaseModel struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;unsigned"`           //主键ID
	GmtCreate   time.Time `gorm:"autoCreateTime;not null"`                     //创建时间
	GmtModified time.Time `gorm:"autoUpdateTime;not null"`                     //更新时间
	IsDelete    bool      `gorm:"type:tinyint(1);unsigned;default:0;not null"` //是否删除
	Remark      *string   `gorm:"type:varchar(2048)"`                          //备注
}
