package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Is_Student bool   `json:"is_student"`
}
type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func main() {
	/*
		x := make([]Person, 0)
		f, err := os.Open("1.json")
		if err != nil {
			panic(err)
		}
		r := json.NewDecoder(f)
		err = r.Decode(&x)
		if err != nil {
			panic(err)
		}
		fmt.Println(x) */
	/*
		p := Person{Name: "Said", Age: 20, Is_Student: true}
		data, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}
		file, err := os.OpenFile("example.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	*/
	b := Book{Title: "cmekce", Author: "vesvoskce", Price: 388.43}
	data, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("book.json", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	data, err = json.MarshalIndent(b, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}
