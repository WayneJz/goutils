package errors

import (
	"fmt"
	"strings"
)

// MultiError ... Collect multiple errors
type MultiError struct {
	msg []string
}

// Error ... Implement Error method
func (g *MultiError) Error() string {
	builder := new(strings.Builder)
	for i, v := range g.msg {
		builder.WriteString(fmt.Sprintf("multi error: sequence=%v, err=%v", i, v))
	}
	return builder.String()
}

// Add ... Add an error to MultiError
func (g *MultiError) Add(err error) {
	if err == nil {
		return
	}
	g.msg = append(g.msg, err.Error())
}

// Return ... Should be called when return an error outbound
func (g *MultiError) Return() error {
	if len(g.msg) == 0 {
		return nil
	}
	return g
}
