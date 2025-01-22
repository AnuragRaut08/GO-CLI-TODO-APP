func main() {
	// Initialize an empty Todos slice
	todos := &Todos{}

	// Add some todos
	todos.add("Buy milk")
	todos.add("Buy bread")

	// Print the todos after adding items
	fmt.Println("Todos after adding items:")
	for i, todo := range *todos {
		fmt.Printf("%d: %+v\n", i, todo)
	}
	fmt.Println()

	// Delete the first todo
	fmt.Println("Deleting the first todo:")
	if err := todos.delete(0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Todo deleted successfully!")
	}

	// Print the todos after deletion
	fmt.Println("\nTodos after deletion:")
	for i, todo := range *todos {
		fmt.Printf("%d: %+v\n", i, todo)
	}

	package main

	func main() {
	todos := Todos{}
	
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	
	storage.Save(todos)
	}
}
