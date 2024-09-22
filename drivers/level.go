package drivers

// LogLevel defines the level of logging
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

func levelToString(level LogLevel) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
