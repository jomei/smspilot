package usersDb

import (
	"SmsPilot2/pkg/config"
	"SmsPilot2/pkg/loggme"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func IsUserActive(idUser int) bool {
	var u UserDB
	db, err := gorm.Open(mysql.Open(config.Conf.DNS), &gorm.Config{})
	loggme.Log.PrintAndWriteError("IsUserValid - ", 1, 1, err)
	// делаем запрос и логин и пароля, если будет выборка вернем true если выборки нет вернем false
	db.Where("id = ?", idUser).Table("users").Find(&u)
	if u.Status != 1 {
		return false
	}
	return true
}
