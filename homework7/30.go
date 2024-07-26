package main

import "fmt"

type Project struct {
	Name    string
	Manager Person2
}

func GetProject(p Project) {
	fmt.Printf("Название проекта %s, Имя менеджера: %s, Его адрес: %s %s\n", p.Name, p.Manager.Name, p.Manager.City, p.Manager.Street)
}

