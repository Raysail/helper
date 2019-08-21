package helper

import (
	"testing"
)

func Test_WriteLogger(t *testing.T) {
	logger := NewLogger()
	logger.Debug()
}
