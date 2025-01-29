package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(done)
}

var done = &cobra.Command{
	Use:   "done shortname",
	Short: "mark todo as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := database.Get(args[0])
		if err != nil {
			cmd.Printf("Error: %v\n", err)
			return
		} else if len(todos) != 1 {
			cmd.Println("Error: todo not found")
			return
		}

		todo := todos[0]
		todo.Done = true
		err = database.Update(&todo)
		if err != nil {
			cmd.Printf("Error: %v\n", err)
			return
		}
		cmd.Printf("Todo marked as done: %+v\n", todo)
	},
}
