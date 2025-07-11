package utils

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	// Debugf(template string, args ...interface{})
	// Debugw(template string, keysAndValues ...interface{})
	// Infof(template string, args ...interface{})
	// Infow(template string, keysAndValues ...interface{})
	// Warnf(template string, args ...interface{})
	// Warnw(template string, keysAndValues ...interface{})
	// Errorf(template string, args ...interface{})
	// Errorw(template string, keysAndValues ...interface{})
	// Panicf(template string, args ...interface{})
	// Panicw(template string, keysAndValues ...interface{})

	Debug(string, ...zap.Field)
	DebugCtx(context.Context, string, ...zap.Field)
	Info(string, ...zap.Field)
	InfoCtx(context.Context, string, ...zap.Field)
	Warn(string, ...zap.Field)
	WarnCtx(context.Context, string, ...zap.Field)
	Error(string, ...zap.Field)
	ErrorCtx(context.Context, string, ...zap.Field)
	Panic(string, ...zap.Field)
	PanicCtx(context.Context, string, ...zap.Field)
}

type logger struct {
	l *zap.Logger
	// s      *zap.SugaredLogger
}

var (
	devEncoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	prodEncoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
)

func NewLogger() Logger {
	var (
		level zapcore.Level
		core  zapcore.Core
	)

	level = zap.InfoLevel
	// switch cfg.Log.Level {
	// case "debug":
	// 	level = zap.DebugLevel
	// case "info":
	// 	level = zap.InfoLevel
	// case "warn":
	// 	level = zap.WarnLevel
	// default:
	// 	level = zap.InfoLevel
	// }

	// if !(cfg.Mode == "dev") {
	// 	// core = zapcore.NewCore(prodEncoder, zapcore.AddSync(&lumberjack.Logger{
	// 	// 	Filename:   cfg.Log.FileName,
	// 	// 	MaxBackups: cfg.Log.MaxBackups,
	// 	// 	MaxAge:     cfg.Log.MaxAge,
	// 	// },
	// 	// ), level)
	// 	core = zapcore.NewCore(prodEncoder, os.Stdout, level)
	// } else {
	// 	core = zapcore.NewCore(devEncoder, os.Stdout, level)
	// }
	core = zapcore.NewCore(devEncoder, os.Stdout, level)

	return &logger{
		// s: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(),
		l: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)),
	}
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *logger) DebugCtx(ctx context.Context, msg string, fields ...zap.Field) {
	f := l.trace_extract(ctx, fields...)
	l.l.Debug(msg, f...)
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func (l *logger) InfoCtx(ctx context.Context, msg string, fields ...zap.Field) {
	f := l.trace_extract(ctx, fields...)
	l.l.Info(msg, f...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}
func (l *logger) WarnCtx(ctx context.Context, msg string, fields ...zap.Field) {
	f := l.trace_extract(ctx, fields...)
	l.l.Warn(msg, f...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}
func (l *logger) ErrorCtx(ctx context.Context, msg string, fields ...zap.Field) {
	f := l.trace_extract(ctx, fields...)
	l.l.Error(msg, f...)
}

func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.l.Panic(msg, fields...)
}
func (l *logger) PanicCtx(ctx context.Context, msg string, fields ...zap.Field) {
	f := l.trace_extract(ctx, fields...)
	l.l.Panic(msg, f...)
}

func (l *logger) trace_extract(ctx context.Context, fields ...zap.Field) []zap.Field {
	// span := trace.SpanFromContext(ctx)
	// fields = append(fields,
	// 	zap.String("trace_id", span.SpanContext().TraceID().String()),
	// 	zap.String("span_id", span.SpanContext().SpanID().String()),
	// 	zap.String("trace_flags", span.SpanContext().TraceFlags().String()),
	// )
	return fields
}
