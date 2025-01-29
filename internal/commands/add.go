package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "add new todo",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("adding stuff %v, %v, %v", args[0], args[1], args[2])
	},
}
