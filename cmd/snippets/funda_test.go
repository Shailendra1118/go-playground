package snippets

import (
	"fmt"
	"testing"
	"math"
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

 func TestCompareBin(t *testing.T) {
	str1 := "0011"
	str2 := "1100"
	length := math.Max(float64(len(str1)), float64(len(str2)))
	var res[] string
	for i :=0; i<int(length); i++ {
		if str1[len(str1)-1] == '1' || str2[len(str2)-1] == '1' {
			res = append(res, "1") // '1'
		} else {
			res = append(res, "0")
		}

		//str1 = str1 >> 1
		//str2 = str2 >> 1
	}

	fmt.Printf("Res is %v\n", res)

 }

 func TestMapContains(t *testing.T) {
	m := make(map[string]string)
	m["one"] = "1"
	fmt.Printf("Map: %v\n", m["one"])


	// declare and initialize new map in same line like below
	anotherMap := map[string]int32{"Shailendra": 100, "Singh": 500}
	fmt.Printf("Another: %v\n", anotherMap["Shailendra"])

	delete(anotherMap, "Singh")
	fmt.Printf("Another: %v\n", anotherMap["Singh"]) //return 0

	val, ok := anotherMap["Shailendra"] //case sensitive
	fmt.Printf("Value: %v and Ok: %v\n", val, ok)
 }

 func TestDetectCycle(t *testing.T) {
	var head *Node
	visited := make(map[*Node]bool)

	first := newNode(1, nil)
	sec := newNode(2, nil)
	th := newNode(3, nil)
	//four := newNode(4, nil)

	//first := newNode(1, newNode(2, newNode(2, newNode(4, nil))))
	first.next = sec
	sec.next = th
	th.next = first
	head = first

	
	for n:= head; n != nil; n = n.next {
		if visited[n] {
			fmt.Print("There is a cycle in LL")
			break
		}
		visited[n] = true
		fmt.Printf("Visted: %v\n", n.data)
	}

 }