package priorityquery_test

import (
	"fmt"
	"github.com/Genry72/helper/heapslice"
	"strings"
)

type Item struct {
	value    string
	priority int
}

func ExampleHeap() {
	// Cписок продуктов и их приоритеты
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	h := priorityquery.NewHeap[Item](len(items), func(a, b Item) bool {
		return a.priority > b.priority
	})
	for name, priority := range items {
		h.PushElement(Item{
			value:    name,
			priority: priority,
		})
	}

	h.PushElement(Item{
		value:    "orange",
		priority: 1,
	})

	h2 := priorityquery.NewHeap[Item](h.Len(), func(a, b Item) bool {
		return a.priority > b.priority
	})

	// Вычитываем все элементы из кучи и записываем в новую
	for h.Len() > 0 {
		item, _ := h.PopElement()
		h2.PushElement(item)
		fmt.Printf("%s:%.2d\n", item.value, item.priority)
	}

	fmt.Println("len(h) = ", h.Len())

	// Так как куча отсортирована по значению приоритета, сортируем по имени для выполнения бинарного поиска
	h2.SortByOtherFn(func(a, b Item) bool {
		return a.value < b.value
	})

	var orangeIdx int

	for idx, ok := range h2.BinarySearch(Item{value: "orange"}, func(a Item, b Item) int {
		return strings.Compare(a.value, b.value)
	}) {
		if !ok {
			panic("orange not fount") // !!! не использовать в проде)
		}
		orangeIdx = idx
	}

	h2.SetByIndex(orangeIdx, Item{
		value:    "orange",
		priority: 100,
	})

	// Применяем исходную сортировку
	h2.Sort()

	// Проход по всем элементам, оставляя в куче. Здесь сортировка на гарантируется
	for _, v := range h2.Iter() {
		fmt.Printf("%s:%.2d\n", v.value, v.priority)
	}

	// Output:
	// pear:04
	// banana:03
	// apple:02
	// orange:01
	// len(h) =  0
	// orange:100
	// pear:04
	// banana:03
	// apple:02
}
