package commands

import (
	"todo/internal/db"

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
		database := db.NewSQLiteDB(db.DefaultDBConfig)

		if All {
			data, err := database.ListAll()
			if err != nil {
				cmd.Printf("Error: %v\n", err)
				return
			}
			cmd.Printf("All todos: %+v\n", data)
		} else if Done {
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
