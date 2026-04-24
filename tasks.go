package main
import (
	"time"
	"fmt"
)
// task has a name, isCompleted, createdat, updatedat
type Task struct{
	Name 			string
	IsCompleted 	bool
	CreatedAt 		time.Time
	UpdatedAt 		time.Time

}

func AddTask(name string, tasks []Task) []Task {
	now := time.Now()
	t := Task{
		Name: name,
		IsCompleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	tasks = append(tasks, t)
	return tasks
	
}

func ListTasks(tasks []Task){
	for _,task := range tasks{
		fmt.Printf("Task Name: %v isCompleted: %v\n",task.Name,task.IsCompleted)
	}
}

func DeleteTask(name string, tasks []Task) ([]Task, error) {
	for i,task := range tasks{
		if task.Name == name{
			tasks = append(tasks[:i],tasks[i+1:]...)
			return tasks,nil
		}
	}
	return tasks,fmt.Errorf("task not found")
}

func EditTask(oldName, NewName string, isCompleted *bool, tasks []Task) ([]Task, error){
	for i,task := range tasks{
		if task.Name == oldName{
			if NewName!=""{
				tasks[i].Name = NewName
			}
			if isCompleted != nil {
				tasks[i].IsCompleted = *isCompleted
			}

			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks,fmt.Errorf("task not found")
}