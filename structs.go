package main

import "fmt"

//define struct
type Person struct {
	Name string
	Age  int
	Subjects
}

type Subjects struct {
	History     string
	English     string
	Mathematics string
}

//function to struct

func (student Person) Details() {
	fmt.Println("Within function name of studen is:", student.Name)
}

func main() {
	fmt.Println("Hello, erratic coder!")

	student := Person{Name: "John", Age: 34, Subjects: Subjects{History: "yes", English: "No", Mathematics: "yes"}}

	student2 := Person{Name: "Andreas", Age: 67, Subjects: Subjects{History: "No", English: "Yes", Mathematics: "yes"}}

	fmt.Println("person type variale is:", student)
	fmt.Println("person type variale is:", student2)

	fmt.Println("Name is:", student.Name)
	fmt.Println("Age is:", student.Age)

	//struct withing struct

	//Attach method
	student.Details()

	//Use  anonymous fields within struct

}
