package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	var x map[string]int
	fmt.Println(x == nil)
	fmt.Println(x["something"])
	_, ok := x["something"]
	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not found")
	}
	m := make(map[string]int, 2)
	m["Said"] = 20
	m["Jamshed"] = 18
	m["Sodik"] = 19
	m["Rahim"] = 21
	fmt.Println(len(m))
	// Есть 4 разных способа создать мапу
	delete(m, "1")
	fmt.Println(m)
	// Итерация
	set := make(map[string]struct{})
	set["Item1"] = struct{}{}
	set["Item2"] = struct{}{}
	// Проверка на действительность в мапе
	if x, exists := set["Item1"]; exists {
		fmt.Println("item exists", x)
	}
	// Ключи
	t := make(map[string]*User)
	t["User1"] = &User{"Killer", "uehweouhcew"}
	fmt.Println(t["User1"].Name)

	c := [5]int{1, 2, 3, 4, 5}
	s := [5]int{1, 2, 3, 2, 5}
	if s == c {
		fmt.Println("tt")
	}
	h := make(map[[5]int]string)

	h[c] = "1"
	h[s] = "2"
	fmt.Println(h)
	for key := range h {
		delete(h, key)
	}

}
