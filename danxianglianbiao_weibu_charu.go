package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var head Student

	head.Name = "suse"
	head.Age = 18
	head.Score = 100

	var tail = &head
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   i,
			Score: rand.Float32() * 100,
		}
		tail.next = stu
		tail = stu
	}

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
