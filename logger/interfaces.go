package logger

import (
	_ "github.com/brianvoe/gofakeit/v6"
	"time"
)

type EventLevel int16
type ServiceType int16
type EventType int16

const (
	//TODO Consider about the way to inform about logging level changes
	LOGGER_EVENT EventLevel = -1
	//----------------
	OFF EventLevel = 0
)
const (
	Event_FATAL EventLevel = (iota + 1) * 1000
	Event_ERROR
	Event_WARNING
	Event_INFO
	Event_DEBUG
	Event_TRACE
)

const (
	ServiceUnspecified ServiceType = 0
	EventUnspecified   EventType   = 0
)

const Service_Logger_itself = -1
const (
	Logger_Started EventType = iota + 1
	Logger_WriterLevelChanged
	//...
	Logger_Terminated EventType = 9999
)

type LogMessage struct {
	EventLevel     `fake:"{number:1,6}"`
	ServiceType    `fake:"{uint8}"`
	EventType      `fake:"{uint8}"`
	Number1        int    `fake:"{uint8}"`
	Number2        int    `fake:"{uint8}"`
	Number3        int    `fake:"{uint8}"`
	AdditionalInfo string `fake:"{sentence:3}"`
	TextMessage    string `fake:"{sentence:3}"`
	Timestamp      time.Time
}

type StructuralLogger interface {
	Log(message LogMessage)
	LogRuntime(message LogMessage, skip int)
}

type LegacyLogger interface {
	Debug(message string)
	Debugf(message string, a ...interface{})

	Info(message string)
	Infof(message string, a ...interface{})

	Warning(message string)
	Warningf(message string, a ...interface{})

	Error(message string)
	Errorf(message string, a ...interface{})

	Fatal(message string)
	Fatalf(message string, a ...interface{})
}

type StringWriter interface {
	WriteLn(string) error
}

type StringMessageBuilder interface {
	MessageToString(message *LogMessage) (string, error)
}

type TimeProvider interface {
	GetCurrentTime() time.Time
}

type TimeFormatter interface {
	FormatTime(t time.Time) string
}
