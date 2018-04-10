package main

import (
	"fmt"
	"math/rand"
)

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

func 头部插入_bug(p *Student) {

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("头部插入stu%d", i),
			Age:   i,
			Score: rand.Float32() * 100,
		}
		//stu.next = head_1
		//head_1 = &stu

		stu.next = p
		p = &stu
	}

}

func 头部插入(p **Student) { // 传递指针的指针
	//  指针的指针 要改变它的变量,要传指针的地址进去

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("头部插入stu%d", i),
			Age:   i,
			Score: rand.Float32() * 100,
		}
		//stu.next = head_1
		//head_1 = &stu

		stu.next = *p
		*p = &stu
	}

	//	 要改变 "变量" 的地址,就传变量的指针进去
	//	 如果要改变 "指针变量" 的值, 传递指针变量的地址进去(指针的指针)
}

func delNode(p *Student) {

	var prev *Student = p // 上一个节点
	for p != nil {
		if p.Name == "尾部插入5" {
			prev.next = p.next
			break
		}
		prev = p   // 上一个节点
		p = p.next // 如果不是要往前挪动(下一个节点)
	}

	/*
	   有问题代码, 如果是头的话没有上一节点, 则是nil,会报错
	   解决方案,指针的指针
	*/

}

func addNode(p, newNode *Student) {

	//	先找到这个节点的位置
	for p != nil {
		if p.Name == "尾部插入6" {
			newNode.next = p.next
			p.next = newNode
			break
		}

		p = p.next // 如果不是要往前挪动(下一个节点)
	}

}
func main() {

	// 声明并分配空间
	//var  head *Student = &Student{} == 等价于 var  head *Student = new(Student)

	var head *Student = new(Student)

	head.Name = "suse"
	head.Age = 18
	head.Score = 100

	weibucharu_danxianglianbiao(head)

	fmt.Printf("\n")
	println("头部插入")

	// 头部插入_bug(&head)

	头部插入(&head) // 指针的指针

	delNode(head) // 删除节点 尾部插入5

	// add new node
	var newnode_tmp *Student = new(Student)
	newnode_tmp.Name = "newnode_tmp"
	newnode_tmp.Age = 18
	newnode_tmp.Score = 100

	addNode(head, newnode_tmp)

	trans(head)

}
