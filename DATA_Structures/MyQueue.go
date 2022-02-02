

package main
import "fmt"

type Queue struct{
	item []int
}


// enqueue
func (s *Queue) enqueue(i int){
	s.item = append(s.item,i)
} 

// dequeue
func (s *Queue) dequeue(){
	s.item = s.item[1:len(s.item)]
} 

func main(){
	myQueue := Queue{}
	myQueue.enqueue(100)
	myQueue.enqueue(200)
	myQueue.enqueue(300)
	myQueue.enqueue(400)
	
	fmt.Println(myQueue)
	
	myQueue.dequeue()
	myQueue.dequeue()

	fmt.Println(myQueue)
}