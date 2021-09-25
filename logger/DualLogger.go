package logger

import (
	"fmt"
	"time"
)

type DualLogger struct {

}

func (d *DualLogger) Log(message LogMessage) {
	panic("implement me")
}

func (d *DualLogger) LogRuntime(message LogMessage, skip int) {
	panic("implement me")
}

func (d *DualLogger) Debug(message string) {
	m:=LogMessage{
		EventLevel:     Event_DEBUG,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    message,
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Debugf(message string, a ...interface{}) {
	//panic("implement me")
	m:=LogMessage{
		EventLevel:     Event_DEBUG,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    fmt.Sprintf(message, a...),
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Info(message string) {
	m:=LogMessage{
		EventLevel:     Event_INFO,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    message,
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Infof(message string, a ...interface{}) {
	m:=LogMessage{
		EventLevel:     Event_INFO,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    fmt.Sprintf(message, a...),
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Warning(message string) {
	m:=LogMessage{
		EventLevel:     Event_WARNING,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    message,
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Warningf(message string, a ...interface{}) {
	m:=LogMessage{
		EventLevel:     Event_WARNING,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    fmt.Sprintf(message, a...),
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Error(message string) {
	m:=LogMessage{
		EventLevel:     Event_ERROR,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    message,
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Errorf(message string, a ...interface{}) {
	m:=LogMessage{
		EventLevel:     Event_ERROR,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    fmt.Sprintf(message, a...),
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Fatal(message string) {
	m:=LogMessage{
		EventLevel:     Event_FATAL,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    message,
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

func (d *DualLogger) Fatalf(message string, a ...interface{}) {
	m:=LogMessage{
		EventLevel:     Event_FATAL,
		ServiceType:    0,
		EventType:      0,
		Number1:        0,
		Number2:        0,
		Number3:        0,
		AdditionalInfo: "",
		TextMessage:    fmt.Sprintf(message, a...),
		Timestamp:      time.Time{},
	}
	d.LogRuntime(m, 2)
}

