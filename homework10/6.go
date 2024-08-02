package main

type Task struct {
	title     string
	completed bool
}
type TaskList struct {
	Tasks []Task
}

func (t *TaskList) AddTask(title string) {
	t.Tasks = append(t.Tasks, Task{title: title})
}
func (t TaskList) CompletedTasks() int {
	number := 0
	for _, Task := range t.Tasks {
		if Task.completed {
			number++
		}
	}
	return number
}
