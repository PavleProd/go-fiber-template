package cmd

import (
	"context"
	"github.com/catadoo/daily-mini-games/pkg/config"
	"github.com/catadoo/daily-mini-games/pkg/constants"
	"github.com/catadoo/daily-mini-games/pkg/container"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func addSubcommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(Serve())
}

func root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:                "DMG-API",
		Short:              "DMG-API is a REST API for mobile application daily mini games",
		PersistentPreRunE:  initCommand,    // children will also execute this command
		PersistentPostRunE: cleanupCommand, // children will also execute this command
	}

	addSubcommands(rootCmd)

	return rootCmd
}

func initCommand(cmd *cobra.Command, _ []string) error {
	cfg, err := config.New()
	if err != nil {
		log.Error().Err(err).Msg("failed to create config")

		return err
	}

	cnt := container.New(cfg)

	ctx := cmd.Context()
	ctx = context.WithValue(ctx, constants.ContextKeyContainer, cnt)
	cmd.SetContext(ctx)

	return nil
}

func cleanupCommand(_ *cobra.Command, _ []string) error {
	// TODO: add any global cleanup logic here
	return nil
}

func Execute() {
	if err := root().Execute(); err != nil {
		log.Error().Err(err).Msg("failed to start CLI command")
	}
}
