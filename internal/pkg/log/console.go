package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type console struct {
	logger *logrus.Logger
	name   string
}

func NewConsole(name string) Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: false})
	logger.SetOutput(logrus.StandardLogger().Out)

	return &console{logger: logger, name: name}
}

func (service *console) concat(messages []any) []any {
	nameFormatted := fmt.Sprintf("[%v] ", service.name)
	info := []interface{}{nameFormatted}

	return append(info, messages...)
}

func (service *console) Info(messages ...any) {
	service.logger.Info(service.concat(messages)...)
}

func (service *console) Warn(messages ...any) {
	service.logger.Warn(service.concat(messages)...)
}

func (service *console) Error(messages ...any) {
	service.logger.Error(service.concat(messages)...)
}

func (service *console) Fatal(messages ...any) {
	service.logger.Fatal(service.concat(messages)...)
}

func (service *console) Debug(messages ...any) {
	service.logger.Debug(service.concat(messages)...)
}
