package main

import "os"
import "TaskTracker/Utilities"
import "fmt"

func main() {
	args := os.Args
	decision := -1
	//fmt.Println(args)
	if len(args) == 2 && args[1] == "-h" {

		fmt.Println(`Task Tracker CLI - Usage:

  task-cli add "<task description>"          
      → Add a new task

  task-cli update <task-id> "<new description>"
      → Update an existing task

  task-cli delete <task-id>                 
      → Delete a task

  task-cli mark-in-progress <task-id>       
      → Mark a task as in progress

  task-cli mark-done <task-id>              
      → Mark a task as done

  task-cli list                             
      → List all tasks

  task-cli list [done|todo|inprogress]     
      → List tasks by status

  task-cli -h | --help                      
      → Show this help message\n
`)
		return
	}
	if !Utilities.Sanitize(args, &decision) {
		println("Invalid usage use TaskTracker -h for help")
		return
	}
	if decision == 1 {
		Utilities.ListTasks(-1)
	}
	if decision == 2 {
		Utilities.DeleteTask(args[2])
	}
	if decision == 4 {
		Utilities.AddTask(args[2])
	}
	if decision == 3 {
		Utilities.UpdateTask(args[2], args[3])
	}
	if decision == 6 {
		Utilities.MarkDone(args[2])
	}
	if decision == 5 {
		Utilities.MarkProgress(args[2])
	}
	if decision == 8 {
		Utilities.ListTasks(1)
	}
	if decision == 7 {
		Utilities.ListTasks(2)
	}
	if decision == 9 {
		Utilities.ListTasks(3)
	}
}
