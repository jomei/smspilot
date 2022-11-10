package sms

import (
	"SmsPilot2/repository/smsDB/queueDB"
	"SmsPilot2/repository/smsDB/templateDB"
	"regexp"
	"strings"
)

func GeneratorMessage(sms Sms, template templateDB.TemplateDB, phoneInt int, Userid int) queueDB.JobSendSmsDB {
	var message queueDB.JobSendSmsDB

	message.Userid = Userid
	message.Phone = phoneInt
	message.DatetimeSend = sms.DataTimeSendUTC

	// заменяем макросы в шаблоне тем, что прислали в смс
	// выпарсим все макросы из темплата
	re := regexp.MustCompile(`{{.*?}}`)
	match := re.FindAllString(template.Text, -1)
	// если макросов не было найдено
	if len(match) == 0 {
		message.Message = template.Text
		return message
	}
	// если макросы были найдены
	// перебираем макросы найденные в шаблоне из базы
	for _, nameMacros := range match {
		template.Text = strings.Replace(template.Text, nameMacros, sms.Macros[nameMacros], -1)
	}
	message.Message = template.Text

	return message
}
