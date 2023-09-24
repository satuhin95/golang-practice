package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("============= Todo List ##################")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Exit")
		fmt.Println("===========================================")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Enter task name :")
			scanner.Scan()
			taskName := scanner.Text()
			todoList.AddTask(taskName)
		case "2":
			todoList.ViesTask()
		case "3":
			fmt.Println("Enter task id :")
			scanner.Scan()
			taskIdStr := scanner.Text()
			taskId, err := strconv.Atoi(taskIdStr)
			if err != nil {
				fmt.Println("Invalid task id")
				continue
			}
			todoList.MarkTaskAsDone(taskId)
		case "4":
			fmt.Println("Exit...")
			return
		default:
			fmt.Println("Invalid choice. Please try again")
		}
	}
}
