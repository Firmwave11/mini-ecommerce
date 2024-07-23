package logger

import (
	"context"
	"fmt"
)

type msg struct {
	text string
}

// LogMessage create contract for log any message
type LogMessage interface {
	To(ctx context.Context)
}

func (m msg) To(ctx context.Context) {
	value, ok := extract(ctx)
	if !ok {
		return
	}

	value.StoreMessage(m.text)
}

// Message default log message
func Message(log ...interface{}) LogMessage {
	return msg{
		text: fmt.Sprint(log...),
	}
}

// Messagef log message with format
func Messagef(format string, log ...interface{}) LogMessage {
	return msg{
		text: fmt.Sprintf(format, log...),
	}
}
