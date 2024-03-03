package logger

import (
	"context"

	"go.uber.org/zap"
)

type Fields map[string]any

type Logger interface {
	Info(ctx context.Context, msg string, fields Fields)
	Error(ctx context.Context, msg string, fields Fields)
	Sync()
}

type logger struct {
	logger *zap.Logger
}

func InitLogger() Logger {
	zapLogger, _ := zap.NewProduction()
	return &logger{
		logger: zapLogger,
	}
}

func (l *logger) Sync() {
	_ = l.logger.Sync()
}

func (l *logger) Info(ctx context.Context, msg string, fields Fields) {
	zapFields := l.getFields(ctx, fields)
	l.logger.Info(msg, zapFields...)
}

func (l *logger) Error(ctx context.Context, msg string, fields Fields) {
	zapFields := l.getFields(ctx, fields)
	l.logger.Error(msg, zapFields...)
}

func (l *logger) getFields(ctx context.Context, fields Fields) []zap.Field {
	var zapFields []zap.Field

	if fields == nil {
		return zapFields
	}

	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}

	return zapFields
}
