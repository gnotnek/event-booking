package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Execute() {
	var command = &cobra.Command{
		Use:   "event-booking",
		Short: "Event booking application",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(apiCmd())

	if err := command.Execute(); err != nil {
		log.Fatal().Err(err).Msg("could not execute command")
	}
}
