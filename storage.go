package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Storage[T any] struct {
	FileName string
}

// NewStorage creates a new storage instance
func NewStorage[T any](fileName string) (*Storage[T], error) {
	if fileName == "" {
		return nil, errors.New("fileName cannot be empty")
	}
	return &Storage[T]{FileName: fileName}, nil
}

// Save saves data to the file in JSON format
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile(s.FileName, fileData, 0600)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", s.FileName, err)
	}
	return nil
}

// Load loads data from the file into the provided reference
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", s.FileName, err)
	}

	err = json.Unmarshal(fileData, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return nil
}

func main() {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	storage, err := NewStorage[User]("user_data.json")
	if err != nil {
		fmt.Println("Error creating storage:", err)
		return
	}

	// Example data
	user := User{Name: "John Doe", Email: "john@example.com", Age: 30}

	// Save data
	err = storage.Save(user)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	fmt.Println("Data saved successfully.")

	// Load data
	var loadedUser User
	err = storage.Load(&loadedUser)
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}
	fmt.Printf("Data loaded successfully: %+v\n", loadedUser)
}
