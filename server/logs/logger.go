package logs

import (
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/OthloTech/OthloCalendar/server/config"
)

var (
	Fatal *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Debug *log.Logger
	Trace *log.Logger
)

type logLevel int

const (
	fatal logLevel = 1 + iota
	err
	warn
	info
	debug
	trace
)

func init() {
	level := logLevel(math.Min(float64(trace), math.Max(float64(fatal), float64(config.NewConfig().LogLevel))))

	handle := ioutil.Discard
	if level >= fatal {
		handle = os.Stderr
	}
	Fatal = log.New(handle, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	handle = ioutil.Discard
	if level >= err {
		handle = os.Stderr
	}
	Error = log.New(handle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	handle = ioutil.Discard
	if level >= warn {
		handle = os.Stdout
	}
	Warn = log.New(handle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	handle = ioutil.Discard
	if level >= info {
		handle = os.Stdout
	}
	Info = log.New(handle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	handle = ioutil.Discard
	if level >= debug {
		handle = os.Stdout
	}
	Debug = log.New(handle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	handle = ioutil.Discard
	if level >= trace {
		handle = os.Stdout
	}
	Trace = log.New(handle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
}
