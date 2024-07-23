package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// all variables middleware
const (
	// any error message
	ErrTimezone = "Sorry. server configuration not available time Asia/Jakarta"
	LayoutDate  = "2006-01-02 15:04:05"
	tzLocation  = "Asia/Jakarta"
)

type client struct {
	log *logrus.Logger
	loc *time.Location
}

// Clients create contract middleware packages
type Clients interface {
	// record custom log using Logrus library
	Logger(next http.Handler) http.Handler
}

// NewMiddleware will create an object that represent clients interface
func NewMiddleware(log *logrus.Logger, loc *time.Location) Clients {
	return &client{
		log: log,
		loc: loc,
	}
}

// Timezone load timezone area
func Timezone() *time.Location {
	loc, err := time.LoadLocation(tzLocation)
	if err != nil {
		panic(ErrTimezone)
	}
	return loc
}

// Logrus loggrus formatter json
func Logrus() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: LayoutDate,
	})
	return log
}
