package main

import (
	"cuan-tracker/internal/app/server"

	"github.com/spf13/cobra"
)

func newCLI() *cobra.Command {
	cli := &cobra.Command{
		Use:   "cuan-tracker",
		Short: "cuan-tracker",
	}
	cli.AddCommand(newServerCmd())
	cli.AddCommand(newMigrateCmd())
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

func newMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "migrate",
		Short:   "start migration",
		Aliases: []string{"server", "s"},
		Run: func(_ *cobra.Command, _ []string) {
			// cfg := config.GetConfig("./config.yml")
		},
	}
}
