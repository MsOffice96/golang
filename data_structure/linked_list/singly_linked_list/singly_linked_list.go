package main

import (
	"fmt"
	"reflect"
)

// Singly Linked List 구현

type Linked_List struct {
	Header *Node
	Tail   *Node
	Count  int
}

func NewLinkedList() *Linked_List {
	linked_List := &Linked_List{
		Header: nil,
		Tail:   nil,
		Count:  0,
	}
	return linked_List
}

func (ll *Linked_List) FindByValue(val interface{}) *Node {
	for n := ll.Header; n != nil; {
		if reflect.DeepEqual(n.Value, val) {
			return n
		} else {
			n = n.Next
		}
	}
	return nil
}

func (ll *Linked_List) FindeByIndex(index int) *Node {
	node := ll.Header
	if node == nil {
		return nil
	}

	for n := 0; n < index; n++ {
		if node == nil {
			return nil
		} else {
			node = node.Next
		}
	}

	return node
}

// singly Linked List 의 Insert Method 구현시 Linked List 관점 + Node 관점으로 구현하면 됨
func (ll *Linked_List) InsertFirst(node *Node) {
	if ll.Header == nil {
		ll.Header = node
		ll.Tail = node
		node.Next = nil
	} else {
		node.Next = ll.Header
		ll.Header = node
	}

	ll.Count++
}

func (ll *Linked_List) InsertLast(node *Node) {
	if ll.Tail == nil {
		ll.Header = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}
	ll.Count++
}

func (ll *Linked_List) InsertMiddle(node *Node, preNode *Node) bool {
	if preNode == nil {
		return false
	} else {
		node.Next = preNode.Next
		preNode.Next = node.Next
		if ll.Tail == preNode {
			ll.Tail = node
		}
		ll.Count++
		return true
	}
}

func (ll *Linked_List) RemoveFirst() {
	if ll.Header == nil {
		return
	} else if ll.Header == ll.Tail {
		ll.Header = nil
		ll.Tail = nil
	} else {
		ll.Header = ll.Header.Next
	}
	ll.Count--
}

func (ll *Linked_List) RemoveLast() {
	if ll.Tail == nil {
		return
	}

	if ll.Header == ll.Tail {
		ll.Header = nil
		ll.Tail = nil
	} else {
		for n := ll.Header; n != nil; {
			if n.Next == ll.Tail {
				n.Next = nil
				ll.Tail = n
			}
			n = n.Next
		}
	}
	ll.Count--
}

func (ll *Linked_List) Print() {
	str := "["

	if ll.Header == nil {
		str += "]"
		// return str
		fmt.Printf("%s\n", str)
		return
	}

	for n := ll.Header; n != nil; {
		str += fmt.Sprintf(" -> %+v", n.Value)
		n = n.Next
	}
	str += "]"
	fmt.Printf("list : %s / len: %d\n", str, ll.Count)
	// return str

}

type Node struct {
	Next  *Node
	Value interface{}
}

func NewNode(v interface{}) *Node {
	node := &Node{
		Next:  nil,
		Value: v,
	}
	return node
}

func main() {

	linked_list := NewLinkedList()
	linked_list.Print()

	linked_list.InsertFirst(NewNode(1))
	linked_list.Print()

	linked_list.InsertLast(NewNode(2))
	linked_list.Print()

	linked_list.InsertFirst(NewNode(3))
	linked_list.Print()

	linked_list.InsertLast(NewNode(4))
	linked_list.Print()

	linked_list.RemoveFirst()
	linked_list.Print()

	linked_list.RemoveLast()
	linked_list.Print()

}
