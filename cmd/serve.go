package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/catadoo/daily-mini-games/pkg/constants"
	"github.com/catadoo/daily-mini-games/pkg/container"
	"github.com/catadoo/daily-mini-games/pkg/core"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	shutdownTimeout = 25 * time.Second
)

var (
	ErrContextValueNotFound = errors.New("context value not found")
)

func Serve() *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "serve starts REST API Server for daily mini games",
		RunE:  serve,
	}

	return serveCmd
}

func serve(cmd *cobra.Command, _ []string) error {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cont, ok := ctx.Value(constants.ContextKeyContainer).(*container.Container)
	if !ok {
		log.Error().Err(ErrContextValueNotFound).Any("key", constants.ContextKeyContainer).Send()

		return fmt.Errorf("%w. key: %s", ErrContextValueNotFound, constants.ContextKeyContainer)
	}

	app := core.New(cont)
	app.RegisterRoutes()

	errCh := make(chan error)

	go func() {
		if err := app.Listen(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		// Done channel is closed when SIGTERM signal is received
		// Instead of stopping server right away, we can do a graceful shutdown to close resources and connections
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := app.Close(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("error while shutting down server")

			return err
		}

		log.Info().Msg("server shutdown completed successfully")
	case err := <-errCh:
		if err != nil {
			log.Error().Err(err).Msg("server listen error")
		}

		log.Info().Msg("server listen ended without errors")
	}

	return nil
}
