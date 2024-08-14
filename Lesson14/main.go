package main

import (
	"os"
)

func main() {
	/*
		r := strings.NewReader("Hello, Reader!")
		buf := make([]byte, 8)
		for {
			n, err := r.Read(buf)
			if err == io.EOF {
				break
			}
			fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
		}
	*/
	/*
		file, err := os.Create("output.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		data := []byte("Hello, Writer!\n")
		n, err := file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Printf("Wrote %d bytes to file,\n", n)
	*/
	/*
		file, err := os.Create("NEw.txt")
		if err != nil {
			log.Fatal(err)
		}
		file, err = os.OpenFile("NEw.txt", os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		b := bufio.NewWriter(file)
		b.WriteString("vsvjiowehvjwoivhwo\n")
		b.Flush()
		defer file.Close()
	*/
	/*
		file, err := os.OpenFile("Said.txt", os.O_RDWR|os.O_APPEND, 0755)
		if err != nil {
			log.Fatal(err)
		}
		b := bufio.NewWriter(file)
		b.WriteString("11111\n")
		b.WriteString("2222222222\n")
		b.Flush()
		buf := make([]byte, 5)
		for {
			n, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			fmt.Printf("%s\n", buf[:n])
		} */
	/*
		x, err := os.OpenFile("NewFile.txt", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		defer x.Close()

		x.WriteString("Hello, World!")

		if err != nil {
			panic(err)
		}
		buf := make([]byte, 8)
		for {
			n, err := x.Read(buf)
			if err == io.EOF {
				break
			}
			fmt.Printf("Read: %s\n", buf[:n])
		} */
	/*
		x, err := os.OpenFile("NewFile.txt", os.O_RDONLY, 0755)
		if err != nil {
			panic(err)
		}
		var input string
		fmt.Print("Enter your name: ")
		fmt.Fscan(x, &input)
		fmt.Printf("Hello, %s!\n", input) */
	_, err := os.OpenFile("NNNN", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}

}
