package heapslice

import (
	"container/heap"
	"slices"
	"sort"
)

type Heap[V any] struct {
	heap     []V
	less     ComparerFunc[V]
	isSorted bool // Флаг для избежания повторных сортировок при вызовах ForEach
}

func (h *Heap[V]) Len() int           { return len(h.heap) }
func (h *Heap[V]) Less(i, j int) bool { return h.less(h.heap[i], h.heap[j]) }
func (h *Heap[V]) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]

	h.isSorted = false
}

type ComparerFunc[V any] func(a, b V) bool

func NewHeap[V any](count int, less ComparerFunc[V]) *Heap[V] {
	sl := &Heap[V]{
		heap: make([]V, 0, count),
		less: less,
	}
	heap.Init(sl)

	return sl
}

func NewFromSlice[V any](sl []V, less ComparerFunc[V]) *Heap[V] {
	res := NewHeap(0, less)
	res.heap = sl
	heap.Init(res)

	return res
}

func (h *Heap[V]) Push(v any) {
	pr := v.(V)
	(*h).heap = append((*h).heap, pr)
}

// Init Восстанавливает кучу, если она была изменена. Не делает ничего, если куча уже отсортирована.
func (h *Heap[V]) Init() {
	if !h.isSorted {
		heap.Init(h)
	}
}

// PopElement Возвращает первый элемент из кучи. И удаляет его. Восстановление свойства кучи
func (h *Heap[V]) PopElement() (V, bool) {
	if h.Len() == 0 {
		var zero V
		return zero, false
	}

	el := heap.Pop(h)
	return el.(V), true
}

// PushElement Добавление элемента, восстановление кучи
func (h *Heap[V]) PushElement(v V) {
	heap.Push(h, v)
}

// Pop Реализация интерфейса heap.Interface. Возвращает первый элемент из кучи. И удаляет его. Не восстанавливает кучу
func (h *Heap[V]) Pop() any {
	old := (*h).heap
	n := len(old)
	x := old[n-1]
	(*h).heap = old[0 : n-1]
	return x
}

// ForEach Итерация согласно порядку Less. Происходит сортировка всего массива функцией Less
func (h *Heap[V]) ForEach() func(yield func(int, V) bool) {
	h.Sort()
	return func(yield func(int, V) bool) {
		for i := range h.heap {
			if !yield(i, h.heap[i]) {
				return
			}
		}
	}
}

// Iter итерация по массиву, сортировка не гарантирована
func (h *Heap[V]) Iter() func(yield func(int, V) bool) {
	return func(yield func(int, V) bool) {
		for i := range h.heap {
			if !yield(i, h.heap[i]) {
				return
			}
		}
	}
}

/*
BinarySearch Убедитесь что куча отсортирована в порядке возрастания (либо измените приведенный пример ниже),
по нужному ключу (вызовите метод Sort) иначе воспользуйтесь линейным поиском используя Iter, либо отсортируйте массив по другому принципу, используя метод SortByOtherFn

	if a == b {
		return 0
	}

	if a < b {
		return -1
	}

return 1
*/
func (h *Heap[V]) BinarySearch(target V, compareFn func(a V, b V) int) func(yield func(int, bool) bool) {
	return func(yield func(int, bool) bool) {
		for i := 0; i < h.Len(); {
			nextIndex, found := slices.BinarySearchFunc(h.heap[i:], target, compareFn)
			if !found {
				if i == 0 {
					yield(-1, false)
				}
				return
			}

			if !yield(nextIndex+i, true) {
				return
			}

			i = nextIndex + i + 1
		}
	}
}

// DeleteElement Удаление элемента по индексу
func (h *Heap[V]) DeleteElement(idx int) {
	if idx < h.Len() && idx > 0 {
		heap.Remove(h, idx)
	}
}

// Fix Восстановление свойства кучи по индесу
func (h *Heap[V]) Fix(i int) {
	heap.Fix(h, i)
}

// SetByIndex Устанавливает новое значение
func (h *Heap[V]) SetByIndex(idx int, v V) {
	h.heap[idx] = v
	h.Fix(idx)
}

func (h *Heap[V]) GetByIndex(idx int) V {
	return h.heap[idx]
}

// Sort Принудительная сортировка всех элементов кучи
func (h *Heap[V]) Sort() {
	// Нужна именно сортировка из пакета sort, так как здесь задействуется метод swap структуры
	if !h.isSorted {
		sort.Sort(h)
		h.isSorted = true
	}
}

// SortByOtherFn Сортировка массива по отличной, переданной сортировки при инициализации кучи
// Когда данная сортировка будет не нужна, нужно вызвать метод Init, для восстановления состояния кучи
func (h *Heap[V]) SortByOtherFn(sortFn ComparerFunc[V]) {
	sort.Slice(h.heap, func(i, j int) bool {
		return sortFn(h.heap[i], h.heap[j])
	})
	h.isSorted = false
}
