package main

import "fmt"

func main() {

	var head Student

	head.Name = "suse"
	head.Age = 18
	head.Score = 100

	var stu1 Student
	stu1.Name = "stu1 suse"
	stu1.Age = 18
	stu1.Score = 100

	var stu2 Student
	stu2.Name = "stu1 suse"
	stu2.Age = 18
	stu2.Score = 100

	head.next = &stu1
	stu1.next = &stu2

	trans(&head)

}

type Student struct {
	Name  string
	Age   int
	Score float32

	next *Student
}

func trans(p *Student) {

	for p != nil {

		fmt.Println(*p)
		p = p.next
		fmt.Printf("\n")

	}

}
