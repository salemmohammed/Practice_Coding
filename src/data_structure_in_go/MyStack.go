

package main
import "fmt"

type Stack struct{
	item []int
}


// push
func (s *Stack) push(i int){
	s.item = append(s.item,i)
} 

// pop
func (s *Stack) pop(){
	s.item = s.item[:len(s.item)-1]
} 

func main(){
	mystack := Stack{}
	mystack.push(100)
	mystack.push(200)
	mystack.push(300)
	mystack.push(400)
	
	fmt.Println(mystack)
	
	mystack.pop()
	mystack.pop()

	fmt.Println(mystack)
}