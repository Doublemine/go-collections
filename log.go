package collections

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *LoggerType
)

const (
	LoggerEncodingConsole = "console"
	LoggerEncodingJSON    = "json"

	// LoggerLevelInfo Logger level info
	LoggerLevelInfo string = "info"
	// LoggerLevelDebug Logger level debug
	LoggerLevelDebug string = "debug"
	// LoggerLevelWarn Logger level warn
	LoggerLevelWarn string = "warn"
	// LoggerLevelError Logger level error
	LoggerLevelError string = "error"
	// LoggerLevelFatal Logger level fatal
	LoggerLevelFatal string = "fatal"
	// LoggerLevelPanic Logger level panic
	LoggerLevelPanic string = "panic"
)

func init() {
	var err error
	if Logger, err = NewLogger("info"); err != nil {
		panic(fmt.Sprintf("create logger: %+v", err))
	}
}

type LoggerType struct {
	*zap.Logger
	level zap.AtomicLevel
}

// ChangeLevel change logger level
func (l *LoggerType) ChangeLevel(level string) (err error) {
	switch level {
	case LoggerLevelDebug:
		l.level.SetLevel(zap.DebugLevel)
	case LoggerLevelInfo:
		l.level.SetLevel(zap.InfoLevel)
	case LoggerLevelWarn:
		l.level.SetLevel(zap.WarnLevel)
	case LoggerLevelError:
		l.level.SetLevel(zap.ErrorLevel)
	case LoggerLevelFatal:
		l.level.SetLevel(zap.FatalLevel)
	case LoggerLevelPanic:
		l.level.SetLevel(zap.PanicLevel)
	default:
		return fmt.Errorf("log level only be debug/info/warn/error/fatal/panic")
	}
	return
}

// NewLogger create new logger
func NewLogger(level string, opts ...zap.Option) (l *LoggerType, err error) {
	return NewLoggerWithName("", level, opts...)
}

// NewLoggerWithName create new logger with name
func NewLoggerWithName(name, level string, opts ...zap.Option) (l *LoggerType, err error) {
	zl := zap.NewAtomicLevel()
	cfg := zap.Config{
		Level:            zl,
		Development:      false,
		Encoding:         LoggerEncodingConsole,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapLogger, err := cfg.Build(opts...)
	if err != nil {
		return nil, fmt.Errorf("build zap logger: %+v", err)
	}
	zapLogger = zapLogger.Named(name)

	l = &LoggerType{
		Logger: zapLogger,
		level:  zl,
	}
	return l, l.ChangeLevel(level)
}
