package cmd

import (
	"github.com/sgash708/nde-clock-io/internal/application"
	"github.com/spf13/cobra"
)

var clockoutCmd = &cobra.Command{
	Use:   "clockout",
	Short: "clock out in a certain site",
	Run: func(cmd *cobra.Command, args []string) {
		if err := application.RunClockOut(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clockoutCmd)
}
