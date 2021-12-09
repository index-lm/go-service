package initserver

//func initDb(db *gorm.DB) {
//	var err error
//	// 创建表时添加后缀
//	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&po.User{})
//	if err != nil {
//		global.MyLogger.WithFields(logrus.Fields{"err": err}).Error("迁移失败")
//	}
//}
