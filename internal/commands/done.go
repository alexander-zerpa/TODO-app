package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(done)
}

var done = &cobra.Command{
	Use:   "done",
	Short: "mark todo as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("this is done %v", args[0])
	},
}
