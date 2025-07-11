package logging

import "github.com/amirhasanpour/car-sale-management-wep-api/src/config"

type Logger interface {
	Init()

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any)
	Debugf(template string, args ...any)

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any)
	Infof(template string, args ...any)

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any)
	Warnf(template string, args ...any)

	Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any)
	Errorf(template string, args ...any)

	Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any)
	Fatalf(template string, args ...any)
}

func NewLogger(cfg *config.Config) Logger {
	switch cfg.Logger.Logger {
	case "zap":
		return newZapLogger(cfg)
	case "zerolog":
		return newZeroLogger(cfg)
	}
	panic("logger not supported")
}

// file <- filebeat -> elasticsearch -> kibana