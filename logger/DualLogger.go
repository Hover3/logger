package logger

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type DualLogger struct {
	consoleLogLevel      EventLevel
	consoleLevelMutex    sync.RWMutex
	consoleWriteMutex    sync.Mutex
	consoleStringBuilder StringMessageBuilder

	csvLogLevel      EventLevel
	csvLevelMutex    sync.RWMutex
	csvWriteMutex    sync.Mutex
	csvWriter        StringWriter
	csvStringBuilder StringMessageBuilder

	timeSource TimeProvider
}

func NewDualLogger(consoleLogLevel EventLevel,
	consoleStringBuilder StringMessageBuilder,
	csvLogLevel EventLevel,
	csvWriter StringWriter,
	csvStringBuilder StringMessageBuilder,
	timeSource TimeProvider,
) *DualLogger {

	return &DualLogger{
		consoleLogLevel:      consoleLogLevel,
		consoleStringBuilder: consoleStringBuilder,
		csvLogLevel:          csvLogLevel,
		csvWriter:            csvWriter,
		csvStringBuilder:     csvStringBuilder,
		timeSource:           timeSource,
	}
}

func (d *DualLogger) SetConsoleLogLevel(level EventLevel) {
	d.consoleLevelMutex.Lock()
	defer d.consoleLevelMutex.Unlock()
	d.consoleLogLevel = level
}

func (d *DualLogger) GetConsoleLogLevel() EventLevel {
	d.consoleLevelMutex.RLock()
	defer d.consoleLevelMutex.RUnlock()
	return d.consoleLogLevel
}

func (d *DualLogger) SetCSVLogLevel(level EventLevel) {
	d.csvLevelMutex.Lock()
	defer d.csvLevelMutex.Unlock()
	d.csvLogLevel = level
}

func (d *DualLogger) GetCSVLogLevel() EventLevel {
	d.csvLevelMutex.RLock()
	defer d.csvLevelMutex.RUnlock()
	return d.csvLogLevel
}

func (d *DualLogger) Log(message *LogMessage) {
	message.Timestamp = d.timeSource.GetCurrentTime()
	//TODO review

	//printing to screen

	if message.EventLevel <= d.consoleLogLevel {
		d.writeToConsole(message)
	}

	//Writing to file
	if message.EventLevel <= d.csvLogLevel {
		d.writeToCSV(message)
	}

}
func (d *DualLogger) writeToConsole(message *LogMessage) {
	d.consoleWriteMutex.Lock()
	defer d.consoleWriteMutex.Unlock()
	if d.consoleStringBuilder == nil {
		fmt.Println("Logger: Console string builder is not set!")
		return
	}
	msg, _ := d.consoleStringBuilder.MessageToString(message)
	fmt.Println(msg)
}
func (d *DualLogger) writeToCSV(message *LogMessage) {
	d.csvWriteMutex.Lock()
	defer d.csvWriteMutex.Unlock()
	if d.csvStringBuilder == nil {
		fmt.Println("Logger: CSV string builder is not set!")
		return
	}
	msg, _ := d.csvStringBuilder.MessageToString(message)
	d.csvWriter.WriteLn(msg + "\n")
}

func (d *DualLogger) LogRuntime(message *LogMessage, skip int) {
	fn, line := GetRuntimeInfo(skip)
	message.TextMessage = fmt.Sprintf("%s %s %s", fn, strconv.Itoa(line), message.TextMessage)
	d.Log(message)
}

func (d *DualLogger) Debug(message string) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Debugf(message string, a ...interface{}) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Info(message string) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Infof(message string, a ...interface{}) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Warning(message string) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Warningf(message string, a ...interface{}) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Error(message string) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Errorf(message string, a ...interface{}) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Fatal(message string) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

func (d *DualLogger) Fatalf(message string, a ...interface{}) {
	m := LogMessage{
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
	d.LogRuntime(&m, 2)
}

type CSVFileWriter struct {
	csvFile        *os.File
	filesCount     uint32
	fileNamePrefix string
	stringsCount   uint16
}

func NewCSVFileWriter() *CSVFileWriter {
	prefix := time.Now().Format("02_01_2006_15_04_05_")
	var count uint32 = 0
	f, err := newFile(fmt.Sprintf("%s%06d.csv", prefix, count))
	if err != nil {
		panic(err)
	}

	return &CSVFileWriter{
		csvFile:        f,
		filesCount:     count,
		fileNamePrefix: prefix,
		stringsCount:   0,
	}
}

func (c *CSVFileWriter) WriteLn(s string) error {
	c.stringsCount++
	_, err := c.csvFile.Write([]byte(s))
	if c.stringsCount == ^uint16(0) {
		c.csvFile.Close()
		c.filesCount++
		c.csvFile, err = newFile(fmt.Sprintf("%s%06d.csv", c.fileNamePrefix, c.filesCount))
	}
	return err
}

func newFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeDir)
	if err != nil {
		return nil, err
	}
	_, err = f.Write([]byte("EventLevel,ServiceType,EventType,Number1,Number2,Number3,AdditionalInfo,TextMessage,Timestamp\n"))
	if err != nil {
		return nil, err
	}
	return f, nil
}
