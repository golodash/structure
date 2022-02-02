package ll

import (
	"reflect"
	"errors"
	"fmt"

	"github.com/golodash/structure/internal"
)

type (
	Node[T interface{}] struct {
		data *T
		Prev *Node[T]
		Next *Node[T]
	}

	LinkList[T interface{}] struct {
		Head *Node[T]
		Tail *Node[T]
		Size int
		functions map[string]interface{}
	}
)

// `functions` variable does not support 'generics' because we are still in 1.18beta1
//
// TODO: add support for generics
func New[T interface{}](functions map[string]interface{}) LinkList[T] {
	ll := LinkList[T]{
		Head: nil,
		Tail: nil,
		functions: map[string]interface{}{},
	}

	for k, v := range functions {
		if reflect.TypeOf(v).Kind() == reflect.Func {
			ll.functions[k] = v
		}
	}

	return ll
}

// O(n)
func (ll *LinkList[T]) Clear() {
	for trav := ll.Head; trav != nil; trav = trav.Next {
		if trav.Prev != nil {
			trav.Prev.Next = nil
		}
		trav.data = nil
		trav.Prev = nil
	}

	ll.Head = nil
	ll.Tail = nil
	ll.Size = 0
}

func (ll *LinkList[T]) Run(function string, params... interface{}) ([]reflect.Value, error) {
	if _, ok := ll.functions[function]; ok {
		return internal.CallJobFuncWithParams(ll.functions[function], params)
	}

	return nil, errors.New(fmt.Sprintf("%s not found", function))
}

func (ll *LinkList[T]) oneSizePlus() {
	ll.Size++
}

func (ll *LinkList[T]) oneSizeMinus() {
	ll.Size--
	if ll.Size < 0 {
		ll.Size = 0
	}
}

func (ll *LinkList[T]) GetSize() int {
	return ll.Size
}

func (ll *LinkList[T]) IsEmpty() bool {
	return ll.Size == 0
}

// O(1)
func (ll *LinkList[T]) AddLast(data *T) *T {
	var node *Node[T] = nil

	node = &Node[T]{
		data: data,
		Prev: nil,
		Next: nil,
	}

	if ll.Size == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		node.Prev = ll.Tail
		ll.Tail.Next = node
		ll.Tail = node
	}

	ll.oneSizePlus()

	return data
}

// O(1)
func (ll *LinkList[T]) AddFirst(data *T) *T {
	var node *Node[T] = nil

	node = &Node[T]{
		data: data,
		Prev: nil,
		Next: nil,
	}

	if ll.Size == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		node.Next = ll.Head
		ll.Head.Prev = node
		ll.Head = node
	}

	ll.oneSizePlus()

	return data
}

func (ll *LinkList[T]) addBeforeNode(node *Node[T], data *T) *T {
	var tempNode *Node[T] = &Node[T] {
		data: data,
		Prev: nil,
		Next: nil,
	}

	if node.Prev != nil {
		node.Prev.Next = tempNode
		tempNode.Prev = node.Prev
	}
	tempNode.Next = node
	node.Prev = tempNode

	ll.oneSizePlus()

	return data
}

// O(n)
func (ll *LinkList[T]) Add(index int, data *T) *T {
	if index == 0 {
		return ll.AddFirst(data)
	} else if index == ll.Size {
		return ll.AddLast(data)
	} else if index < 0 || index > ll.Size {
		return nil
	}

	for trav, i := ll.Head, 0; trav != nil; trav, i = trav.Next, i+1 {
		if i == index {
			return ll.addBeforeNode(trav, data)
		}
	}

	return nil
}

// O(n)
func (ll *LinkList[T]) RemoveData(data *T) *T {
	for trav := ll.Head; trav != nil; trav = trav.Next {
		if trav.data == data {
			return ll.RemoveNode(trav)
		}
	}

	return nil
}

// O(n)
func (ll *LinkList[T]) Remove(index int) *T {
	if index == 0 {
		return ll.RemoveFirst()
	} else if index == ll.Size {
		return ll.RemoveLast()
	} else if index < 0 || index > ll.Size {
		return nil
	}

	for trav, i := ll.Head, 0; trav != nil; trav, i = trav.Next, i+1 {
		if i == index {
			return ll.RemoveNode(trav)
		}
	}

	return nil
}

// O(1)
func (ll *LinkList[T]) RemoveNode(node *Node[T]) *T {
	if node == nil {
		return nil
	}
	
	var data *T = nil

	if ll.Head == node {
		return ll.RemoveFirst()
	} else if ll.Tail == node {
		return ll.RemoveLast()
	} else {
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}

		data = node.data
		node.Next = nil
		node.Prev = nil
		node.data = nil
	}

	ll.oneSizeMinus()

	return data
}

