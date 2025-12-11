package todo

import (
	"fmt"
	"os"
	"text/tabwriter"

	"example.com/todo-cli/input"
)

type TodoList struct {
	todos  []Todo
	nextID int
}

func (todoList *TodoList) Add() {
	id := todoList.nextID
	content := input.Get("Enter your todo item: ")
	dueAt := input.Get("When is this item due?: ")
	creator := input.Get("Your name: ")

	todoItem := NewTodo(id, content, dueAt, creator)

	todoList.todos = append(todoList.todos, todoItem)
	todoList.nextID++
	fmt.Println("TODO item has been added!")
}

func (todoList TodoList) ViewList() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Header
	fmt.Fprintln(w, "\n\nID\tTODO\tDUE\tCREATOR\tCREATED AT")
	fmt.Fprintln(w, "--\t-------\t---\t-------\t----------")

	for _, t := range todoList.todos {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			t.id, t.content, t.dueAt, t.creator, t.createdAt.Format("2006-01-02"))
	}

	w.Flush()
}
