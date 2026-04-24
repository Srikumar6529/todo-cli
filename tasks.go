package main
import (
	"time"
)
// task has a name, isCompleted, createdat, updatedat
type Task struct{
	Name 			string
	isCompleted 	bool
	createdAt 		time.Time
	updatedAt 		time.Time

}

func AddTask(name string, tasks []Task){
	t := Task{
		Name: name,
		isCompleted: false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	tasks = append(t, tasks)
}