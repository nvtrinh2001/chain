package yoda

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "config [key] [value]",
		Aliases: []string{"c"},
		Short:   "Set yoda configuration environment",
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			values := args[1:]

			for _, value := range values {
				viper.Set(args[0], value)
			}
			return viper.WriteConfig()
		},
	}
	return cmd
}
