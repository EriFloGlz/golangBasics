package main

import "fmt"

type User struct {
	Name string
}
type Task struct {
	Name        string
	Description string
	Status      string
	User        User
}
type TaskList []Task

func createTask() Task {
	name := ""
	description := ""
	userName := ""
	fmt.Println("Name")
	fmt.Scanln(&name)
	fmt.Println("Description:")
	fmt.Scanln(&description)
	fmt.Println("User")
	fmt.Scanln(&userName)
	task := Task{
		Name:        name,
		Description: description,
		Status:      "todo",
		User: User{
			Name: userName,
		},
	}
	return task

}
func (arr *TaskList) addTask(t Task) {
	*arr = append(*arr, t)
}

func (arr *TaskList) removeTask(taskToRemove string) {
	for i, v := range *arr {
		if v.Name == taskToRemove {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
		}
	}
}
func (tasks TaskList) printTasks() {
	for _, v := range tasks {
		fmt.Println("Task name", v.Name)
		fmt.Println("description", v.Description)
		fmt.Println("User", v.User.Name)
		fmt.Println("Status", v.Status)
	}
}
func menu() string {
	var option string
	fmt.Println("1. Add")
	fmt.Println("2. Remove")
	fmt.Println("3. Print tasks")
	fmt.Println("4. Update status")

	fmt.Scanln(&option)

	return option
}
func main() {

	var tasks TaskList

	for {
		opc := menu()

		switch opc {
		case "1":
			newTask := createTask()
			tasks.addTask(newTask)

		case "2":
			if len(tasks) == 0 {
				fmt.Println("No tasks yet")
				continue
			}
			taskToRemove := ""
			fmt.Println("Task to remove")
			fmt.Scanln(&taskToRemove)
			tasks.removeTask(taskToRemove)
		case "3":
			tasks.printTasks()

		case "4":
			fmt.Println("Task to change status")
			var taskToUpdate string
			fmt.Scanln(&taskToUpdate)
			var statusTask string
			fmt.Println("Ingresa el estado de la tarea : in-progress, resolve, ")
			fmt.Scanln(&statusTask)
			isFound := false

			for index, value := range tasks {
				if tasks[index].Name == taskToUpdate {
					tasks[index].Status = statusTask
					fmt.Print(value)
					isFound = true
					fmt.Println("success")
					break
				}

			}
			if !isFound {
				fmt.Println("Fail, task not found")
			}

		case "5":
			fmt.Print("Bye")
			return
		default:
			fmt.Print("Opcion no valida")
		}
	}

}
