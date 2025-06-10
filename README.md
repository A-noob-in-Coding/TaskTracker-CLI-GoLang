# Task Tracker CLI

A simple command-line interface for managing your tasks. Built with Go using only native libraries.

## Features

- Add, update, and delete tasks
- Mark tasks as in-progress or done
- List all tasks or filter by status
- Persistent storage using JSON file
- Automatic task ID generation and timestamps

## Installation

1. Clone the repository:
```bash
git clone [<repository-url>](https://github.com/A-noob-in-Coding/TaskTracker-CLI-GoLang.git)
cd TaskTracker
```

2. Build the application:
```bash
go build -o task-cli
```

## Usage

### Basic Commands

```bash
# Add a new task
./task-cli add "Buy groceries"

# Update a task
./task-cli update 1 "Buy groceries and cook dinner"

# Delete a task
./task-cli delete 1

# Mark task status
./task-cli mark-in-progress 1
./task-cli mark-done 1

# List all tasks
./task-cli list

# List tasks by status
./task-cli list done
./task-cli list todo  
./task-cli list inprogress

# Show help
./task-cli -h
```

## Task Properties

Each task contains:
- `id`: Unique identifier
- `description`: Task description
- `status`: `todo`, `in-progress`, or `done`
- `createdAt`: Creation timestamp
- `updatedAt`: Last modification timestamp

## Data Storage

Tasks are stored in a `tasks.json` file in the current directory. The file is created automatically if it doesn't exist.

## Project Structure

```
TaskTracker/
├── main.go
├── Utilities/
│   └── [utility functions]
└── tasks.json (created automatically)
```

## Requirements

- Go 1.16 or higher
- No external dependencies required
