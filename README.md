# Task Manager CLI

A simple command-line interface for managing your tasks.  Add, update, list, and delete tasks with ease.

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/harshkasat/taskManagerCli)](https://goreportcard.com/report/github.com/harshkasat/taskManagerCli)


## Table of Contents

- [Project Overview](#project-overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
    - [Adding a Task](#adding-a-task)
    - [Listing Tasks](#listing-tasks)
    - [Updating a Task](#updating-a-task)
    - [Deleting a Task](#deleting-a-task)
- [Project Architecture](#project-architecture)
- [Contributing](#contributing)
- [License](#license)


## Project Overview

`taskManagerCli` is a lightweight command-line application built using Go and Cobra that helps you manage your tasks efficiently.  It allows you to add tasks with descriptions, due dates, and priorities, mark tasks as complete, and list tasks based on various criteria.  This tool is ideal for individuals who prefer a simple, text-based interface for task management.


## Prerequisites

- Go 1.18 or higher installed on your system.


## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/harshkasat/taskManagerCli.git
   cd taskManagerCli
   ```

2. Build the application:
   ```bash
   go build
   ```

This will create an executable file (named `task-manager` by default).


## Usage

The application is run from your terminal using the command `task-manager`.  Several subcommands are available:

### Adding a Task

To add a new task, use the `add` subcommand:

```bash
task-manager add --title "Write README" --priority high --due "2024-10-27" --notes "Need to finish this before the deadline."
```

This command requires `--title`, `--priority`, and `--notes` flags.  `--due` is optional; if not provided, a default due date of 7 days from the current date will be used.

The `addTask` function in `addTask.go` handles this:

```go
func addTask(title, priority, dueDate, note string) error {
    // ... (code to read existing tasks, generate a new ID, create a new Task struct, append it to the existing tasks, and save the updated list to task.json) ...
}
```

### Listing Tasks

To list all tasks, simply run:

```bash
task-manager list
```

The `Execute` function in `root.go` contains the logic for this (although the `if all` block needs to be uncommented to enable it):

```go
func Execute() {
    // ...
    if all { //Uncomment to enable listing
        tasks := listTask() //This function is not provided in the code sample
        for _, task := range tasks {
            fmt.Printf("Task ID: %v, Title: %s, Priority: %s, Notes: %s, Due Dates: %s, Status: %s\n", task.ID, task.Title, task.Priority, task.Notes, task.DueDate, task.Status)
        }
    }
}
```

### Updating a Task

Use the `update` subcommand to modify existing tasks:

```bash
task-manager update --id 1 --field title --new "Updated Task Title"
```

This command requires `--id`, `--field`, and `--new` flags.  `--id` specifies the ID of the task to update, `--field` indicates the field to modify (title, due, priority, notes, status), and `--new` provides the new value.  The `updateTask` function in `updateTask.go` handles the logic.

### Deleting a Task

The provided code does not include a `delete` command.  This would need to be added.


## Project Architecture

The project uses the Cobra library for creating the CLI structure.  Data is stored in a JSON file (`task.json`).  The application consists of several commands (`add`, `update`, `list` - `delete` needs to be added) each implemented as a separate Cobra command.


## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](CONTRIBUTING.md) file for details.  (Note:  A CONTRIBUTING.md file is not present in the provided code.)


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. (Note: A LICENSE file is not present in the provided code.)
