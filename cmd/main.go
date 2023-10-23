package main

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"pt-test/internal/controller"
	v1 "pt-test/internal/controller/v1"
	"pt-test/internal/repo"
	"pt-test/internal/service"
	"pt-test/pkg/app"
	"pt-test/pkg/app/configuration"
	"pt-test/pkg/app/logger"
	"pt-test/pkg/app/termination"
	"pt-test/pkg/postgres"
)

func main() {
	cfg := new(config)
	configuration.MustProcessConfig(cfg)

	log := logger.MustCreateLogger(cfg.Env)

	db := postgres.MustConnect(cfg.DB)
	defer func() { _ = db.Close() }()

	application := app.New(
		log,
		app.AppendWorks(
			controller.New(
				v1.NewRouter(
					service.New(
						repo.New(db),
					),
					log,
				),
				controller.WithAddress(cfg.HTTPServerAddress),
				controller.WithEnvironment(cfg.Env),
				controller.WithLogger(log),
			),
		),
	)

	if err := application.Run(context.Background()); err != nil && !errors.Is(err, termination.ErrStopped) {
		log.Error("running error", zap.Error(err))
	}
}
