package main
import "fmt"

type Node struct{
	Data int
	Next *Node
}

type LinkedList struct{
	head *Node
	length int
}

func (l LinkedList) PrintList(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.length --
	}
	fmt.Printf("\n")
}

// Dnode is to delete the node with value
func (l *LinkedList) Dnode(value int) {
	if l.length == 0 {
		return
	}
	l.head.Data == value {
		l.head = l.head.Next
		l.length--
		return
	}
	PrevioustoDelete := l.head
	for PrevioustoDelete.Next.Data != value {
		if PrevioustoDelete.Next.Next == nil{
			return
		}
		PrevioustoDelete = PrevioustoDelete.Next
	}
	PrevioustoDelete.Next = PrevioustoDelete.Next.Next
	l.length --
}

func (l *LinkedList) insert(n *Node){
	second := l.head
	l.head = n 
	l.head.Next = second
	l.length++
}

func main(){
	mylist := LinkedList{}
	node1  := &Node{Data:10}
	node2  := &Node{Data:11}
	node3  := &Node{Data:12}
	
	mylist.insert(node1)
	mylist.insert(node2)
	mylist.insert(node3)
	
	fmt.Println(mylist)
	mylist.PrintList()
	mylist = LinkedList{}
	mylist.Dnode(10)
	
	fmt.Println(mylist)
	mylist.PrintList()
}