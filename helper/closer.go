package helper

import (
	"io"
	"log/slog"
)

// CloseAndLogError尝试关闭资源并记录任何错误。
func CloseAndLogError(closer io.Closer, message string) {
	CloseAndLogErrorWithLogger(closer, message, slog.Default())
}

// CloseAndLogErrorWithLogger尝试关闭资源，并将任何错误记录到给定的记录器。
func CloseAndLogErrorWithLogger(closer io.Closer, message string, logger *slog.Logger) {
	err := closer.Close()
	if err != nil {
		logger.Error(message, "error", err)
	}
}
