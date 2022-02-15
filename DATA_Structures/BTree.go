package main

import "fmt"
type Node struct{
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) insert(key int){

	if(n.Key < key){
		if(n.Right == nil){
			n.Right = &Node{Key:key}
	    }else{
	    	n.Right.insert(key)
	    }
	}else if(n.Key > key){
		if(n.Left == nil){
			n.Left = &Node{Key:key}
	    }else{
	    	n.Left.insert(key)
	    }
	}
}

func (n *Node) search(key int)bool{

	if(n == nil){
		return false
	}
	
	if(n.Key < key){
		return n.Right.search(key)
	}else if(n.Key > key){
		return n.Left.search(key)
	}
	return true
}

func main(){
	Tree := &Node{Key:100}
	Tree.insert(11)
	fmt.Println(Tree.search(1))
	fmt.Println(Tree)
	
}