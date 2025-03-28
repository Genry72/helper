package heapmap_test

import (
	"container/heap"
	"github.com/Genry72/helper/heapmap"
	"github.com/Genry72/helper/heapmap/gentestdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type myStrct struct {
	name string
	age  int
}

func TestNewFromSliceInts(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name string
	}

	slInts := gentestdata.GetIntSlice[int, int](100)

	tests := []testCase[int, int]{
		{
			name: "#1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := heapmap.NewFromSlice[int, int](*slInts, slInts.FnKey(), slInts.FnCompate())
			hres := gentestdata.InitHeap(slInts)
			if got.Len() != hres.Len() {
				t.Fatal("lenGot != lenHres")
			}
			// Проверка инициадизации и корректности заполнения мапы с индексами
			hre := *hres.(*gentestdata.IntHeap[int, int])
			for i := range hre {
				idx, ok := got.FindIdxByKey(hre[i])
				if !ok {
					t.Fatal("не найден ключ")
				}
				if i != idx {
					t.Fatal("не верный индекс")
				}
			}

			// Проверка результирующего массива
			hresSl := make([]int, 0)

			gotSl := make([]int, 0)

			for hres.Len() != 0 {
				hresSl = append(hresSl, heap.Pop(hres).(int))
			}

			for got.Len() != 0 {
				v, _ := got.PopElement()
				gotSl = append(gotSl, v.Value)
			}
			assert.Equal(t, hresSl, gotSl)
		})
	}
}

func TestOtherStruct(t *testing.T) {
	sl := []myStrct{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Ключ сортировки имя (строка)
	h := heapmap.NewFromSlice[string, myStrct](sl, func(v myStrct) (string, myStrct) {
		return v.name, v
	}, func(k1, k2 string) bool {
		// Обратный порядок сортировки
		return k1 > k2
	})

	h.DeleteElement("Charlie")

	h.PushElement("Willi", myStrct{"Willi", 28})

	resultSl := []myStrct{}

	for h.Len() != 0 {
		v, _ := h.PopElement()
		resultSl = append(resultSl, v.Value)
	}

	assert.Equal(t, []myStrct{
		{
			name: "Willi",
			age:  28,
		},
		{
			name: "Bob",
			age:  25,
		},
		{
			name: "Alice",
			age:  30,
		},
	}, resultSl)

	h2 := heapmap.NewFromSlice[int, myStrct](sl, func(v myStrct) (int, myStrct) {
		return v.age, v
	}, func(k1, k2 int) bool {
		// По убыванию возраста
		return k1 > k2
	})
	// Так как ключом является возраст, то установив новое значение для ключа 30, мы меняем имя
	h2.PushElement(30, struct {
		name string
		age  int
	}{name: "Marina", age: 30})

	resultSl = []myStrct{}

	for h2.Len() != 0 {
		v, _ := h2.PopElement()
		resultSl = append(resultSl, v.Value)
	}

	assert.Equal(t, []myStrct{
		{"Charlie", 35},
		{"Marina", 30},
		{"Bob", 25},
	}, resultSl)
}

func TestHeapMapDelete_ForEach(t *testing.T) {
	sl := []myStrct{
		{"Charlie", 35},
		{"Marina", 30},
		{"Bob", 25},
	}

	getHeap := func() *heapmap.HeapMap[int, myStrct] {
		return heapmap.NewFromSlice[int, myStrct](sl, func(v myStrct) (int, myStrct) {
			return v.age, v
		}, func(k1, k2 int) bool {
			return k1 < k2
		})
	}

	type args struct {
		fnDelete func(int, *heapmap.HeapMap[int, myStrct])
	}

	type testCase struct {
		name string
		args args
		want []myStrct
	}
	tests := []testCase{
		{
			name: "delete all",
			args: args{
				fnDelete: func(i int, v *heapmap.HeapMap[int, myStrct]) {
					v.DeleteElement(i)
				},
			},
			want: []myStrct{},
		},
		{
			name: "delete Marina",
			args: args{
				fnDelete: func(i int, v *heapmap.HeapMap[int, myStrct]) {
					if i == 30 {
						v.DeleteElement(i)
					}
				},
			},
			want: []myStrct{
				{"Bob", 25},
				{"Charlie", 35},
			},
		},
		{
			name: "delete Marina and Bob",
			args: args{
				fnDelete: func(i int, v *heapmap.HeapMap[int, myStrct]) {
					if i == 30 || i == 25 {
						v.DeleteElement(i)
					}
				},
			},
			want: []myStrct{
				{"Charlie", 35},
			},
		},
		{
			name: "просто итерация",
			args: args{
				fnDelete: func(i int, v *heapmap.HeapMap[int, myStrct]) {},
			},
			want: []myStrct{
				{"Bob", 25},
				{"Marina", 30},
				{"Charlie", 35},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getHeap()
			for k, v := range h.Iter() {
				_ = v
				tt.args.fnDelete(k, h)
			}
			actual := make([]myStrct, 0)
			for h.Len() != 0 {
				v, _ := h.PopElement()
				actual = append(actual, v.Value)
			}
			assert.Equal(t, tt.want, actual)
		})
	}
}
