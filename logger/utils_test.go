package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkActualTimeProvider_GetCurrentTime(b *testing.B) {
	atp := ActualTimeProvider{}
	for i := 0; i < b.N; i++ {
		atp.GetCurrentTime()
	}
}

func TestStdTimeFormatter_FormatTime(t *testing.T) {
	stf := NewStdTimeFormatter("_", " ", ":")
	tn := time.Now()
	expected := tn.Format("02_01_2006 15:04:05")
	actual := stf.FormatTime(tn)
	assert.Equal(t, expected, actual)
}

func BenchmarkStdTimeFormatter_FormatTime(b *testing.B) {
	stf := NewStdTimeFormatter("_", " ", ":")
	for i := 0; i < b.N; i++ {
		stf.FormatTime(time.Now())
	}
}

func BenchmarkStdTimeFormatter_GetAndFormatTime(b *testing.B) {
	atp := ActualTimeProvider{}
	stf := NewStdTimeFormatter("_", " ", ":")
	for i := 0; i < b.N; i++ {
		tn := atp.GetCurrentTime()
		stf.FormatTime(tn)
	}
}

func BenchmarkStdTimeFormatter_GetAndFormatTimeParallel(b *testing.B) {
	atp := ActualTimeProvider{}
	stf := NewStdTimeFormatter("_", " ", ":")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tn := atp.GetCurrentTime()
			stf.FormatTime(tn)
		}
	})
}
