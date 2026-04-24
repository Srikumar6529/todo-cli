package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("hello world!!")
	//lets build a simple cli based todo app
	//it needs to do 4 things
	//list current todo tasks
	//create a new todo task
	//edit a todo task
	//delete a todo task
	tasks := []Task
	AddTask("sampleTask1",tasks)
	fmt.Println("this is a sample task")
	fmt.Println(task)
}