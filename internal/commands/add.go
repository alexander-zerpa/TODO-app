package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "add new todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adding stuff")
	},
}
