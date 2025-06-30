package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	// log level - debug, info, warning, error, fatal
	Level string `yaml:"level" default:"info" json:"level,omitempty"`

	// format strategy - plain, json
	Format string `yaml:"format" default:"json" json:"format,omitempty"`

	// file out location - "./log.json"
	WFilePath string `yaml:"w_file_path" default:"json" json:"wFilePath,omitempty"`

	// audit system events - none(default), stdout, db
	AuditEvents string `yaml:"audit_events" default:"none" json:"audit_events,omitempty"`

	// IgnoredAuditEvents contains list of events which should be ignored in audit logs
	IgnoredAuditEvents []string `yaml:"ignored_audit_events" json:"ignored_audit_events,omitempty"` 
}

type Logger = *zap.Logger

var logger *zap.Logger

func Set(cfg Config) Logger {
	zapcfg := zap.NewProductionEncoderConfig()
	zapcfg.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(zapcfg)
	logLevel := zap.NewAtomicLevelAt(atomicLevel(cfg.Level))

	var core zapcore.Core
	if (strings.Compare(cfg.WFilePath, "") != 0) {
		fileEncoder := zapcore.NewJSONEncoder(zapcfg)
		logFile, _ := outFile(cfg.WFilePath)

		writer := zapcore.AddSync(logFile)

		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, logLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		)
	}

	logger = zap.New(core, zap.AddCaller())

	return logger
}

func ZapInt(key string, val int) zap.Field {
	return zap.Int(key, val)
} 

func ZapString(key string, val string) zap.Field {
	return zap.String(key, val)
}

func ZapTime(key string, val time.Time) zap.Field {
	return zap.Time(key, val)
}

func WrapGRPCCtx(ctx context.Context) context.Context {
	return grpczap.ToContext(ctx, logger)
}
func UnWrapGRPCCtx(ctx context.Context) *zap.Logger {
	return grpczap.Extract(ctx)
}

func atomicLevel(level string) zapcore.Level {
	switch level {
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func outFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if (err != nil) {
		log.Fatal(fmt.Printf("couldn't open file for logger: %s", err))
	}

	return logFile, nil
}
