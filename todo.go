package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// Todo represents a single todo item
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// Todos represents a list of Todo items
type Todos []Todo

// add adds a new todo to the list
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

// validateIndex ensures the given index is valid
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("invalid index")
	}
	return nil
}

// delete removes a todo by index
func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

// print displays the todos in a formatted table
func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for i, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
