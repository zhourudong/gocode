package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	left  *Student
	right *Student
}

func main() {

	var root *Student = new(Student)

	root.Name = "root"
	root.Age = 18
	//root.left = nil    // default is nil
	//root.right = nil   // default is nil

	var left1 *Student = new(Student)

	left1.Name = "left1"
	left1.Age = 111

	root.left = left1 // 插入子节点左树

	var right1 *Student = new(Student)

	right1.Name = "right1"
	right1.Age = 222

	root.right = right1 // 插入子节点右子树

	var left1_1 *Student = new(Student)

	left1_1.Name = "left1_1"
	left1_1.Age = 111

	left1.left = left1_1 // 插入子节点左树

	trans(root)

}
func trans(root *Student) {

	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}
