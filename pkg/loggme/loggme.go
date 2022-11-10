package loggme

import (
	"SmsPilot2/pkg/config"
	"log"
	"os"
)

type LoggMe struct {
}

var Log = LoggMe{}

func (l LoggMe) PrintAndWriteError(message string, typeInfo int, togglePanic int, err error) {
	//switch {
	//// выводим и записываем все типы ошибок и логов
	//case config.Conf.LogLevel == 3:
	//	if typeInfo <= 3 {
	//		// запись в файл
	//		writerInFile("INFO: "+message, err)
	//		// вывод на экран принтом
	//		fmt.Println("INFO: "+message, err)
	//		if togglePanic != 1 {
	//			log.Panic(err)
	//		}
	//	}
	//// выводим и записываем только критические ошибки
	//case config.Conf.LogLevel == 2:
	//	if typeInfo <= 2 {
	//		// запись в файл
	//		writerInFile("WARNING: "+message, err)
	//		// вывод на экран принтом
	//		fmt.Println("WARNING: "+message, err)
	//		if togglePanic != 1 {
	//			log.Panic(err)
	//		}
	//	}
	//// выводим и записываем только системные ошибки типа нет конекта с базой данных
	//case config.Conf.LogLevel == 1:
	//	if typeInfo == 1 {
	//		// запись в файл
	//		writerInFile("DANGER: "+message, err)
	//		// вывод на экран принтом
	//		fmt.Println("DANGER: "+message, err)
	//		if togglePanic != 1 {
	//			log.Panic(err)
	//		}
	//	}
	//default:
	//	return
	//}
}

func writerInFile(message string, err error) {
	f, e := os.OpenFile(config.Conf.LOGFILE, os.O_APPEND|os.O_WRONLY, 0644)
	if e != nil {
		log.Panic(e)
	}
	defer f.Close()

	if _, e = f.WriteString(message + ": " + err.Error() + "\n"); e != nil {
		log.Panic(e)
	}
}
