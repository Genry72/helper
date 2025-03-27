package heapmap_test

import (
	"fmt"
	"github.com/Genry72/helper/heapmap"
)

func ExampleHeapMapSlice() {
	type myStruct struct {
		froutName string
		id        int
	}
	// Cписок продуктов и их id
	items := []*myStruct{
		{
			froutName: "banana",
			id:        3,
		},
		{
			froutName: "apple",
			id:        2,
		},
		{
			froutName: "pear",
			id:        4,
		},
	}

	h := heapmap.NewFromSlice[int, *myStruct](items,
		func(v *myStruct) (int, *myStruct) { // Функция для получения ключа из переданной структуры
			return v.id, v
		}, func(k1, k2 int) bool { // Сортировка по id от большего к меньшему
			return k1 > k2
		})

	h.PushElement(5, &myStruct{
		froutName: "orange",
		id:        5,
	})

	// Получение эдемента без удаления из кучи по ключу
	el, _ := h.GetElement(2)
	fmt.Printf("%+v", *el.Value)

	// Удаляем элемент с id =4
	h.DeleteElement(4)

	// Проход по всем элементам, удаляя из кучи
	for h.Len() > 0 {
		val, _ := h.PopElement()
		fmt.Printf("%.2d:%s\n", val.Key, val.Value.froutName)
	}

	// Output:
	// {froutName:apple id:2}05:orange
	// 03:banana
	// 02:apple
}
