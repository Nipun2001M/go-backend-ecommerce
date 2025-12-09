package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Nipun2001M/go-backend-ecommerce/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config{
		address: ":8080",
		db: dbConfig{
			dsn: env.GetString(
				"GOOSE_DBSTRING",
				"host=localhost port=5435 user=adminNipungo password=gogogo dbname=ecom sslmode=disable",
			),
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	slog.Info("connected to database sucessfully!")

	api := application{
		config: cfg,
	}

	h := api.mount()
	err = api.run(h)
	if err != nil {
		slog.Error("Server Failed to start", "error", err)
		os.Exit(1)
	}

}
