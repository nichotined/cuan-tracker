package main

import (
	"cuan-tracker/internal/app"
	"cuan-tracker/internal/app/server"

	"github.com/spf13/cobra"
)

func newCLI() *cobra.Command {
	cli := &cobra.Command{
		Use:   "cuan-tracker",
		Short: "cuan-tracker",
	}
	cli.AddCommand(newServerCmd())
	cli.AddCommand(newMigrateUpCmd())
	cli.AddCommand(newMigrateDownCmd())
	return cli
}

func newServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "server",
		Short:   "start cuan-tracker",
		Aliases: []string{"server", "s"},
		Run: func(_ *cobra.Command, _ []string) {
			server.Start()
		},
	}
}

func newMigrateUpCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "migrate_up",
		Short:   "start migration up",
		Aliases: []string{"server", "s"},
		Run: func(_ *cobra.Command, _ []string) {
			app.MigrateDBUp()
		},
	}
}

func newMigrateDownCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "migrate_down",
		Short:   "start migration down",
		Aliases: []string{"server", "s"},
		Run: func(_ *cobra.Command, _ []string) {
			app.MigrateDBDown()
		},
	}
}
