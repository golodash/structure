package array

import (
	"fmt"
	"errors"
	"reflect"

	"github.com/golodash/structure/internal"
)

type (
	Array[T any] struct {
		Values []*T
		size int
		functions map[string]any
	}
)

func New[T any](functions map[string]any) (*Array[T], error) {
	a := &Array[T]{
		Values: []*T{},
		size: 0,
		functions: map[string]any{},
	}

	for k, v := range functions {
		if reflect.TypeOf(v).Kind() == reflect.Func {
			f := reflect.TypeOf(v)
			// checks if function's first input is type of *Array[T]
			if f.NumIn() > 0 && f.In(0).Kind() == reflect.Ptr && f.In(0).String() == reflect.TypeOf(a).String() {
				a.functions[k] = v
			} else {
				return nil, errors.New(fmt.Sprintf("`%s` is a function but its first input must be type of %s", k, reflect.TypeOf(a).String()))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("`%s` is not a function", k))
		}
	}

	return a, nil
}

func (a *Array[T]) Clear() {
	a.Values = []*T{}
	a.size = 0
}

func (a *Array[T]) Run(function string, params... any) ([]reflect.Value, error) {
	params = append([]any{a}, params...)
	if v, ok := a.functions[function]; ok {
		return internal.CallJobFuncWithParams(v, params)
	}

	return nil, errors.New(fmt.Sprintf("%s not found", function))
}

func (a *Array[T]) oneSizePlus() {
	a.size++
}

func (a *Array[T]) oneSizeMinus() {
	a.size--
	if a.size < 0 {
		a.size = 0
	}
}

func (a *Array[T]) PeekFirst() *T {
	if a.size == 0 {
		return nil
	}

	return a.Values[0]
}

func (a *Array[T]) PeekLast() *T {
	if a.size == 0 {
		return nil
	}

	return a.Values[a.size-1]
}

func (a *Array[T]) Peek(index int) *T {
	if index < 0 || index >= a.size {
		return nil
	}

	return a.Values[index]
}

func (a *Array[T]) ReturnAsSlice() []T {
	values := []T{}
	for i := range a.Values {
		values = append(values, *a.Values[i])
	}

	return values
}

func (a *Array[T]) PushLast(data *T) *T {
	a.Values = append(a.Values, data)
	a.oneSizePlus()
	return data
}

func (a *Array[T]) PushFirst(data *T) *T {
	a.Values = append([]*T{data}, a.Values...)
	a.oneSizePlus()
	return data
}

func (a *Array[T]) Push(data *T, index int) *T {
	if index < 0 || index > a.size {
		return nil
	}

	if index == 0 {
		return a.PushFirst(data)
	} else if index == a.size {
		return a.PushLast(data)
	}

	a.Values = append(a.Values[:index+1], a.Values[index:]...)
	a.Values[index] = data
	a.oneSizePlus()
	return data
}

func (a *Array[T]) PopLast() *T {
	if a.size == 0 {
		return nil
	}

	a.oneSizeMinus()
	data := a.Values[a.size-1]
	a.Values = a.Values[:a.size]
	return data
}

func (a *Array[T]) PopFirst() *T {
	if a.size == 0 {
		return nil
	}

	a.oneSizeMinus()
	data := a.Values[0]
	a.Values = a.Values[1:]
	return data
}

func (a *Array[T]) Pop(index int) *T {
	if index < 0 || index >= a.size {
		return nil
	}

	if index == 0 {
		return a.PopFirst()
	} else if index == a.size-1 {
		return a.PopLast()
	}

	data := a.Values[index]
	a.Values = append(a.Values[:index], a.Values[index+1:]...)
	a.oneSizeMinus()
	return data
}

func (a *Array[T]) Displace(index1 int, index2 int) bool {
	if index1 < 0 || index1 >= a.size || index2 < 0 || index2 >= a.size {
		return false
	}

	v1 := a.Values[index1]
	v2 := a.Values[index2]
	
	a.ReplaceValue(index1, v2)
	a.ReplaceValue(index2, v1)
	
	return true
}

func (a *Array[T]) ReplaceValue(index int, data *T) bool {
	if index < 0 || index >= a.size {
		return false
	}

	a.Values[index] = data

	return true
}

func (a *Array[T]) GetSize() int {
	return a.size
}
