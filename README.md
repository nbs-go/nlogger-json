# nlogger-json

[![Go Report Card](https://goreportcard.com/badge/github.com/nbs-go/nlogger-json)](https://goreportcard.com/report/github.com/nbs-go/nlogger-json)
[![GitHub license](https://img.shields.io/github/license/nbs-go/nlogger-json)](https://github.com/nbs-go/nlogger/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/nbs-go/nlogger-json/branch/master/graph/badge.svg?token=K2K4FA6ZX2)](https://codecov.io/gh/nbs-go/nlogger-json)

`nbs-go/nlogger` Implementation that output to JSON

## Install

```shell
go get github.com/nbs-go/nlogger-json
```

## Usage

```go
package main

import (
	jsonLogger "github.com/nbs-go/nlogger-json"
	"github.com/nbs-go/nlogger/v2"
	"os"
)

// Declare global logger
var log nlogger.Logger

func init() {
	// Register logger here
	lvl := os.Getenv("LOG_LEVEL")
	namespace := os.Getenv("LOG_NAMESPACE")
	nlogger.Register(jsonLogger.NewLogger(namespace, lvl, nil))

	// Retrieve logger instance
	log = nlogger.Get()
}

func main() {
	log.Info("Hello World")
}

```

## Contributors

- Saggaf Arsyad <saggaf.arsyad@gmail.com>