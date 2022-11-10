package usersDb

import (
	"SmsPilot2/pkg/config"
	"SmsPilot2/pkg/loggme"
	"SmsPilot2/pkg/md5Hash"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func IsUserValid(u *UserDB) bool {
	db, err := gorm.Open(mysql.Open(config.Conf.DNS), &gorm.Config{})
	loggme.Log.PrintAndWriteError("IsUserValid - ", 1, 1, err)
	// делаем запрос и логин и пароля, если будет выборка вернем true если выборки нет вернем false
	db.Where("login = ? AND password >= ?", u.Login, md5Hash.Get(u.Password)).Table("users").Find(&u)
	if u.Status != 1 {
		return false
	}
	return true
}
