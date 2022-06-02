package model

import (
	"fmt"
	"go-service/pkg/db"
	"go-service/pkg/log"
	"gorm.io/gorm"
)

func InitDb(db *gorm.DB) {
	var err error
	// 迁移账号表
	err = db.Set("gorm:table_options",
		"ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin").AutoMigrate(&Account{})
	if err != nil {
		log.Error("sys", fmt.Sprintf("数据库迁移失败:%s",err.Error()))
	}
	// 迁移客户端表
	err = db.Set("gorm:table_options",
		"ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin").AutoMigrate(&Client{})
	if err != nil {
		log.Error("sys", fmt.Sprintf("数据库迁移失败:%s",err.Error()))
	}
	// 迁移设备表
	err = db.Set("gorm:table_options",
		"ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin").AutoMigrate(&Device{})
	if err != nil {
		log.Error("sys", fmt.Sprintf("数据库迁移失败:%s",err.Error()))
	}
	// 迁移好友表
	err = db.Set("gorm:table_options",
		"ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin").AutoMigrate(&Friend{})
	if err != nil {
		log.Error("sys", fmt.Sprintf("数据库迁移失败:%s",err.Error()))
	}
	// 迁移群组表
	err = db.Set("gorm:table_options",
		"ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin").AutoMigrate(&Group{})
	if err != nil {
		log.Error("sys", fmt.Sprintf("数据库迁移失败:%s",err.Error()))
	}
	// 迁移群组成员表
	err = db.Set("gorm:table_options",
		"ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin").AutoMigrate(&GroupUser{})
	if err != nil {
		log.Error("sys", fmt.Sprintf("数据库迁移失败:%s",err.Error()))
	}
}

// 账号表
type Account struct {
	db.BaseModel
	ClientId  uint64  `gorm:"type:bigint;unsigned;not null"`               //客户端ID
	Nickname  string  `gorm:"type:varchar(32);not null"`                   //昵称
	Sex       uint8   `gorm:"type:tinyint(3);unsigned;not null;default:0"` //性别
	AvatarUrl *string `gorm:"type:varchar(255)"`                           //头像地址
	Extra     *string `gorm:"type:varchar(1024)"`                          //附加属性
}

func (Account) TableName() string {
	return "tb_account"
}

//客户端表
type Client struct {
	db.BaseModel
	Name      string  `gorm:"type:varchar(32);not null"` //客户端名称
	code      string  `gorm:"type:varchar(32);not null"` //客户端编码
	SecretKey string  `gorm:"type:varchar(32);not null"` //客户端密钥
	Extra     *string `gorm:"type:varchar(1024)"`        //附加属性
}

func (Client) TableName() string {
	return "tb_client"
}

// 设备表
type Device struct {
	db.BaseModel
	AccountId     uint64  `gorm:"type:bigint;unsigned;not null"`     //账号ID
	Type          uint8   `gorm:"type:tinyint(3);not null;unsigned"` //设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
	Brand         string  `gorm:"type:varchar(20);not null"`         //手机厂商
	Model         string  `gorm:"type:varchar(20);not null"`         // 机型
	SystemVersion string  `gorm:"type:varchar(10);not null"`         // 系统版本
	SDKVersion    string  `gorm:"type:varchar(10);not null"`         // SDK版本
	Status        uint8   `gorm:"type:tinyint(3);not null;unsigned"` // 在线状态，0：离线；1：在线
	ConnAddr      string  `gorm:"type:varchar(25);not null"`         // 连接层服务层地址
	ClientAddr    string  `gorm:"type:varchar(25);not null"`         // 客户端地址
	Extra         *string `gorm:"type:varchar(1024)"`                //附加属性
}

func (Device) TableName() string {
	return "tb_device"
}

//好友表
type Friend struct {
	db.BaseModel
	AccountId uint64  `gorm:"type:bigint;unsigned;not null"`     //账号ID
	FriendId  uint64  `gorm:"type:bigint;unsigned;not null"`     //好友ID
	Remarks   string  `gorm:"type:varchar(32);not null"`         //备注名
	Status    uint8   `gorm:"type:tinyint(3);unsigned;not null"` //状态，1：申请，2：同意
	Extra     *string `gorm:"type:varchar(1024)"`                //附加属性
}

func (Friend) TableName() string {
	return "tb_friend"
}

//群组表
type Group struct {
	db.BaseModel
	GroupName  string  `gorm:"type:varchar(32);not null"`  //备注名
	GroupIntro string  `gorm:"type:varchar(255);not null"` //群组简介
	AvatarUrl  string  `gorm:"type:varchar(255);not null"` //群组头像
	Extra      *string `gorm:"type:varchar(1024)"`         //附加属性
}

func (Group) TableName() string {
	return "tb_group"
}

//群组成员表
type GroupUser struct {
	db.BaseModel
	AccountId  uint64  `gorm:"type:bigint;unsigned;not null"`     //账号ID
	GroupId    uint64  `gorm:"type:bigint;unsigned;not null"`     //群组ID
	MemberType uint8   `gorm:"type:tinyint(3);unsigned;not null"` //成员类型，1：管理员；2：普通成员
	RemarkName string  `gorm:"type:varchar(32);not null"`         //用户备注名
	Extra      *string `gorm:"type:varchar(1024)"`                //附加属性
}

func (GroupUser) TableName() string {
	return "tb_group_user"
}
