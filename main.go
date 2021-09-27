package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"logger/logger"
	"time"
)

var Logger *logger.DualLogger

func main() {
	fmt.Println("Logger")

	consoleStringBuilder := &logger.CSVStringBuilder{
		ColumnSeparator: ",",
		TimeFormatter:   logger.NewStdTimeFormatter(".", " ", ":"),
	}

	csvWriter := logger.NewCSVFileWriter()

	Logger = logger.NewDualLogger(
		logger.Event_DEBUG,
		consoleStringBuilder,
		logger.Event_DEBUG,
		csvWriter,
		consoleStringBuilder,
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

	var lm logger.LogMessage
	for i := 0; i < 250000; i++ {
		gofakeit.Struct(&lm)
		lm.EventLevel *= 1000
		lm.Timestamp = time.Now()
		Logger.Log(&lm)
	}
}
