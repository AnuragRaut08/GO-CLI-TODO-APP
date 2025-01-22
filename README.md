
# Go CLI TODO App

A simple Command Line Interface (CLI) application written in Go to manage a list of TODOs. The app supports saving and loading tasks from a JSON file for persistent storage.

## Features

- Add tasks to a TODO list.
- Toggle tasks as "Done" or "Pending."
- Display the list of tasks with their status.
- Save tasks to a JSON file for persistent storage.
- Load tasks from the JSON file on startup.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AnuragRaut08/GO-CLI-TODO-APP.git
   cd GO-CLI-TODO-APP
   ```

2. Initialize dependencies (if required):
   ```bash
   go mod init cli-todo-app
   go mod tidy
   ```

3. Build the application (optional):
   ```bash
   go build -o todo-app
   ```

## How to Run

1. Run the application directly:
   ```bash
   go run main.go todos.go storage.go
   ```

2. Or, run the built binary:
   ```bash
   ./todo-app
   ```

## Usage

The program will:
1. Load existing TODOs from `todos.json` (if the file exists).
2. Allow you to add tasks to the TODO list.
3. Enable toggling the status of tasks.
4. Display the list of tasks with their status.
5. Save updated tasks to `todos.json`.

### Example Output
```plaintext
1. Buy Milk [Done]
2. Buy Bread [Pending]
```

## File Structure

```plaintext
├── main.go          # Entry point for the application
├── todos.go         # Implementation of Todos struct and methods
├── storage.go       # Generic JSON storage logic
├── todos.json       # JSON file where tasks are saved (auto-created)
├── README.md        # Documentation
```

## Notes

- The `todos.json` file is automatically created in the current working directory to store tasks persistently.
- If errors occur, ensure proper read/write permissions for the `todos.json` file.

## Contributing

Contributions are welcome!  
Feel free to fork this repository, create a branch, and submit a pull request with your improvements or new features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Steps to Add This to GitHub

1. Create a `README.md` file in your project directory:
   ```bash
   touch README.md
   ```

2. Open the file in an editor (e.g., VS Code, nano, or vim) and paste the content above.

3. Stage, commit, and push the changes to your repository:
   ```bash
   git add README.md
   git commit -m "Add README file"
   git push origin main
   ```
```
