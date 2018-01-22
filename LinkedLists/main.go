package main

import "fmt"

//Node - single linked list class implementation
type Node struct {
	data interface{}
	next *Node
}

func main() {
	b := Node{}
	b.addToTail(5)
	fmt.Println(b)
	b.addToTail(9)
	fmt.Println(b)
	b.addToTail(98)
	fmt.Println(b.showElements())
}

func (n *Node) addToTail(d interface{}) {

	end := &Node{}
	end.data = d

	end.next = nil
	n.next = end
	//fmt.Println(n)
}

func (n *Node) showElements() (a []interface{}) {

	for n.next != nil {
		fmt.Println("saa", n.data)
		a = append(a, n.data)
		fmt.Println("jhk", n.next)
		n = n.next
	}
	return a
}
