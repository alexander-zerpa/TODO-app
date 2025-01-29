package commands

import (
	"github.com/spf13/cobra"
)

var listAll bool
var listDone bool

func init() {
	list.Flags().BoolVarP(&listAll, "all", "", false, "list all todos")
	list.Flags().BoolVarP(&listDone, "done", "", false, "list done todos")
	rootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "list saved todos",
	Run: func(cmd *cobra.Command, args []string) {
		if listAll {
			data, err := database.ListAll()
			if err != nil {
				cmd.Printf("Error: %v\n", err)
				return
			}
			cmd.Printf("All todos: %+v\n", data)
		} else if listDone {
			data, err := database.List(true)
			if err != nil {
				cmd.Printf("Error: %v\n", err)
				return
			}
			cmd.Printf("Done todos: %+v\n", data)
		} else {
			data, err := database.List(false)
			if err != nil {
				cmd.Printf("Error: %v\n", err)
				return
			}
			cmd.Printf("Pending todos: %+v\n", data)
		}
	},
}
