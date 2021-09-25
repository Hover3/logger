package logger

import "time"

type EventLevel int
type ServiceType int
type EventType int
const (
	//TODO Consider about the way to inform about logging level changes
	LOG_MODE EventLevel = -1
	//----------------
	OFF EventLevel= 0
	Event_FATAL EventLevel= (iota +1) *1000
	Event_ERROR
	Event_WARNING
	Event_INFO
	Event_DEBUG
	Event_TRACE
)

type LogMessage struct {
	EventLevel
	ServiceType
	EventType
	Number1 int
	Number2 int
	Number3 int
	AdditionalInfo string
	TextMessage string
}

type TimeStampedLogMessage struct {
	LogMessage
	Timestamp time.Time
}
