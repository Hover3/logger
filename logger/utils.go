package logger

import (
	"fmt"
	"runtime"
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

func (csb *CSVStringBuilder) MessageToString(message *LogMessage) (string, error) {
	tmpStr:= fmt.Sprint(GetEventLevelChar(message.EventLevel), csb.ColumnSeparator,
			csb.TimeFormatter.FormatTime(message.Timestamp), csb.ColumnSeparator,
			message.ServiceType, csb.ColumnSeparator,
			message.EventType, csb.ColumnSeparator,
			message.Number1, csb.ColumnSeparator,
			message.Number2, csb.ColumnSeparator,
			message.Number3, csb.ColumnSeparator,
			message.AdditionalInfo, csb.ColumnSeparator,
			message.TextMessage, csb.ColumnSeparator,
			)

	return tmpStr, nil
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

func GetRuntimeInfo(skip int) (string, int) {
	_, fn, line, _ := runtime.Caller(skip+1)
	return fn, line
}