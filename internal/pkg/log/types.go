package log

type Logger interface {
	Info(messages ...any)
	Warn(messages ...any)
	Error(messages ...any)
	Fatal(messages ...any)
	Debug(messages ...any)
}
