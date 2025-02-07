package main

import (
	"context"
	"interceptor/gen/interceptor"
	"interceptor/internal/lib/logger"
	"interceptor/internal/server"
	"interceptor/internal/storage/sqlite"
	"log/slog"
	"net/http"
)

func main() {
	// logging
	logger := logger.SetupLogger()
	logger.Info("Starting server", slog.Any("address", "localhost"), slog.Any("port", 80))

	// database.db
	storage, err := sqlite.New("storage\\database.db")
	if err != nil {
		panic(err)
	}

	handlers := server.NewHandler(logger, storage)
	srv, err := interceptor.NewServer(handlers)

	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", RequestToContext(srv))

	if err := http.ListenAndServe("localhost:80", mux); err != nil {
		panic(err)
	}
}

// RequestToContext — пример middleware, который добавляет что-то в контекст запроса.
func RequestToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Здесь можно добавить логику, например, добавление данных в контекст.
		ctx := context.WithValue(r.Context(), "req", r)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
