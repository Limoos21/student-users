package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger инициализирует логгер с заданным уровнем.
func NewLogger(level string) (*zap.Logger, error) {
	var zapLevel zapcore.Level

	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// Создаём конфигурацию логгера
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		Development:      false,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "time",
			LevelKey:     "level",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	// Инициализируем логгер
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
