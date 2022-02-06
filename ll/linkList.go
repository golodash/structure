package ll

import (
	"reflect"
	"errors"
	"fmt"

	"github.com/golodash/structure/internal"
)

type (
	Node[T any] struct {
		data *T
		Prev *Node[T]
		Next *Node[T]
	}

	LinkList[T any] struct {
		Head *Node[T]
		Tail *Node[T]
		size int
		functions map[string]any
	}
)

func New[T any](functions map[string]any) (*LinkList[T], error) {
	l := &LinkList[T]{
		Head: nil,
		Tail: nil,
		functions: map[string]any{},
	}

	for k, v := range functions {
		if reflect.TypeOf(v).Kind() == reflect.Func {
			f := reflect.TypeOf(v)
			// checks if function's first input is type of *LinkList[T]
			if f.NumIn() > 0 && f.In(0).Kind() == reflect.Ptr && f.In(0).String() == reflect.TypeOf(l).String() {
				l.functions[k] = v
			} else {
				return nil, errors.New(fmt.Sprintf("`%s` is a function but its first input must be type of %s", k, reflect.TypeOf(l).String()))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("`%s` is not a function", k))
		}
	}

	return l, nil
}

// O(n)
func (l *LinkList[T]) Clear() {
	for trav := l.Head; trav != nil; trav = trav.Next {
		if trav.Prev != nil {
			trav.Prev.Next = nil
		}
		trav.data = nil
		trav.Prev = nil
	}

	l.Head = nil
	l.Tail = nil
	l.size = 0
}

func (l *LinkList[T]) Run(function string, params... any) ([]reflect.Value, error) {
	params = append([]any{l}, params...)
	if v, ok := l.functions[function]; ok {
		return internal.CallJobFuncWithParams(v, params)
	}

	return nil, errors.New(fmt.Sprintf("%s not found", function))
}

func (l *LinkList[T]) oneSizePlus() {
	l.size++
}

func (l *LinkList[T]) oneSizeMinus() {
	l.size--
	if l.size < 0 {
		l.size = 0
	}
}

func (l *LinkList[T]) GetSize() int {
	return l.size
}

func (l *LinkList[T]) IsEmpty() bool {
	return l.size == 0
}

// O(1)
func (l *LinkList[T]) AddLast(data *T) *T {
	var node *Node[T] = nil

	node = &Node[T]{
		data: data,
		Prev: nil,
		Next: nil,
	}

	if l.size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}

	l.oneSizePlus()

	return data
}

// O(1)
func (l *LinkList[T]) AddFirst(data *T) *T {
	var node *Node[T] = nil

	node = &Node[T]{
		data: data,
		Prev: nil,
		Next: nil,
	}

	if l.size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}

	l.oneSizePlus()

	return data
}

func (l *LinkList[T]) addBeforeNode(node *Node[T], data *T) *T {
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

	l.oneSizePlus()

	return data
}

// O(n)
func (l *LinkList[T]) Add(index int, data *T) *T {
	if index == 0 {
		return l.AddFirst(data)
	} else if index == l.size {
		return l.AddLast(data)
	} else if index < 0 || index > l.size {
		return nil
	}

	for trav, i := l.Head, 0; trav != nil; trav, i = trav.Next, i+1 {
		if i == index {
			return l.addBeforeNode(trav, data)
		}
	}

	return nil
}

// O(n)
func (l *LinkList[T]) RemoveData(data *T) *T {
	for trav := l.Head; trav != nil; trav = trav.Next {
		if trav.data == data {
			return l.RemoveNode(trav)
		}
	}

	return nil
}

// O(n)
func (l *LinkList[T]) Remove(index int) *T {
	if index == 0 {
		return l.RemoveFirst()
	} else if index == l.size {
		return l.RemoveLast()
	} else if index < 0 || index > l.size {
		return nil
	}

	for trav, i := l.Head, 0; trav != nil; trav, i = trav.Next, i+1 {
		if i == index {
			return l.RemoveNode(trav)
		}
	}

	return nil
}

// O(1)
func (l *LinkList[T]) RemoveNode(node *Node[T]) *T {
	if node == nil {
		return nil
	}
	
	var data *T = nil

	if l.Head == node {
		return l.RemoveFirst()
	} else if l.Tail == node {
		return l.RemoveLast()
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

	l.oneSizeMinus()

	return data
}

// O(1)
func (l *LinkList[T]) RemoveLast() *T {
	var data *T = nil

	if l.size != 0 {
		if l.Head == l.Tail {
			l.Head = nil
		}
		node := l.Tail
		l.Tail = node.Prev
		if l.Tail != nil {
			l.Tail.Next = nil
		}

		data = node.data
		node.Next = nil
		node.Prev = nil
		node.data = nil
	}

	l.oneSizeMinus()

	return data
}

// O(1)
func (l *LinkList[T]) RemoveFirst() *T {
	var data *T = nil

	if l.size != 0 {
		if l.Head == l.Tail {
			l.Tail = nil
		}
		node := l.Head
		l.Head = node.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}

		data = node.data
		node.Next = nil
		node.Prev = nil
		node.data = nil
	}

	l.oneSizeMinus()

	return data
}

// O(n)
func (l *LinkList[T]) GetIndex(data *T) int {
	for i, trav := 0, l.Head; i < l.size; i, trav = i+1, trav.Next {
		if trav.data == data {
			return i
		}
	}

	return -1
}

// O(n)
func (l *LinkList[T]) Contains(data *T) bool {
	return l.GetIndex(data) != -1
}

// O(n)
func (l *LinkList[T]) GetNode(index int) *Node[T] {
	if index > l.size || index < 0 {
		return nil
	}

	for i, trav := 0, l.Head; i < l.size; i, trav = i+1, trav.Next {
		if i == index {
			return trav
		}
	}

	return nil
}

// O(n)
func (l *LinkList[T]) FindNode(data *T) *Node[T] {
	for trav := l.Head; trav != nil; trav = trav.Next {
		if trav.data == data {
			return trav
		}
	}

	return nil
}

func (l *LinkList[T]) DisplaceTo(node *Node[T], index int) bool {
	if index >= l.size || index < 0 {
		return false
	}

	data := node.data

	l.RemoveNode(node)
	l.Add(index, data)

	return true
}

func (l *LinkList[T]) DisplaceIndex(index1 int, index2 int) bool {
	return l.Displace(l.GetNode(index1), l.GetNode(index2))
}

// O(1)
func (l *LinkList[T]) Displace(node1 *Node[T], node2 *Node[T]) bool {
	if node1 == node2 {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	if node1 == l.Head {
		l.Head = node2
	} else if node2 == l.Head {
		l.Head = node1
	}
	if node2 == l.Tail {
		l.Tail = node1
	} else if node1 == l.Tail {
		l.Tail = node2
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
func (l *LinkList[T]) ReturnAsSlice() []T {
	slice := []T{}
	for trav := l.Head; trav != nil; trav = trav.Next {
		slice = append(slice, *trav.data)
	}

	return slice
}

func (l *LinkList[T]) FirstNode() *Node[T] {
	return l.Head
}

func (l *LinkList[T]) LastNode() *Node[T] {
	return l.Tail
}

func (l *LinkList[T]) First() *T {
	if l.Head != nil {
		return l.Head.data
	}
	return nil
}

func (l *LinkList[T]) Last() *T {
	if l.Head != nil {
		return l.Tail.data
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
