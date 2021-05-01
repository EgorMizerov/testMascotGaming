package logger

import "go.uber.org/zap"

func NewLogger(level string) (*zap.Logger, error) {
	return ConsoleConfig(level)
}
