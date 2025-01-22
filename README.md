# GO CLI TODO App

A simple Command Line Interface (CLI) application written in Go to manage a list of TODOs. The app supports saving and loading tasks from a JSON file for persistent storage.

## Features

- Add tasks to a TODO list.
- Toggle tasks as "Done" or "Pending."
- Display the list of tasks with their status.
- Save tasks to a JSON file for persistent storage.
- Load tasks from the JSON file on startup.

## Prerequisites

Make sure you have the following installed:

1. **Go Programming Language**:
   - [Install Go](https://go.dev/dl/)
   - Check installation:
     ```bash
     go version
     ```

2. **Git** (optional for cloning the repository):
   - [Install Git](https://git-scm.com/)
   - Check installation:
     ```bash
     git --version
     ```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AnuragRaut08/GO-CLI-TODO-APP.git
   cd GO-CLI-TODO-APP
Initialize dependencies (if required):

bash
Copy
Edit
go mod init cli-todo-app
go mod tidy
Build the application (optional):

bash
Copy
Edit
go build -o todo-app
How to Run
Run the application directly:

bash
Copy
Edit
go run main.go todos.go storage.go
Or, run the built binary:

bash
Copy
Edit
./todo-app
Usage
The program will:
Load existing TODOs from todos.json (if the file exists).
Allow you to add tasks to the TODO list.
Enable toggling the status of tasks.
Display the list of tasks with their status.
Save updated tasks to todos.json.
Example Output:
plaintext
Copy
Edit
1. Buy Milk [Done]
2. Buy Bread [Pending]
File Structure
plaintext
Copy
Edit
├── main.go          # Entry point for the application
├── todos.go         # Implementation of Todos struct and methods
├── storage.go       # Generic JSON storage logic
├── todos.json       # JSON file where tasks are saved (auto-created)
├── README.md        # Documentation
Notes
The todos.json file is automatically created in the current working directory to store tasks persistently.
If errors occur, ensure proper read/write permissions for the todos.json file.
Contributing
Contributions are welcome!
Feel free to fork this repository, create a branch, and submit a pull request with your improvements or new features.

License
This project is licensed under the MIT License. See the LICENSE file for details.

perl
Copy
Edit

### Steps to Add This to GitHub

1. Create a `README.md` file in your project directory.
   ```bash
   touch README.md
Open the file in an editor (e.g., VS Code, nano, or vim) and paste the content above.

Stage, commit, and push the changes:

bash
Copy
Edit
git add README.md
git commit -m "Add README file"
git push origin main
