package commands

import (
	"fmt"
	"os"
	"todo/internal/db"

	"github.com/spf13/cobra"
)

var database db.TodoDBManager
var DBPath string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo cli app",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		database = db.NewSQLiteDB(db.DBConfig{Path: DBPath})
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
