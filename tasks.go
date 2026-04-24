package main
import (
	"time"
	"fmt"
)
// task has a name, isCompleted, createdat, updatedat
type Task struct{
	Name 			string
	isCompleted 	bool
	createdAt 		time.Time
	updatedAt 		time.Time

}

func AddTask(name string, tasks []Task) []Task {
	t := Task{
		Name: name,
		isCompleted: false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	tasks = append(tasks, t)
	return tasks
	
}

func ListTasks(tasks []Task){
	for _,task := range tasks{
		fmt.Printf("Task Name: %v isCompleted: %v\n",task.Name,task.isCompleted)
	}
}