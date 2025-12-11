package todo

import (
	"time"
)

type Todo struct {
	id        int
	content   string
	createdAt time.Time
	dueAt     string
	creator   string
}

func NewTodo(id int, content string, dueAt string, creator string) Todo {
	return Todo{
		id:        id,
		content:   content,
		dueAt:     dueAt,
		createdAt: time.Now(),
		creator:   creator,
	}
}

func (todo *Todo) Edit(content string, dueAt string) {
	if content != "" {
		todo.content = content
	}

	if dueAt != "" {
		todo.dueAt = dueAt
	}
}
