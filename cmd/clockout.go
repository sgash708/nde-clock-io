package cmd

import (
	"github.com/sgash708/nde-clock-io/internal/application"
	"github.com/spf13/cobra"
)

var clockoutCmd = &cobra.Command{
	Use:   "clockout",
	Short: "clock out",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := application.RunClockOut(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clockoutCmd)
}
