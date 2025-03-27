package heapslice_test

import (
	"fmt"
	"github.com/Genry72/helper/heapslice"
)

type Items struct {
	value    string
	priority int
}

func ExampleHeapFromSlice() {
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

	h := heapslice.NewFromSlice[*Items](items, func(a, b *Items) bool {
		return a.priority > b.priority
	})

	h.PushElement(&Items{
		value:    "orange",
		priority: 1,
	})

	// Вычитываем все элементы из кучи и записываем в новую
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
