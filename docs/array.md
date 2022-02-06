# Array (Stack, Queue or Array)

## Different Variables:

In Array implementation there is one major structure:

1. **Array**[T any] struct:
   1. **Values** []*T
      - Represents the actual values inside array struct.
   2. **size** int (inaccessible)
      - Represents size of the `Values` array.
   3. **functions** map[string]any (inaccessible)
      - Represents custom functions that you want to run on `Array` object by `Run` function.

## Functions

1. **New**[T any](functions map[string]any) (*Array[T], error):
   - Creates and returns a `Array` object.
   - Adds functions inside the `Array` object that can be run from `Run` function.
   - **Important Note**: Functions that you send inside `New` function, have to have their first input as type of `*Array[T]`.
2. **Array**[T any] struct:
   1. **Run**(function string, params... any) ([]reflect.Value, error):
      - Receives function key as string and some parameters and runs the function with given parameters.
   2. **Clear**():
      - Clears the array variable(`Values`) and sets `size` to 0.
   3. **PeekFirst**() *T:
      - Returns the first element of array.
   4. **PeekLast**() *T:
      - Returns the last element of array.
   5. **Peek**(index int) *T:
      - Returns index'th element of array.
   6. **PopFirst**() *T:
      - Pops and returns the first element of array.
   7. **PopLast**() *T:
      - Pops and returns the last element of array.
   8. **Pop**(index int) *T:
      - Pops and returns index'th element of array.
   9.  **PushFirst**(data *T) *T:
       - Adds the given data to the beginning of array.
   10. **PushLast**(data *T) *T:
       - Adds the given data to the end of array.
   11. **Push**(data *T, index int) *T:
       - Adds the given data to index'th place of array.
   12. **GetSize**() int:
       - Returns length of the array.
   13. **ReturnAsSlice**() []T:
       - Returns a **copy** of the actual array(`Values` variable) inside `Array` object.
   14. **ReplaceValue**(index int, data *T) bool:
       - Replaces the value of the index element of the array with the given data.
   15. **Displace**(index1 int, index2 int) bool:
       - Changes the value of the two given indexes to each other in array variable.

## Example Usage

This is just a random simple code of how to use this structure:

```go
package main

import (
	"fmt"
	// "reflect"

	"github.com/golodash/structure/array"
)

type Human struct {
	FirstName string
	LastName  string
	IDNum     string
}

func fun(a *array.Array[Human], i int) []Human {
	fmt.Println("We are inside `fun` function and i value equals to: ", i, "\n")
	return a.ReturnAsSlice()
}

func main() {
	a, err := array.New[Human](map[string]any{
		"fun": fun,
	})
	if err != nil {
		panic(err.Error())
	}

	a.Push(&Human{
		FirstName: "Mahmood",
		LastName: "Heidari",
		IDNum: "0860215489",
	}, 0)
	a.Push(&Human{
		FirstName: "Ali",
		LastName: "Ahmadi",
		IDNum: "0860211245",
	}, 0)
	a.Push(&Human{
		FirstName: "Ahmad",
		LastName: "Mostofi",
		IDNum: "0860651245",
	}, 1)
	a.PushLast(&Human{
		FirstName: "Gholam",
		LastName: "Kholi",
		IDNum: "0860659845",
	})
	a.PushLast(&Human{
		FirstName: "GholamAli",
		LastName: "Akbari",
		IDNum: "0869859845",
	})
	a.PushFirst(&Human{
		FirstName: "Karim",
		LastName: "Khavari",
		IDNum: "9869859845",
	})

	values, err := a.Run("fun", 1)
	if err == nil {
		fmt.Println(values[0])
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println(a.GetSize())
	a.Displace(1, 2)
	a.Pop(0)
	fmt.Println(a.ReturnAsSlice())
	a.Displace(0, 3)
	fmt.Println(a.ReturnAsSlice())
	a.Pop(2)
	fmt.Println(a.ReturnAsSlice())
	fmt.Println(a.GetSize())
}
```

Output:
```bash
We are inside `fun` function and i value equals to:  1 

[{Karim Khavari 9869859845} {Ali Ahmadi 0860211245} {Ahmad Mostofi 0860651245} {Mahmood Heidari 0860215489} {Gholam Kholi 0860659845} {GholamAli Akbari 0869859845}]
6
[{Ahmad Mostofi 0860651245} {Ali Ahmadi 0860211245} {Mahmood Heidari 0860215489} {Gholam Kholi 0860659845} {GholamAli Akbari 0869859845}]
[{Gholam Kholi 0860659845} {Ali Ahmadi 0860211245} {Mahmood Heidari 0860215489} {Ahmad Mostofi 0860651245} {GholamAli Akbari 0869859845}]
[{Gholam Kholi 0860659845} {Ali Ahmadi 0860211245} {Ahmad Mostofi 0860651245} {GholamAli Akbari 0869859845}]
4
```
