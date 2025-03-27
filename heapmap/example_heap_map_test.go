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
	for k, v := range h.ForEach() {
		fmt.Printf("%s:%.2d ", k, v)
	}

	// Output:
	// apple 2
	// apple:02 banana:03 orange:05 pear:04
}
