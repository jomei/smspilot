package templateDB

import (
	"SmsPilot2/pkg/config"
	"SmsPilot2/pkg/loggme"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetTemplateById(idTemplate int) (TemplateDB, error) {
	var template TemplateDB
	db, err := gorm.Open(mysql.Open(config.Conf.DNS), &gorm.Config{})
	loggme.Log.PrintAndWriteError("IsUserValid - ", 1, 1, err)
	// делаем запрос и логин и пароля, если будет выборка вернем true если выборки нет вернем false
	db.Where("id = ?", idTemplate).Table("templates").Find(&template)

	if template.Status != 1 {
		return template, errors.New("bad template")
	}
	return template, nil
}
