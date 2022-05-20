/*
	This is an implementation of 
	data structure called hash table.
*/
package main
import "fmt"

const ArraySize = 7
// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}
// bucket structure
type bucket struct {
	head *bucketNode
}
// bucketNode is a linked list data structure
type bucketNode struct {
	key string
	next *bucketNode
}
// insert
func (h *HashTable)Insert(key string){
	index := hash(key)
	h.array[index].insert(key)
}
// search
func (h *HashTable)Search(key string) bool{
	index := hash(key)

	return h.array[index].search(key)
}
// delete
func (h *HashTable)Delete(key string){
	index := hash(key)
	h.array[index].delete(key)
}
// insert
// insert will take a key and create a node with the key and insert the node in
// the bucket
func (b *bucket) insert(k string){
	if !b.search(k){
		newNode := &bucketNode{key:k}
		newNode.next = b.head
		b.head = newNode
	}else{
		fmt.Println(k, "the value is already exists")
	}
}
// search
func (b *bucket) search(k string) bool{
	currentNode := b.head
	for currentNode != nil{
		if currentNode.key == k{
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// delete
func (b *bucket) delete(k string) bool{
	if b.head.key == k{
		b.head = b.head.next
		return true
	}
	previousNode := b.head
	for previousNode.next != nil{
		if previousNode.next.key == k{
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
	return false
}
// hash
func hash(key string)int{
	sum :=0
	for _,v := range key{
		sum+=int(v)
	}
	return sum % ArraySize
}
// init
func Init() *HashTable{
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main(){

	testHash := Init()
	list := []string{
		"Salem",
		"Ali",
		"Faris",
	}
	for _,v := range list{
		testHash.Insert(v)
	}
	testHash.Delete("Salem")
	fmt.Println("Ali", testHash.Search("Ali"))
	fmt.Println("Salem", testHash.Search("Salem"))

}