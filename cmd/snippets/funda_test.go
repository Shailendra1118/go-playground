package snippets

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	//fmt.Printf("Hello Test")

	var name = []string{"shailendra", "singh", "yadav"}
	for i := 0; i < 2; i++ {
		fmt.Printf("i is %v\n", name[i])
	}
	dest := []string{"james", "bond"}
	copy(name, dest)
	for i, v := range dest {
		fmt.Printf("coying %v dest: %v\n", v, dest[i])
	}
 }


 func TestList(t *testing.T) {
	//slices are refer to underlying arrays, they are references
	animals := []string{"Lion", "Monkey", "Snake", "Bees", "Cheetah"}
	//monkey := animals[1]
	jungle := animals[0:3]
	fmt.Printf("Monkey is: %v\n", animals[1])
	animals[1] = "Tiger"
	fmt.Printf("Monkey is: %v\n", animals[1])

	for _, item := range jungle {
		fmt.Printf("Animal is: %v\n", item)
	}

	fmt.Printf("Last is: %v\n", jungle[2])
 }


 func TestSlic(t *testing.T) {
	slice := []int{100,5,499}
	slice[1] = 88
	for _, item := range slice {
		fmt.Printf("Item is: %v\n", item)
	}
 }

 // Singlely Linkedlist Implementation

 func TestLinkedListOps(t *testing.T) {
	//head := newNode(100, nil)
	//fmt.Printf("Node is: %v\n", head)
	list := LinkedList{}
	list.add(20)
	list.add(10)
	list.add(80)
	list.display()
 }

 type Node struct { data int; next *Node }

 type LinkedList struct {
	head *Node
 }

 func newNode(data int, next *Node) *Node {
	return &Node{data: data, next: next}
 }

 func (ll *LinkedList) add(data int) *Node {
	node := newNode(data, nil)
	if ll.head == nil {
		ll.head = node;
	} else {
		cur := ll.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node;
	}
	return ll.head
 }

 func (ll *LinkedList) display() {
	cur := ll.head
	for cur != nil {
		fmt.Printf("%v --> ", cur.data)
		cur = cur.next
	}
	fmt.Println("nil")
 }