package gentestdata

import (
	"container/heap"
	"github.com/Genry72/helper/heapmap"
	"math/rand"
	"slices"
)

type IntHeap[K int, V int] []V

func (h *IntHeap[K, V]) Len() int { return len(*h) }
func (h *IntHeap[K, V]) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap[K, V]) FnCompate() heapmap.ComparerFunc[K] {
	return func(k1, k2 K) bool {
		return k1 < k2
	}
}

func (h *IntHeap[K, V]) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap[K, V]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(V))
}

func (h *IntHeap[K, V]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap[K, V]) Copy() *IntHeap[K, V] {
	cop := slices.Clone(*h)
	return &cop
}

func (h *IntHeap[K, V]) FnKey() heapmap.FnKV[K, V] {
	fn := func(v V) (K, V) { return K(v), v }
	return fn

}

type HeapInterface[K comparable, V any] interface {
	heap.Interface
	Copy() *IntHeap[int, int]
	FnKey() heapmap.FnKV[K, V]
	FnCompate() heapmap.ComparerFunc[K]
}

// InitHeap Добавляет свойства кучи, возвращает новый объект
func InitHeap[K comparable, V any](h HeapInterface[K, V]) heap.Interface {
	res := h.Copy()
	heap.Init(res)
	return res
}

// GetIntSlice Без свойств кучи
func GetIntSlice[K comparable, V any](count int) *IntHeap[int, int] {
	rand.New(rand.NewSource(1))
	res := make(IntHeap[int, int], count)
	for i := 0; i < count; i++ {
		res[i] = rand.Int()
	}
	return &res
}
