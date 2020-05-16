package orms

import (
	"fmt"
	"strings"
)

// Collect errors for database manipulation
type GormMultiError struct {
	msg []string
}

// Implement Error method
func (g *GormMultiError) Error() string {
	builder := new(strings.Builder)
	for i, v := range g.msg {
		builder.WriteString(fmt.Sprintf("gorms multi error, num: %v, msg: %v", i, v))
	}
	return builder.String()
}

// Add an error to it
func (g *GormMultiError) Add(err error) {
	g.msg = append(g.msg, err.Error())
}

// Should be called when return an error outbound
func (g *GormMultiError) Return() error {
	if len(g.msg) == 0 {
		return nil
	}
	return g
}
