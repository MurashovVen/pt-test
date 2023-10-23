package controller

import (
	"context"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"pt-test/pkg/app"
)

var _ app.Work = (*Server)(nil)

func (s *Server) Runner(ctx context.Context) func() error {
	return func() error {
		eg, egCtx := errgroup.WithContext(ctx)

		eg.Go(func() error {
			return s.server.ListenAndServe()
		})

		eg.Go(func() error {
			<-egCtx.Done()

			shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer shutdownCtxCancel()

			if err := s.server.Shutdown(shutdownCtx); err != nil {
				s.logger.Error("shutdown error", zap.Error(err))

				return err
			}

			s.logger.Debug("shutdown successfully completed")

			return nil
		})

		return eg.Wait()
	}
}

func (s *Server) Name() string {
	return "HTTPServer"
}
