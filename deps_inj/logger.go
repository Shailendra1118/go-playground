package depsinj

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type DefaultLogger struct {
}

// implementation of interface
func (logger DefaultLogger) Log(message string) {
	fmt.Println("Logging message:", message)
}
