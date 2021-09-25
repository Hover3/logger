package main

import (
	"fmt"
	"logger/logger"
	"time"
)

var Logger *logger.DualLogger
func main() {
	fmt.Println("Logger")


	consoleStringBuilder:= &logger.CSVStringBuilder{
		ColumnSeparator: ";",
		TimeFormatter:   logger.NewStdTimeFormatter(".", " ", ":"),
	}

	Logger=logger.NewDualLogger(
		logger.Event_ERROR,
		consoleStringBuilder,
		logger.Event_DEBUG,
		nil,
		nil,
		&logger.ActualTimeProvider{})

	Logger.Error("Эррор!")
	Logger.Fatal("Паникэ!")
	Logger.Warning("Предупреждение минздрава")
	Logger.Info("Информацыэ")
	Logger.Debug("Отладке")

	Logger.Log(&logger.LogMessage{
		EventLevel:     logger.Event_ERROR,
		ServiceType:    100,
		EventType:      20,
		Number1:        1,
		Number2:        2,
		Number3:        3,
		AdditionalInfo: "",
		TextMessage:    "Превед",
		Timestamp:      time.Time{},
	})

}
