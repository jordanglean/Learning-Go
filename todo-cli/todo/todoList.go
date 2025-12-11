package todo

import (
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
}
