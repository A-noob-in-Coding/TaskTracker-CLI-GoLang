package Utilities

import (
	_ "encoding/json"
	"fmt"
	_ "io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func getNewId() (int, error) {
	const filename = "id.txt"

	// Read current ID from file
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, start from 1
			err := os.WriteFile(filename, []byte("1"), 0666)
			if err != nil {
				return 0, fmt.Errorf("failed to create id file: %w", err)
			}
			return 1, nil
		}
		return 0, fmt.Errorf("failed to read id file: %w", err)
	}

	// Parse current ID
	content := strings.TrimSpace(string(data))
	if content == "" {
		// Empty file, start from 1
		err := os.WriteFile(filename, []byte("1"), 0666)
		if err != nil {
			return 0, fmt.Errorf("failed to write initial id: %w", err)
		}
		return 1, nil
	}

	currentID, err := strconv.Atoi(content)
	if err != nil {
		return 0, fmt.Errorf("failed to parse id '%s': %w", content, err)
	}

	// Increment ID
	newID := currentID + 1

	// Write new ID back to file
	err = os.WriteFile(filename, []byte(strconv.Itoa(newID)), 0666)
	if err != nil {
		return 0, fmt.Errorf("failed to write new id: %w", err)
	}

	return newID, nil
}
func AddTask(description string) (bool, error) {
	newID, err := getNewId()
	if err != nil {
		return false, fmt.Errorf("failed to get new id for task: %w", err)
	}
	newTask := Task{ID: newID, Description: description, Status: "To Do", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	AddTaskToFile(newTask)
	return true, nil
}
func DeleteTask(id string) error {
	err := deleteTaskFromFile(id)
	return err
}

func UpdateTask(id string, description string) error {
	err := updateTask(id, description)
	return err
}

func MarkDone(id string) error {
	err := markDone(id, true)
	return err
}

func MarkProgress(id string) error {
	err := markDone(id, false)
	return err
}
