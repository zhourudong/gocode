package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 声明并分配空间
	//var  head *Student = &Student{} == 等价于 var  head *Student = new(Student)

	var head Student

	head.Name = "suse"
	head.Age = 18
	head.Score = 100

	weibucharu_danxianglianbiao(&head)

	var head_1 *Student
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("头部插入stu%d", i),
			Age:   i,
			Score: rand.Float32() * 100,
		}
		stu.next = head_1
		head_1 = &stu
	}

	trans(&head)

	fmt.Printf("\n")
	println("头部插入")
	trans(head_1)

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

func weibucharu_danxianglianbiao(tail *Student) {
	// 尾部插入

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:  fmt.Sprintf("尾部插入%d", i),
			Age:   i,
			Score: rand.Float32() * 100,
		}
		tail.next = stu
		tail = stu
	}

}
