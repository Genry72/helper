package stackquery_test

import (
	"fmt"
	"github.com/Genry72/helper/stackquery"
)

func ExampleStack() {
	stack := stackquery.NewStack[int](0)
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}

	fmt.Println("len: ", stack.Len())

	// Output:
	// 4
	// 3
	// 2
	// 1
	// 0
	// len:  0

}
