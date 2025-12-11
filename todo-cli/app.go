package main

import (
	"fmt"

	"example.com/todo-cli/input"
	"example.com/todo-cli/todo"
)

func main() {
	list := todo.TodoList{}

	for {
		fmt.Println("\n\n== TODO CLI ==")
		fmt.Println("What would you like to do? ")
		fmt.Printf("1. Add todo\n2. View todo list\n3. Edit todo\n4. Delete todo\n5. Exit\n")

		userOption := input.Get(">> ")

		switch userOption {
		case "1":
			list.Add()
		case "2":
			list.ViewList()
		case "5":
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}
