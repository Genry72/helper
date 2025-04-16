package stackquery_test

import (
	"fmt"
	"github.com/Genry72/helper/stackquery"
)

func ExampleQuery() {
	query := stackquery.NewQuery[int](0)
	for i := 0; i < 5; i++ {
		query.Push(i)
	}
	for query.Len() > 0 {
		fmt.Println(query.Pop())
	}

	fmt.Println("len: ", query.Len())

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// len:  0

}
