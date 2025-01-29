package main

import (
	"todo/internal/commands"
)

func main() {
	commands.DBPath = "todo.db"
	commands.Execute()
}
