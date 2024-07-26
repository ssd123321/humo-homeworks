package main

type Task struct {
	Title     string
	Completed bool
}

func ChangeStatus(array *[]Task, index string) {
	array2 := *array
	for i := 0; i < len(array2); i++ {
		if array2[i].Title == index {
			array2[i].Completed = true
		}
	}
}
