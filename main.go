package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tasks []Task
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Task CLI ---")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Edit task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			tasks = AddTask(name, tasks)
			fmt.Println("Task added successfully")

		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				continue
			}
			ListTasks(tasks)

		case 3:
			fmt.Print("Enter existing task name: ")
			oldName, _ := reader.ReadString('\n')
			oldName = strings.TrimSpace(oldName)

			fmt.Print("Enter new task name, or leave empty to keep same: ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)

			fmt.Print("Change completion status? yes/no: ")
			answer, _ := reader.ReadString('\n')
			answer = strings.ToLower(strings.TrimSpace(answer))

			var completedPtr *bool

			if answer == "yes" || answer == "y" {
				fmt.Print("Is task completed? true/false: ")
				statusInput, _ := reader.ReadString('\n')
				statusInput = strings.ToLower(strings.TrimSpace(statusInput))

				completed, err := strconv.ParseBool(statusInput)
				if err != nil {
					fmt.Println("Invalid completion status")
					continue
				}

				completedPtr = &completed
			}

			var err error
			tasks, err = EditTask(oldName, newName, completedPtr, tasks)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Task edited successfully")

		case 4:
			fmt.Print("Enter task name to delete: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			var err error
			tasks, err = DeleteTask(name, tasks)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Task deleted successfully")

		case 5:
			fmt.Println("Goodbye")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}