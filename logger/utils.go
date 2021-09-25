package logger

import "time"

type ActualTimeProvider struct {

}

func (a *ActualTimeProvider) GetCurrentTime() time.Time {
	return time.Now().UTC()
}
