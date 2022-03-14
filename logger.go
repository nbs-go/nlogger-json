package jsonLogger

import (
	"encoding/json"
	"fmt"
	"github.com/nbs-go/nlogger/v2"
	logContext "github.com/nbs-go/nlogger/v2/context"
	"github.com/nbs-go/nlogger/v2/level"
	logOption "github.com/nbs-go/nlogger/v2/option"
	"io"
	"os"
	"strings"
	"time"
)

// New initiate nlogger.StdLogger with json printer
func New(namespace string, lvStr string, w io.Writer) *nlogger.StdLogger {
	lv := level.Parse(lvStr)
	return nlogger.NewStdLogger(NewPrinter(w), logOption.WithNamespace(namespace), logOption.Level(lv))
}

func NewPrinter(w io.Writer, _ ...logOption.SetterFunc) *jsonPrinter {
	// Init writer if not set
	if w == nil {
		w = os.Stdout
	}

	// Get hostname
	hostname, _ := os.Hostname()

	return &jsonPrinter{
		writer:   json.NewEncoder(w),
		hostname: hostname,
	}
}

type jsonPrinter struct {
	writer   *json.Encoder
	hostname string
}

func (j *jsonPrinter) Print(namespace string, outLevel level.LogLevel, msg string, options *logOption.Options) {
	// Init body
	body := map[string]interface{}{
		"timestamp": time.Now(),
		"level":     outLevel,
		"levelStr":  strings.ToUpper(level.String(outLevel)),
		"hostname":  j.hostname,
	}

	// Set namespace
	if namespace != "" {
		body["namespace"] = namespace
	}

	// Format message if it has formatted arguments
	if len(options.FmtArgs) > 0 {
		msg = fmt.Sprintf(msg, options.FmtArgs...)
	}

	// Set message
	body["message"] = msg

	// Set requestId
	if reqId := logContext.GetRequestId(options.Context); reqId != "" {
		body["requestId"] = reqId
	}

	// Set error if set in option and level is at least level.Warn
	if err := logOption.GetError(options, logOption.ErrorKey); outLevel <= level.Warn && err != nil {
		body["error"] = err.Error()
	}

	// Set metadata
	if len(options.Metadata) > 0 {
		body["metadata"] = options.Metadata
	}

	// Write body
	_ = j.writer.Encode(body)
}
