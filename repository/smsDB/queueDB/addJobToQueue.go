package queueDB

import (
	"SmsPilot2/pkg/config"
	"SmsPilot2/pkg/loggme"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AddJobToQueue(jobSendSmsDB JobSendSmsDB) JobSendSmsDB {
	db, err := gorm.Open(mysql.Open(config.Conf.DNS), &gorm.Config{})
	loggme.Log.PrintAndWriteError("AddJobToQueue - ", 1, 1, err)
	// пытаемся добавить в очередь отправки сообщений
	_ = db.Table("queue-send-sms").Create(&jobSendSmsDB)
	return jobSendSmsDB
}
