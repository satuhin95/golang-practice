package main

import "fmt"

type Task struct {
	ID   int
	Name string
	Done bool
}

func (t *Task) MarkAsDone() {
	if t.Done {
		fmt.Println("Task already Done!")
	}
	t.Done = true
	fmt.Println("Task Mark as  Done!")

}

func NewTask(taskId int, taskName string) Task {
	return Task{ID: taskId, Name: taskName}
}
