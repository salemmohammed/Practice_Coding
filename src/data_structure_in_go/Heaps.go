// https://www.youtube.com/watch?v=3DYIgTC4T1o&list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6
package main
import "fmt"
// Maxheap struct has the slice to hold the array
type Maxheap struct{
	array []int;
}
// insert adds an element to the heap
func (h *Maxheap) Insert(key int){
	// at the begining, we append the key to the array
	h.array = append(h.array,key)
	// index of key because we append it in the array and the index is the last
	h.maxHeapifyUp(len(h.array)-1)
}

// Extract returns the largest key, and removes it form the heap.
func (h *Maxheap) Extract() int{
	extracted := h.array[0]
	l := len(h.array)-1

	if len(h.array) == 0{
		fmt.Println("cannot extract because array length is zero")
		return -1
	}

	// take out the last index and put it in the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}
// maxHeapifyUp from botton top approach
func (h *Maxheap) maxHeapifyUp(index int){
	// we need to swap 
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index),index)
		index = parent(index)
	}
}

// maxHeapifyUp from Top down approach
func (h *Maxheap) maxHeapifyDown(index int){
	
	lastIndex := len(h.array)-1
	l,r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child

	for l <= lastIndex {
		if l == lastIndex { // left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger 
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
	
	// compare array value of current index to larger child and swap in smaller 
	if h.array[index] < h.array[childToCompare]{
		h.swap(index, childToCompare)
		index = childToCompare
		l,r = left(index), right(index)
	} else {
		return
	}
	}
}
// get the parent index
func parent(i int) int{
	return (i-1)/2
}
// get the left child index
func left(i int) int{
	return (2*i)+1
}
// get the right child index
func right(i int) int{
	return (2*i)+2
}
// swap keys in the array
func (h *Maxheap) swap(i1,i2 int){
	h.array[i1],h.array[i2] = h.array[i2],h.array[i1]
}
func main(){
	m := &Maxheap{}
	fmt.Println(m)
	buildHeap := []int{10,20,30,5,7,9,11,13,15,17}
	for _,v := range buildHeap{
		m.Insert(v)
		fmt.Println(m)
	}
	for i:=0; i<5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}