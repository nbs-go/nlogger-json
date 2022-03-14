package jsonLogger_test

import (
	"context"
	"errors"
	jsonLogger "github.com/nbs-go/nlogger-json"
	"github.com/nbs-go/nlogger/v2"
	logContext "github.com/nbs-go/nlogger/v2/context"
	"github.com/nbs-go/nlogger/v2/level"
	logOption "github.com/nbs-go/nlogger/v2/option"
	"testing"
)

func TestPrint(t *testing.T) {
	p := jsonLogger.NewPrinter(nil)
	log := nlogger.NewStdLogger(p, logOption.Level(level.Debug))
	log.Debug("test print")
	log.Debugf("test formatted print. Message = %s", "hello")
	log.Debug("test print with request id", logOption.Context(logContext.SetRequestId(context.Background(), "100")))
	log.Error("test print with error", logOption.Error(errors.New("internal error")))
	log.Debug("test log with metadata", logOption.AddMetadata("key", "value"))
}

func TestRegister(t *testing.T) {
	nlogger.Register(jsonLogger.New("", "DEBUG", nil))
	log := nlogger.Get()
	log.Debug("test register")
}

func TestNamespace(t *testing.T) {
	p := jsonLogger.NewPrinter(nil)
	log := nlogger.NewStdLogger(p, logOption.Level(level.Debug), logOption.WithNamespace("my-app"))
	log.Debug("this log message contains namespace")
}
