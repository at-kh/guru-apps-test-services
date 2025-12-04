package app

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type customEncoder struct{}

func (e *customEncoder) encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339Nano))
}

func (e *customEncoder) encodeDuration(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(d.String())
}

// InitLogger initializes logger for application.
func (a *App) initLogger() {
	ce := customEncoder{}

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			EncodeTime:     ce.encodeTime,
			EncodeDuration: ce.encodeDuration,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logger, err := cfg.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.DPanicLevel),
	)
	if err != nil {
		log.Fatalf("failed to build logger: %v", err)
	}

	a.logger = logger

	a.logger.Info("logger initialized",
		zap.String("app", a.meta.Info.Name),
		zap.String("version", a.meta.Info.BuildVersion),
	)
}
