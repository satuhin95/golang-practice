package main

import (
	"fmt"
)

type TodoList struct {
	tasks []Task
}

func (td *TodoList) AddTask(taskName string) {
	taskId := len(td.tasks) + 1
	task := NewTask(taskId, taskName)
	td.tasks = append(td.tasks, task)
	fmt.Printf("Task %s  Successfully ", taskName)
}

func (td *TodoList) ViesTask() {
	fmt.Println("**********Tasks***********")
	for _, task := range td.tasks {
		status := " "
		if task.Done {
			status = "X"
		}
		fmt.Printf("[%s] Task #%d %s\n", status, task.ID, task.Name)
	}
	fmt.Println("**************************")
}

func (td *TodoList) MarkTaskAsDone(taskId int) {
	if taskId < 1 || taskId > len(td.tasks) {
		fmt.Println("Invalid Task")
		return
	}
	td.tasks[taskId-1].MarkAsDone()
}
