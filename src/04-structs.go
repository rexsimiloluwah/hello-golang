package main

import "fmt"

// Structure Declaration
type Student struct {
	name       string
	age        int16
	department string
	courses    []string
	cgpa       float32
}

type Animal struct {
	name   string
	family string
}

type Bird struct {
	Animal // Embedding of the Animal struct
	speed  float32
	color  string
}

func main() {
	// Initialization
	student := Student{
		"Okunowo Similoluwa",
		20,
		"Electronic and Electrical Engineering",
		[]string{"EEE200", "EEE300", "EEE400"},
		4.90,
	}

	fmt.Printf("Student Details\n--------------------------\n")
	fmt.Printf("Name:- %s\nAge: - %d years\nDepartment: - %s\nCourses: - %v\nCGPA: - %.2f\n", student.name, student.age, student.department, student.courses, student.cgpa)

	// Modifying values in a struct
	student.name = "Okunowo Adetoyosi"
	fmt.Println(student)

	studentOne := createStudent("Yemisi Okunowo", 21, "Accounting", []string{"ECN500", "ACC506", "ACC507"}, 4.99)
	fmt.Println(studentOne)

	aDoctor := struct {
		name           string
		specialization string
	}{name: "Moyo Adepeju", specialization: "Surgical sciences"}
	fmt.Printf("%v, %T\n", aDoctor, aDoctor)

	// Structures embedding
	// Composition - Similar to inheritance in languages that support object oriented programming
	b := Bird{}
	b.name = "Bird"
	b.family = "Aves"
	b.speed = 32.56
	b.color = "green"
	fmt.Println(b)
}

func createStudent(name string, age int16, department string, courses []string, cgpa float32) Student {
	student := Student{}
	student.name = name
	student.age = age
	student.department = department
	student.courses = courses
	student.cgpa = cgpa

	return student
}
