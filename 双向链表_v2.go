package main

import (
	"fmt"
	"math/rand"
)

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

	trans(head)

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


/*

头部插入

{头部插入stu9 9 21.855305 0xc42007a540}

{头部插入stu8 8 67.90846 0xc42007a510}

{头部插入stu7 7 29.310184 0xc42007a4e0}

{头部插入stu6 6 28.303415 0xc42007a4b0}

{头部插入stu5 5 46.888985 0xc42007a480}

{头部插入stu4 4 31.805817 0xc42007a450}

{头部插入stu3 3 38.06572 0xc42007a420}

{头部插入stu2 2 21.426388 0xc42007a3f0}

{头部插入stu1 1 81.36399 0xc42007a3c0}

{头部插入stu0 0 51.521267 0xc42007a1b0}

{suse 18 100 0xc42007a1e0}

{尾部插入0 0 60.466026 0xc42007a210}

{尾部插入1 1 94.05091 0xc42007a240}

{尾部插入2 2 66.45601 0xc42007a270}

{尾部插入3 3 43.77142 0xc42007a2a0}

{尾部插入4 4 42.46375 0xc42007a2d0}

{尾部插入5 5 68.682304 0xc42007a300}

{尾部插入6 6 6.563702 0xc42007a330}

{尾部插入7 7 15.651925 0xc42007a360}

{尾部插入8 8 9.696952 0xc42007a390}

{尾部插入9 9 30.091187 <nil>}

*/
