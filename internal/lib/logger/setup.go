package logger

import (
	"interceptor/internal/domain"
	"interceptor/internal/lib/logger/handlers/slogpretty"
	"interceptor/internal/pkg/directories"

	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	var log *slog.Logger
	log = setupPrettySlog()

	//switch env {
	//case envLocal:
	//	log = setupPrettySlog()
	//case envDev:
	//	log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	//		AddSource: true,
	//		Level:     slog.LevelDebug,
	//	}))
	//case envProd:
	//
	//	log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	//		Level: slog.LevelInfo,
	//	}))
	//}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	// Создание файла для логов
	file, err := os.OpenFile(
		directories.FindDirectoryName(
			"logger",
		)+
			"\\logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
	)
	if err != nil {
		panic(err)
	}
	//defer file.Close()
	customWriter := &domain.CustomFileWriter{File: file}
	handler := opts.NewPrettyHandler(os.Stdout, customWriter)

	return slog.New(handler)
}
