package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var All bool
var Done bool

func init() {
	list.Flags().BoolVarP(&All, "all", "", false, "list all todos")
	list.Flags().BoolVarP(&Done, "done", "", false, "list done todos")
	rootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "list saved todos",
	Run: func(cmd *cobra.Command, args []string) {
		if All {
			fmt.Println("list all things")
		} else if Done {
			fmt.Println("list done things")
		} else {
			fmt.Println("list things")
		}
	},
}
