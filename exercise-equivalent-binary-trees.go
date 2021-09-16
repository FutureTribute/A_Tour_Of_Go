// concurrency/8
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	"reflect"  // Google search "how to compare two slices"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <-t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	values1 := make([]int, 10, 10)
	values2 := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		values1[i] = <-ch1
		values2[i] = <-ch2
	}
	fmt.Println(values1)
	fmt.Println(values2)
	return reflect.DeepEqual(values1, values2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
