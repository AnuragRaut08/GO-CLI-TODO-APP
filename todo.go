package main

import (
	"errors"
	"fmt"
	"time"
)

// Todo struct defines a task
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// Todos is a slice of Todo items
type Todos []Todo

// ValidateIndex ensures the given index is valid for the Todos slice
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

// Delete removes a todo at the specified index
func (todos *Todos) delete(index int) error {
	t:=*todos
	// Validate the index first
	if err := t.validateIndex(index); err != nil {
		return err
	}

	// Remove the item at the index
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}