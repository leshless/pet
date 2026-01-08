package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/leshless/golibrary/graceful"
	"github.com/leshless/pet/cub/internal/config"
	"github.com/leshless/pet/cub/internal/environment"
	"github.com/leshless/pet/cub/internal/telemetry"
)

func InitClient(
	logger telemetry.Logger,
	tel telemetry.Telemetry,
	gracefulRegistrator graceful.Registrator,
	configHolder config.Holder,
	environmentHolder environment.Holder,
) (*client, error) {
	config := configHolder.Config().DB
	dbPassword := environmentHolder.Environment().DBPassword

	logger.Info(context.Background(), "initializing db client")

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		dbPassword,
		config.Database,
		config.SSLMode,
	)

	dbConn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		logger.Error(context.Background(), "failed to connect to db", telemetry.Error(err))
		return nil, fmt.Errorf("connecting to db: %w", err)
	}

	client := newClient(tel, dbConn)

	gracefulRegistrator.Register(client.Close)

	err = client.Ping(context.Background())
	if err != nil {
		logger.Error(context.Background(), "failed to ping db", telemetry.Error(err))
		return nil, fmt.Errorf("pinging db: %w", err)
	}

	logger.Info(context.Background(), "db client successfully initialized")

	return client, nil
}
