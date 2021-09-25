package logger

import (
	"fmt"
	"time"
)

type ActualTimeProvider struct {
}

func (a *ActualTimeProvider) GetCurrentTime() time.Time {
	return time.Now().UTC()
}

type CSVStringBuilder struct {
	ColumnSeparator string
	TimeFormatter
}

//func (C CSVStringBuilder) MessageToString(message LogMessage) (string, error) {
//	var tmpStr string
//
//
//}

func (C *CSVStringBuilder) MessageToString(message LogMessage) (string, error) {
	var tmpStr string


}

func GetEventLevelChar (l EventLevel) string {
	switch l {
	case LOGGER_EVENT:
		return "L"
	case Event_FATAL:
		return "F"
	case Event_ERROR:
		return "E"
	case Event_WARNING:
		return "W"
	case Event_INFO:
		return "I"
	case Event_DEBUG:
		return "D"
	case Event_TRACE:
		return "T"
	default:
		return "U"
	}
}



type StdTimeFormatter struct {
	format string
}

func NewStdTimeFormatter(
	dateComponentSeparator string,
	dateTimeSeparator string,
	timeComponentSeparator string,
) *StdTimeFormatter {

	return &StdTimeFormatter{
		format: fmt.Sprintf("02%s01%s2006%s15%s04%s05", dateComponentSeparator, dateComponentSeparator,
			dateTimeSeparator,
			timeComponentSeparator, timeComponentSeparator),
	}
}
func (stf *StdTimeFormatter) FormatTime(t time.Time) string {

	return t.Format(stf.format)
}
