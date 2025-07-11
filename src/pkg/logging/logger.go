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

	Error(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]any)
	Errorf(err error, template string, args ...any)

	Fatal(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]any)
	Fatalf(err error, template string, args ...any)
}

func NewLogger(cfg *config.Config) Logger {
	return nil
}

// file <- filebeat -> elasticsearch -> kibana