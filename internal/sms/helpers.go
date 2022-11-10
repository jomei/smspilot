package sms

import (
	"SmsPilot2/repository/smsDB/templateDB"
	"errors"
	"fmt"
	"regexp"
)

func CheckMarcos(sms Sms, template templateDB.TemplateDB) error {
	// выпарсим из шаблона регепсом все максроы
	re := regexp.MustCompile(`{{.*?}}`)
	match := re.FindAllString(template.Text, -1)
	fmt.Println("макросы которые достали из шаблона: ", match)
	// если есть в шаблоне макросы будем проверять какие передали
	if len(match) > 0 {
		// проверим есть ли макросы из шаблона в запросе пользователя
		if sms.Macros != nil {
			for _, macrosName := range match {
				if _, ok := sms.Macros[macrosName]; !ok {
					return errors.New(macrosName)
				}
			}
		}
	}
	return nil
}
