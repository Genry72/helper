package priorityquery_test

import (
	"fmt"
	"github.com/Genry72/helper/priorityquery"
)

type Items struct {
	value    string
	priority int
}

func ExampleNewFromSlice() {
	// Cписок продуктов и их приоритеты
	items := []*Items{
		{
			value:    "banana",
			priority: 3,
		},
		{
			value:    "apple",
			priority: 2,
		},
		{
			value:    "pear",
			priority: 4,
		},
	}

	h := priorityquery.NewFromSlice(items, func(a, b *Items) bool {
		return a.priority > b.priority
	})

	h.PushElement(&Items{
		value:    "orange",
		priority: 1,
	})

	// Вычитываем все элементы из кучи
	for h.Len() > 0 {
		item, _ := h.PopElement()
		fmt.Printf("%s:%.2d\n", item.value, item.priority)
	}

	// Output:
	// pear:04
	// banana:03
	// apple:02
	// orange:01
}

func ExampleNewHeap() {
	// Cписок продуктов и их приоритеты
	items := []int{}

	h := priorityquery.NewFromSlice[int](items, func(a, b int) bool {
		return a > b
	})

	for i := 0; i < 10; i++ {
		h.PushElement(i)
	}

	for i, v := range h.Iter() {
		if v > 5 {
			h.DeleteElement(i)
		}
	}

	// Вычитываем все элементы из кучи
	for h.Len() > 0 {
		item, _ := h.PopElement()
		fmt.Printf("%d\n", item)
	}

	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}
