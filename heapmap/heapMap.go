package heapmap

import (
	"container/heap"
	"maps"
	"slices"
	"sort"
)

/*
HeapMap - это структура данных, которая сочетает в себе свойства кучи и хеш-таблицы.
Она позволяет эффективно управлять элементами, обеспечивая быстрый доступ, вставку и удаление элементов по ключу,
а также поддержание порядка элементов согласно заданной функции сравнения ключей.

Используем, когда уверены что выбранный ключ уникален и по нему потребуется быстрый поиск, а так же важен порядок
элементов.
*/
type HeapMap[K comparable, V any] struct {
	heap     []*Item[K, V]
	m        map[K]int // Значение - индекс в heap
	less     ComparerFunc[K]
	isSorted bool // Флаг для избежания повторных сортировок при вызовах ForEach
}

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

// ComparerFunc Функция сравнения ключей для кучи.
type ComparerFunc[K comparable] func(k1, k2 K) bool

// FnKV Функция преобразования элемента в пару ключ-значение. Используется при создании кучи из среза.
type FnKV[K comparable, V any] func(v V) (K, V)

// NewHeapMap Создание кучи с заданной емкостью и функцией сравнения ключей
func NewHeapMap[K comparable, V any](count int, less ComparerFunc[K]) *HeapMap[K, V] {
	sl := &HeapMap[K, V]{
		heap: make([]*Item[K, V], 0, count),
		less: less,
		m:    make(map[K]int, count),
	}
	return sl
}

// NewFromSlice Создание кучи из среза
// fnKV - функция преобразования элемента среза в пару ключ-значение
// less - функция сравнения ключей для создания кучи минимумов/максимумов
// Пример использования смотри в примерах
func NewFromSlice[K comparable, V any](sl []V, fnKV FnKV[K, V], less ComparerFunc[K]) *HeapMap[K, V] {
	res := NewHeapMap[K, V](len(sl), less)
	for i := range sl {
		k, v := fnKV(sl[i])
		res.Push(&Item[K, V]{
			Key:   k,
			Value: v,
		})
	}
	res.Init()
	return res
}

// Len Реализация интерфейса heap.Interface. Возвращает количество элементов в куче.
func (h *HeapMap[K, V]) Len() int { return len(h.heap) }

// Less Реализация интерфейса heap.Interface. Сравнивает элементы по ключам.
func (h *HeapMap[K, V]) Less(i, j int) bool {
	return h.less(h.heap[i].Key, h.heap[j].Key)
}

// Swap Реализация интерфейса heap.Interface. Меняет местами элементы и обновляет индексы в карте.
func (h *HeapMap[K, V]) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]

	h.m[h.heap[i].Key] = i
	h.m[h.heap[j].Key] = j

	h.isSorted = false
}

// Push Реализация интерфейса heap.Interface. Просто добавляет элемент. Без сортировки.
func (h *HeapMap[K, V]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	pr := x.(*Item[K, V])
	if n, found := h.FindIdxByKey(pr.Key); found {
		h.heap[n].Value = pr.Value
		return
	}

	(*h).heap = append((*h).heap, pr)
	// Добавляем в мапу
	h.m[pr.Key] = h.Len() - 1
}

// Init Восстанавливает кучу, если она была изменена. Не делает ничего, если куча уже отсортирована.
func (h *HeapMap[K, V]) Init() {
	if !h.isSorted {
		heap.Init(h)
	}
}

// Pop Реализация интерфейса heap.Interface. Возвращает первый элемент из кучи. И удаляет его. Не восстанавливает кучу
func (h *HeapMap[K, V]) Pop() any {
	old := (*h).heap
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	(*h).heap = old[0 : n-1]
	// Удаляем из мапы
	delete(h.m, x.Key)
	return x
}

// PushElement Добавление элемента, восстановление кучи
func (h *HeapMap[K, V]) PushElement(k K, v V) {
	heap.Push(h, &Item[K, V]{Key: k, Value: v})
}

// PopElement Возвращает первый элемент из кучи. И удаляет его. Восстановление свойства кучи
func (h *HeapMap[K, V]) PopElement() (*Item[K, V], bool) {
	if h.Len() == 0 {
		return nil, false
	}

	el := heap.Pop(h)
	return el.(*Item[K, V]), true
}

// FindIdxByKey Поиск индекса в куче по ключу
func (h *HeapMap[K, V]) FindIdxByKey(key K) (int, bool) {
	if idx, ok := h.m[key]; ok {
		return idx, true
	}
	return -1, false
}

// DeleteElement Удаление элемента по ключу
func (h *HeapMap[K, V]) DeleteElement(key K) {
	if idx, ok := h.FindIdxByKey(key); ok {
		heap.Remove(h, idx)
	}
}

// Fix Восстановление свойства кучи по индесу
func (h *HeapMap[K, V]) Fix(i int) {
	heap.Fix(h, i)
}

// Iter Итерация по всем элементам. Сортировка не гарантируется. Вызоаите дополнительно метод Sort либо
// используйте PopElement
func (h *HeapMap[K, V]) Iter() func(yield func(K, V) bool) {
	oldLen := h.Len()

	return func(yield func(K, V) bool) {
		for i := 0; i < h.Len(); i++ {
			if !yield(h.heap[i].Key, h.heap[i].Value) {
				return
			}

			// В случае удаления элемента
			if oldLen != h.Len() {
				oldLen = h.Len()
				i--
			}
		}
	}
}

// GetElement Возвращение элемента по ключу без его удаления
func (h *HeapMap[K, V]) GetElement(key K) (*Item[K, V], bool) {
	if idx, ok := h.FindIdxByKey(key); ok {
		return (*h).heap[idx], true
	}

	return nil, false
}

// Clone Создание копии кучи
func (h *HeapMap[K, V]) Clone() *HeapMap[K, V] {
	newH := NewHeapMap[K, V](h.Len(), h.less)
	newH.heap = slices.Clone(h.heap)
	maps.Copy(newH.m, h.m)
	return newH
	//return h
}

// Sort Принудительная сортировка всех элементов кучи
func (h *HeapMap[K, V]) Sort() {
	// Нужна именно сортировка из пакета sort, так как здесь задействуется метод swap структуры
	if !h.isSorted {
		sort.Sort(h)
		h.isSorted = true
	}
}
