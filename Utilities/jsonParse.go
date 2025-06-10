package Utilities

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const filename = "Utilities/data.json" // Define filename as a constant
func readTasks() ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // Return empty slice if file doesn't exist
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasks []Task
	if len(data) > 0 { // Only unmarshal if there's data
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
		}
	}
	return tasks, nil
}

// writeTasks writes tasks to the data file.
func writeTasks(tasks []Task) error {
	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	err = os.WriteFile(filename, updatedData, 0666)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func AddTaskToFile(newTask Task) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, newTask)
	return writeTasks(tasks)
}

func ListTasks(flag int) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	fmt.Printf("Found %d tasks:\n\n", len(tasks))
	for i, task := range tasks {
		// Filter tasks based on flag
		switch flag {
		case 1: // Done
			if task.Status != "Done" {
				continue
			}
		case 2: // In-Progress
			if task.Status != "In-Progress" {
				continue
			}
		case 3: // To Do
			if task.Status != "To Do" {
				continue
			}
		case -1: // All tasks, do nothing
		default:
			fmt.Println("Invalid flag, showing all tasks")
		}

		fmt.Printf("Task %d:\n", i+1)
		fmt.Printf("  ID: %d\n", task.ID)
		fmt.Printf("  Description: %s\n", task.Description)
		fmt.Printf("  Status: %s\n", task.Status)
		fmt.Printf("  Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("  Updated: %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println() // Empty line between tasks
	}
	return nil
}

func deleteTaskFromFile(id string) error {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id entered: %w", err)
	}

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	foundIndex := -1
	for i, task := range tasks {
		if task.ID == taskID {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	tasks = append(tasks[:foundIndex], tasks[foundIndex+1:]...)
	return writeTasks(tasks)
}

func updateTask(id string, description string) error {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid task id given: %w", err)
	}

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	foundIndex := -1
	for i, task := range tasks {
		if task.ID == taskID {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	tasks[foundIndex].Description = description
	tasks[foundIndex].UpdatedAt = time.Now()
	return writeTasks(tasks)
}

func markDone(id string, flag bool) error {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid task id given: %w", err)
	}

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	foundIndex := -1
	for i, task := range tasks {
		if task.ID == taskID {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	if flag {
		tasks[foundIndex].Status = "Done"
	} else {
		tasks[foundIndex].Status = "In-Progress"
	}

	return writeTasks(tasks)
}
