package termination

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"pt-test/pkg/app/logger"
)

type (
	Waiter struct {
		logger *logger.Logger
	}

	WaiterOption func(*Waiter)
)

func NewWaiter(options ...WaiterOption) *Waiter {
	w := &Waiter{
		logger: logger.NewNop(),
	}
	for _, opt := range options {
		opt(w)
	}

	return w
}

func WaiterWithLogger(logger *logger.Logger) WaiterOption {
	return func(waiter *Waiter) {
		waiter.logger = logger.Named(waiter.Name())
	}
}

var ErrStopped = errors.New("application is stopped")

func (w *Waiter) Runner(ctx context.Context) func() error {
	return func() error {
		signalsChan := make(chan os.Signal, 1)
		signal.Notify(signalsChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-signalsChan:
			w.logger.Info("got termination signal", zap.String("signal", sig.String()))

			return ErrStopped

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *Waiter) Name() string {
	return "TerminationWaiter"
}
