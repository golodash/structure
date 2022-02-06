# LinkList

## Different Variables:

In LinkList implementation there are two major structure:

1. **Node**[T any] struct:
   1. **Prev** *Node[T]:
      - Represents previous node before the current node.
   2. **Next** *Node[T]:
      - Represents next node after the current node.
   3. **data** *T (inaccessible):
      - Represent the actual data stored inside the node.
2. **LinkList**[T any] struct:
   1. **Head** *Node[T]:
      - Represents start of the nodes.
   2. **Tail** *Node[T]:
      - Represents end of the nodes.
   3. **Size** int:
      - Represents size of the nodes.
   4. **functions** map[string]any (inaccessible):
      - Represents custom functions that you want to run on `LinkList` object by `Run` function.

## Functions

1. **New**[T any](functions map[string]any) (*LinkList[T], error):
   - Creates and returns a `LinkList` object.
   - Adds functions inside the `LinkList` object that can be run from `Run` function.
   - **Important Note**: Functions that you send inside `New` function, have to have their first input as type of `*LinkList[T]`.
2. **LinkList**[T any] struct:
   1. **Run**(function string, params... any) ([]reflect.Value, error):
      - Receives function key as string and some parameters and runs the function with given parameters.
   2. **Clear**():
      - Clears the nodes.
   3. **GetSize**() int:
      - Returns length of all nodes.
   4. **IsEmpty**() bool:
      - Return true of the `LinkList` object is empty.
   5. **AddLast**(data *T) *T:
      - Adds a node at the end of `LinkList` object.
   6. **AddFirst**(data *T) *T:
      - Adds a node at the beginning of `LinkList` object.
   7. **Add**(index int, data *T) *T:
      - Adds a node at desired index location.
   8.  **RemoveData**(data *T) *T:
      - Removes a node that has same data address as variable `data` has in memory. (***addresses*** of them will be compared ***not values***)
   9.  **Remove**(index int) *T:
      - Removes a node based on its index from `LinkList` object.
   10. **RemoveNode**(node *Node[T]) *T:
      - Removes a node from `LinkList` object.
   11. **RemoveLast**() *T:
      - Removes latest node from `LinkList` object.
   12. **RemoveFirst**() *T:
       - Removes first node from `LinkList` object.
   13. **GetIndex**(data *T) int:
       - Returns index of the node that has asked data in it. (if no nodes found, -1 returns)
   14. **Contains**(data *T) bool:
       - Returns true if a node containing asked data exists.
   15. **GetNode**(index int) *Node[T]:
       - Returns a node based on given index.
   16. **FindNode**(data *T) *Node[T]:
       - Returns a node based on data address in memory (not its value).
   17. **Displace**(node1 *Node[T], node2 *Node[T]) bool:
       - Changes place of the two given nodes to each other.
   18. **DisplaceIndex**(index1 int, index2 int) bool:
       - Changes place of the two given nodes in those indexes to each other.
   19. **DisplaceTo**(node *Node[T], index int) bool:
       - Changes place of a node to `index` position.
   20. **ReturnAsSlice**() []T:
       - Returns a **copy** of `LinkList` object nodes inside an slice.
   21. **FirstNode**() *Node[T]:
       - Returns first node of `LinkList` object.
   22. **LastNode**() *Node[T]:
       - Returns last node of `LinkList` object.
   23. **First**() *Node[T]:
       - Returns first `data` of the first node in `LinkList` object.
   24. **Last**() *Node[T]:
       - Returns last `data` of the last node in `LinkList` object.
3. **Node**[T any] struct:
   1. **GetData**() *T:
       - Returns data of the node.
   2. **SetData**(data *T) *T:
       - Replaces data of the node with new given data.

## Example Usage

This is just a random simple code of how to use this structure:

```go
package main

import (
	"fmt"
	// "reflect"

	"github.com/golodash/structure/ll"
)

type Human struct {
	FirstName string
	LastName  string
	IDNum     string
}

func fun(l *ll.LinkList[Human], i int) []Human {
	fmt.Println("We are inside `fun` function and i value equals to: ", i, "\n")
	return l.ReturnAsSlice()
}

func main() {
	l, err := ll.New[Human](map[string]any{
		"fun": fun,
	})
	if err != nil {
		panic(err)
	}
	l.AddLast(&Human{
		FirstName: "Ali",
		LastName:  "Hamidi",
		IDNum:     "0891265769",
	})
	l.AddLast(&Human{
		FirstName: "Jafar",
		LastName:  "Ahmadi",
		IDNum:     "0090565755",
	})
	l.AddFirst(&Human{
		FirstName: "Gholi",
		LastName:  "Ghasemi",
		IDNum:     "0890254698",
	})
	l.AddFirst(&Human{
		FirstName: "Ahmad",
		LastName:  "Abbasi",
		IDNum:     "0750254694",
	})
	l.Add(2, &Human{
		FirstName: "Bob",
		LastName:  "Ross",
		IDNum:     "0320254644",
	})

	values, err := l.Run("fun", 2)
	if err == nil {
		fmt.Println("Returned Value:", values[0])
	} else {
		fmt.Println(err)
	}
	fmt.Println("Removed Item: ", l.Remove(1))
	fmt.Println("After Remove: ", l.ReturnAsSlice())
	fmt.Println("Size: ", l.GetSize())
	l.DisplaceIndex(1, 1)
	l.DisplaceIndex(2, 3)
	l.DisplaceIndex(1, 500)
	fmt.Println("After Displaces: ", l.ReturnAsSlice())
	l.Clear()
	fmt.Println("After Clear: ", l.ReturnAsSlice())
	fmt.Println("Size After Clear: ", l.GetSize())
}
```

Output:
```bash
We are inside `fun` function and i value equals to:  2 

Returned Value: [{Ahmad Abbasi 0750254694} {Gholi Ghasemi 0890254698} {Bob Ross 0320254644} {Ali Hamidi 0891265769} {Jafar Ahmadi 0090565755}]
Removed Item:  &{Gholi Ghasemi 0890254698}
After Remove:  [{Ahmad Abbasi 0750254694} {Bob Ross 0320254644} {Ali Hamidi 0891265769} {Jafar Ahmadi 0090565755}]
Size:  4
After Displaces:  [{Ahmad Abbasi 0750254694} {Bob Ross 0320254644} {Jafar Ahmadi 0090565755} {Ali Hamidi 0891265769}]
After Clear:  []
Size After Clear:  0
```