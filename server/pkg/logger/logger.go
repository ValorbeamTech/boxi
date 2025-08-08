package logger

import (
    "log/slog"
    "os"
    "time"
)

type Logger interface {
    Info(msg string, args ...interface{})
    Error(msg string, err error, args ...interface{})
    Debug(msg string, args ...interface{})
    Warn(msg string, args ...interface{})
    Fatal(msg string, err error)
    With(key string, value interface{}) Logger
}

type slogger struct {
    logger *slog.Logger
}

func New(level string) Logger {
    var logLevel slog.Level
    switch level {
    case "debug":
        logLevel = slog.LevelDebug
    case "warn":
        logLevel = slog.LevelWarn
    case "error":
        logLevel = slog.LevelError
    default:
        logLevel = slog.LevelInfo
    }

    handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: logLevel,
        AddSource: true,
        ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
            if a.Key == slog.TimeKey {
                a.Value = slog.StringValue(time.Now().Format(time.RFC3339))
            }
            return a
        },
    })

    return &slogger{
        logger: slog.New(handler),
    }
}

func (l *slogger) Info(msg string, args ...interface{}) {
    l.logger.Info(msg, args...)
}

func (l *slogger) Error(msg string, err error, args ...interface{}) {
    if err != nil {
        args = append(args, "error", err.Error())
    }
    l.logger.Error(msg, args...)
}

func (l *slogger) Debug(msg string, args ...interface{}) {
    l.logger.Debug(msg, args...)
}

func (l *slogger) Warn(msg string, args ...interface{}) {
    l.logger.Warn(msg, args...)
}

func (l *slogger) Fatal(msg string, err error) {
    l.Error(msg, err)
    os.Exit(1)
}

func (l *slogger) With(key string, value interface{}) Logger {
    return &slogger{
        logger: l.logger.With(key, value),
    }
}