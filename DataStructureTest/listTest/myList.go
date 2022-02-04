package listtest

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n *Node) GetData() int {
	return n.data
}

func (n *Node)SetData(data int) {
	n.data = data
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(p *Node) {
	n.next = p
}



func createNode(data int) *Node {
	var n *Node = new(Node)
	n.SetData(data)
	n.SetNext(nil)
	return n
}
func createHeadNode() *Node {
	var head *Node = new(Node)
	head.SetNext(nil)
	return head
}

func CreateList(num int) *Node{
	var head = createHeadNode()
	var p = head
	for i := 1; i <= num; i++ {
		var number int
		number = i * 3
		var n *Node = createNode(number)
		p.SetNext(n)
		p = n
	}
	return head
}

func (head *Node) ShowList() {
	var p *Node = head.next
	for{
		if p == nil {
			break
		}
		fmt.Println(p.GetData())
		p = p.GetNext()
	}
}