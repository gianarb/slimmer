package logger

import (
	"fmt"
	"testing"
)

func TestReturnCorrentLen(t *testing.T) {
	c := Console{}
	l, _ := c.Write([]byte{'t'})
	fmt.Print(l)
}
