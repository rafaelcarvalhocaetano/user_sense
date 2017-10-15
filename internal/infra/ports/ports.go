package ports

type Logger interface {
	Info(msg string, payload any, sttCode int)
	Debug(msg string, payload any, sttCode int)
	Error(msg string, payload any, sttCode int)
	Warn(msg string, payload any, sttCode int)
}