// O(1)
func (ll *LinkList[T]) RemoveLast() *T {
	var data *T = nil

	if ll.Size != 0 {
		if ll.Head == ll.Tail {
			ll.Head = nil
		}
		node := ll.Tail
		ll.Tail = node.Prev
		if ll.Tail != nil {
			ll.Tail.Next = nil
		}

		data = node.data
		node.Next = nil
		node.Prev = nil
		node.data = nil
	}

	ll.oneSizeMinus()

	return data
}

// O(1)
func (ll *LinkList[T]) RemoveFirst() *T {
	var data *T = nil

	if ll.Size != 0 {
		if ll.Head == ll.Tail {
			ll.Tail = nil
		}
		node := ll.Head
		ll.Head = node.Next
		if ll.Head != nil {
			ll.Head.Prev = nil
		}

		data = node.data
		node.Next = nil
		node.Prev = nil
		node.data = nil
	}

	ll.oneSizeMinus()

	return data
}

// O(n)
func (ll *LinkList[T]) GetIndex(data *T) int {
	for i, trav := 0, ll.Head; i < ll.Size; i, trav = i+1, trav.Next {
		if trav.data == data {
			return i
		}
	}

	return -1
}

// O(n)
func (ll *LinkList[T]) Contains(data *T) bool {
	return ll.GetIndex(data) != -1
}

// O(n)
func (ll *LinkList[T]) GetNode(index int) *Node[T] {
	if index > ll.Size || index < 0 {
		return nil
	}

	for i, trav := 0, ll.Head; i < ll.Size; i, trav = i+1, trav.Next {
		if i == index {
			return trav
		}
	}

	return nil
}

// O(n)
func (ll *LinkList[T]) FindNode(data *T) *Node[T] {
	for trav := ll.Head; trav != nil; trav = trav.Next {
		if trav.data == data {
			return trav
		}
	}

	return nil
}

func (ll *LinkList[T]) DisplaceTo(node *Node[T], index int) bool {
	if index >= ll.Size || index < 0 {
		return false
	}

	data := node.data

	ll.RemoveNode(node)
	ll.Add(index, data)

	return true
}

func (ll *LinkList[T]) DisplaceIndex(index1 int, index2 int) bool {
	return ll.Displace(ll.GetNode(index1), ll.GetNode(index2))
}

// O(1)
func (ll *LinkList[T]) Displace(node1 *Node[T], node2 *Node[T]) bool {
	if node1 == node2 {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	if node1 == ll.Head {
		ll.Head = node2
	} else if node2 == ll.Head {
		ll.Head = node1
	}
	if node2 == ll.Tail {
		ll.Tail = node1
	} else if node1 == ll.Tail {
		ll.Tail = node2
	}

	if node1.Next == node2 {
		if node1.Prev != nil {
			node1.Prev.Next = node2
		}
		if node2.Next != nil {
			node2.Next.Prev = node1
		}

		node1Prev := node1.Prev

		node1.Next = node2.Next
		node1.Prev = node2
		node2.Next = node1
		node2.Prev = node1Prev

		return true
	} else if node2.Next == node1 {
		if node2.Prev != nil {
			node2.Prev.Next = node1
		}
		if node1.Next != nil {
			node1.Next.Prev = node2
		}

		node2Prev := node2.Prev

		node2.Next = node1.Next
		node2.Prev = node1
		node1.Next = node2
		node1.Prev = node2Prev

		return true
	}

	node1Prev := node1.Prev
	if node1.Prev != nil {
		node1.Prev.Next = node2
	}
	node1Next := node1.Next
	if node1.Next != nil {
		node1.Next.Prev = node2
	}
	node2Prev := node2.Prev
	if node2.Prev != nil {
		node2.Prev.Next = node1
	}
	node2Next := node2.Next
	if node2.Next != nil {
		node2.Next.Prev = node1
	}

	node1.Next = node2Next
	node2.Next = node1Next
	node1.Prev = node2Prev
	node2.Prev = node1Prev

	return true
}

// O(n)
func (ll *LinkList[T]) ReturnAsSlice() []T {
	slice := []T{}
	for trav := ll.Head; trav != nil; trav = trav.Next {
		slice = append(slice, *trav.data)
	}

	return slice
}

func (ll *LinkList[T]) FirstNode() *Node[T] {
	return ll.Head
}

func (ll *LinkList[T]) LastNode() *Node[T] {
	return ll.Tail
}

func (ll *LinkList[T]) First() *T {
	if ll.Head != nil {
		return ll.Head.data
	}
	return nil
}

func (ll *LinkList[T]) Last() *T {
	if ll.Head != nil {
		return ll.Tail.data
	}
	return nil
}

func (node *Node[T]) GetData() *T {
	return node.data
}

func (node *Node[T]) SetData(data *T) *T {
	node.data = data
	return node.data
}
