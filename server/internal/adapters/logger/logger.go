package logger

import (
	"log/slog"
	"os"

	"github.com/server/internal/adapters/env"
)

var logger *slog.Logger

func Set(enviroment env.ENVIROMENT) {

	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	if env.Enviroment[enviroment] == "development" {
		logger = slog.New(
			slog.NewJSONHandler(os.Stderr, nil),
		)
	}

	slog.SetDefault(logger)
}
