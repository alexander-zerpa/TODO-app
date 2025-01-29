package commands

import (
	"github.com/spf13/cobra"
	"todo/internal/db"
	"todo/internal/models"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add Shortname Title Description",
	Short: "add new todo",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			cmd.Help()
			return
		}

		todo := models.Todo{ID: args[0], Title: args[1], Description: args[2]}

		database := db.NewSQLiteDB(db.DefaultDBConfig)
		err := database.Add(&todo)
		if err != nil {
			cmd.Printf("Error: %v\n", err)
		}

		cmd.Printf("Todo added: %+v\n", todo)
	},
}
