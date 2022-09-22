package customeError

import (
	"fmt"
)

type ErrorInterface interface {
	Error() string
}

type ServerError struct {
	StatusCode int
	Msg       string
}

func (c *ServerError) Error() string {
	return fmt.Sprintf("status %d: err %v", c.StatusCode, c.Msg)
}
