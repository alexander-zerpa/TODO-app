package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(done)
}

var done = &cobra.Command{
	Use:   "done",
	Short: "mark todo as done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is done")
	},
}
