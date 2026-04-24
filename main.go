package main


func main(){
	//lets build a simple cli based todo app
	//it needs to do 4 things
	//list current todo tasks
	//create a new todo task
	//edit a todo task
	//delete a todo task
	var tasks []Task
	tasks = AddTask("sampleTask1",tasks)
	tasks = AddTask("sampleTask2",tasks)
	tasks = AddTask("sampleTask3",tasks)
	ListTasks(tasks)
}