package logger

import (
	"sync"
	"time"
)

// Locker is container data
type Locker struct {
	data sync.Map
}

type (
	// Key of context
	Key int

	// Flags is key for store context
	Flags string
)

const (
	logKey              = Key(27)
	_StatusCode   Flags = `StatusCode`
	_Response     Flags = `Response`
	_Messages     Flags = `Messages`
	_ThirdParties Flags = `ThirdParties`
)

// DataLogger is standard output to terminal
type DataLogger struct {
	RequestID     string       `json:"request_id"`
	TimeStart     time.Time    `json:"time_start"`
	Host          string       `json:"host"`
	Endpoint      string       `json:"endpoint"`
	RequestHeader string       `json:"request_header"`
	RequestBody   string       `json:"request_body"`
	StatusCode    int          `json:"status_code"`
	Response      interface{}  `json:"response"`
	ExecTime      float64      `json:"exec_time"`
	Messages      []string     `json:"log_messages"`
	ThirdParties  []ThirdParty `json:"outgoing_log"`
}

// ThirdParty is data logging for any request to third party (outside)
type ThirdParty struct {
	URL           string  `json:"url"`
	RequestHeader string  `json:"request_header"`
	RequestBody   string  `json:"request_body"`
	Response      string  `json:"response"`
	Method        string  `json:"method"`
	StatusCode    int     `json:"status_code"`
	ExecTime      float64 `json:"execution_time"`
}
