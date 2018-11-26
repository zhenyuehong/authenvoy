package appcode

type AppCode int

const (
	BadData       AppCode = 400
	LoggingErr    AppCode = 500
	ValidationErr AppCode = 501
)