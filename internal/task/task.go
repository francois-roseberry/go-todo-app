package task

import "fmt"

type Task struct {
	Id   int
	Name string
	Done bool
}

func NewTask(index int) *Task {
	return &Task{
		Id:   index,
		Name: fmt.Sprintf("Task %d", index),
	}
}

func TaskList() []*Task {
	tasks := []*Task{}
	for i := 1; i < 10; i++ {
		tasks = append(tasks, NewTask(i))
	}
	return tasks
}
