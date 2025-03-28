package heapmap_test

import (
	"fmt"
	"github.com/Genry72/helper/heapmap"
)

func ExampleHeapMap() {
	// Cписок продуктов и их id
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	h := heapmap.NewHeapMap[string, int](len(items), func(k1, k2 string) bool {
		// Сортировка по имени
		return k1 < k2
	})
	for name, id := range items {
		h.PushElement(name, id)
	}

	h.PushElement("orange", 5)

	el, _ := h.GetElement("apple")
	fmt.Println(el.Key, el.Value)

	// Проход по всем элементам, оставляя в куче
	for h.Len() > 0 {
		item, _ := h.PopElement()
		fmt.Printf("%s:%.2d\n", item.Key, item.Value)
	}

	fmt.Println(h.Len())

	// Output:
	// apple 2
	// apple:02
	// banana:03
	// orange:05
	// pear:04
	// 0
}
