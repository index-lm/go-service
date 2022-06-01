package model

import (
	"go-service/pkg/db"
	"go-service/pkg/log"
	"gorm.io/gorm"
)

func InitDb(db *gorm.DB) {
	var err error
	// 创建表时添加后缀

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Account{})
	if err != nil {
		log.Error("sys", "迁移数据库失败")
	}
}

type Account struct {
	db.BaseModel
	ClientId  uint64  `gorm:"type:bigint;unsigned;not null"`
	Nickname  string  `gorm:"type:varchar(32);not null"`
	Sex       uint8   `gorm:"type:tinyint;unsigned;not null;default:0"`
	AvatarUrl *string `gorm:"type:varchar(255)"`
	Extra     *string `gorm:"type:varchar(1024)"`
}

func (Account) TableName() string {
	return "am_account"
}
