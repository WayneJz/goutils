package errors

import (
	"fmt"
	"strings"
)

// Errors ... Collect multiple errors
type Errors struct {
	msg []string
}

// Error ... Implement Error method
func (e *Errors) Error() string {
	builder := new(strings.Builder)
	for i, v := range e.msg {
		builder.WriteString(fmt.Sprintf("errors traceback: sequence=%v, err=%v", i, v))
	}
	return builder.String()
}

// Add ... Add an error to Errors
func (e *Errors) Add(err error) {
	if err == nil {
		return
	}
	e.msg = append(e.msg, err.Error())
}

// Return ... Should be called when return an error outbound
func (e *Errors) Return() error {
	if len(e.msg) == 0 {
		return nil
	}
	return e
}
