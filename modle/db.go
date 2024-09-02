package modle

import (
	"ERoBP/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	//Uuid     string `gorm:"column:uuid;type:varchar(36);not null"`
	Name      string `gorm:"column:name;type:varchar(20);not null"`
	Password  string `gorm:"column:password;type:varchar(60);not null"`
	Phone     string `gorm:"column:phone;type:varchar(20);not null"`
	Role      string `gorm:"column:role;type:varchar(20);not null"`
	AvatarUrl string `gorm:"column:avatar_url;type:varchar(255);not null"`
}

var Db = createDb()

func createDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Conf.MySQL.GormDNS), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		logrus.Error(err)
	}
	return db
}
