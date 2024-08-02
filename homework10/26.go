package main

type Task1 struct {
	title       string
	description string
	status      string
}
type Project struct {
	Tasks []Task1
}

func (p *Project) AddTask(title, description, status string) {
	p.Tasks = append(p.Tasks, Task1{title, description, status})
}
func (p Project) CountTasksByStatus(status string) int {
	x := 0
	for _, v := range p.Tasks {
		if v.status == status {
			x++
		}
	}
	return x
}
