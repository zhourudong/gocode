package main

import "fmt"

type Student struct {
	Id   int
	Name string
}
type Node struct {
	Val  Student
	Next *Node
}

func main() {
	node_a := Node{Val: Student{Id: 1, Name: "zhangsan"}}
	node_b := Node{Val: Student{Id: 2, Name: "lisi"}}
	node_c := Node{Val: Student{Id: 3, Name: "wangwu"}}
	node_d := Node{Val: Student{Id: 4, Name: "xxxxxxx"}}

	node_a.Next = &node_b
	node_a.Next.Next = &node_c      //node_b.Next = &node_c
	node_a.Next.Next.Next = &node_d //node_c.Next = &node_d

	p := &node_a
	printList(revrse_list(p))
	for p != nil {
		node := *p
		fmt.Println(node)
		p = node.Next
	}
	fmt.Println()
	printList(revrse_list(p))
}

func revrse_list(p *Node) *Node {
	var pre *Node

	for p != nil {
		//	暂存p下的下一个节点
		next := p.Next
		// 让当前节点p的下一个节点为pre
		p.Next = pre
		// 更新 pre为当前节点
		pre = p
		// 继续迭代, 变更 p 为其原来的next节点
		p = next
	}
	// 此时,p是nil,应当返回pre节点

	return pre
}

func printList(p *Node) {
	for {
		if p == nil {
			fmt.Println("p is nil.")
			break
		}
		fmt.Println(p)
		p = p.Next
	}
}
