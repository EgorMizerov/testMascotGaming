package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"
)

var levels = map[string]int8{
	"debug":  -1,
	"info":   0,
	"warn":   1,
	"error":  2,
	"dpanic": 3,
	"panic":  4,
	"fatal":  5,
}

var colors = map[string]string{
	"colorReset":  "\033[0m",
	"colorBlack":  "\033[30m",
	"colorRed":    "\033[31m",
	"colorGreen":  "\033[32m",
	"colorYellow": "\033[33m",
	"colorBlue":   "\033[34m",
	"colorPurple": "\033[35m",
	"colorCyan":   "\033[36m",
	"colorWhite":  "\033[37m",
}

var background = map[string]string{
	"colorReset":  "\033[0m",
	"colorBlack":  "\033[30m",
	"colorRed":    "\033[31m",
	"colorGreen":  "\033[32m",
	"colorYellow": "\033[33m",
	"colorBlue":   "\033[34m",
	"colorPurple": "\033[35m",
	"colorCyan":   "\033[36m",
	"colorWhite":  "\033[37m",
}

func ConsoleCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	ln := strings.Split(caller.TrimmedPath(), ":")[1]
	enc.AppendString(string(background["colorWhite"]) + "func: " + caller.Function + " line:" + ln + string(background["colorReset"]))
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/2 - 15:04:05"))
}

func ConsoleLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level.CapitalString() {

	case "INFO":
		enc.AppendString(string(colors["colorWhite"]) + "[" + level.CapitalString() + "]" + string(colors["colorReset"]))
	case "WARN":
		enc.AppendString(string(colors["colorYellow"]) + "[" + level.CapitalString() + "]" + string(colors["colorReset"]))
	case "DEBUG":
		enc.AppendString(string(colors["colorBlue"]) + "[" + level.CapitalString() + "]" + string(colors["colorReset"]))
	case "ERROR":
		enc.AppendString(string(colors["colorRed"]) + "[" + level.CapitalString() + "]" + string(colors["colorReset"]))
	case "FATAL":
		enc.AppendString(string(colors["colorBlack"]) + "[" + level.CapitalString() + "]" + string(colors["colorReset"]))
	default:
		enc.AppendString(string(colors["colorReset"]) + "[" + level.CapitalString() + "]" + string(colors["colorReset"]))
	}

}

func ConsoleConfig(level string) (*zap.Logger, error) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(levels[level])),
		Encoding:    "console",
		Development: true,
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "message",
			ConsoleSeparator: "\t",
			CallerKey:        "caller",
			//EncodeCaller: ConsoleCallerEncoder,
			TimeKey:     "time",
			EncodeTime:  SyslogTimeEncoder,
			LevelKey:    "level",
			EncodeLevel: ConsoleLevelEncoder,
		},
	}

	logger, err := cfg.Build()
	return logger, err
}

func JSONConfig(level string) (*zap.Logger, error) {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(levels[level])),
		Encoding:         "json",
		Development:      true,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	return logger, err
}
