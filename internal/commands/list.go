package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "list saved todos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stuff to do")
	},
}
