package logger

import "time"

type ActualTimeProvider struct {

}

func (a *ActualTimeProvider) GetCurrentTime() time.Time {
	return time.Now().UTC()
}

type  CSVStringBuilder struct {
	ColumnSeparator string
	TimeFormatter
}


func (C CSVStringBuilder) MessageToString(message LogMessage) (string, error) {
	var tmpStr string


}


type StdTimeFormatter struct {
	DateComponentSeparator string
	DateTimeSeparator string
	TimeComponentSeparator string
}

func (stf *StdTimeFormatter) FormatTime(t time.Time) string {
	var tmpStr string
	y, m, d:=t.Date()
	h, min, s:=t.Clock()
	tmpStr = string(d) + stf.DateComponentSeparator +
			string(m) + stf.DateComponentSeparator +
			string(y) +stf.DateComponentSeparator +
			stf.DateTimeSeparator +
			string(h) +stf.TimeComponentSeparator +
			string(min) +stf.TimeComponentSeparator +
			string(s)
	return tmpStr
}



