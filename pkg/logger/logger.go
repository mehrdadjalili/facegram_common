package logger

type Logger interface {
	Info(string, map[string]interface{})
	Warning(string, map[string]interface{})
	Error(string, map[string]interface{})
	Sync()
}
