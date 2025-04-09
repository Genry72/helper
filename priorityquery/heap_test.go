package priorityquery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStr struct {
	name     string
	priority int
}

func TestHeap_BinarySearch(t *testing.T) {
	type args[V any] struct {
		target    V
		compareFn func(a V, b V) int
	}
	type testCase[V any] struct {
		name string
		h    Heap[V]
		args args[V]
		want []int
	}

	tests := []testCase[testStr]{
		{
			name: "Должны вернуться все элементы",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 0,
					},
					{
						name:     "1",
						priority: 0,
					},
					{
						name:     "1",
						priority: 0,
					},
				},
			},
			args: args[testStr]{
				target: testStr{"1", 0},
				compareFn: func(a testStr, b testStr) int {
					return a.priority - b.priority
				},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "Должны вернуться idx 1'",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 3,
					},
				},
			},
			args: args[testStr]{
				target: testStr{"", 2},
				compareFn: func(a testStr, b testStr) int {
					return a.priority - b.priority
				},
			},
			want: []int{1},
		},
		{
			name: "Должны вернуться idx 1 и 2'",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 2,
					},
				},
			},
			args: args[testStr]{
				target: testStr{"", 2},
				compareFn: func(a testStr, b testStr) int {
					return a.priority - b.priority
				},
			},
			want: []int{1, 2},
		},
		{
			name: "negative #1",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 2,
					},
				},
			},
			args: args[testStr]{
				target: testStr{"", 4},
				compareFn: func(a testStr, b testStr) int {
					return a.priority - b.priority
				},
			},
			want: []int{},
		},
		{
			name: "negative #2",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 2,
					},
				},
			},
			args: args[testStr]{
				target: testStr{"", 0},
				compareFn: func(a testStr, b testStr) int {
					return a.priority - b.priority
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []int{}
			for idx, ok := range tt.h.BinarySearch(tt.args.target, tt.args.compareFn) {
				if !ok && len(tt.want) != 0 {
					panic("not found")
				} else {
					if ok {
						res = append(res, idx)
					}
				}

			}
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestHeap_DeleteElement(t *testing.T) {
	type args struct {
		idx int
	}
	type testCase[V any] struct {
		name string
		h    Heap[V]
		args args
		want []V
	}
	tests := []testCase[testStr]{
		{
			name: "#1",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 2,
					},
				},
				less: func(a, b testStr) bool {
					return a.priority < b.priority
				},
			},
			args: args{
				idx: 1,
			},
			want: []testStr{
				{
					name:     "1",
					priority: 1,
				},
				{
					name:     "1",
					priority: 2,
				},
			},
		},
		{
			name: "#2",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 2,
					},
				},
				less: func(a, b testStr) bool {
					return a.priority < b.priority
				},
			},
			args: args{
				idx: 0,
			},
			want: []testStr{
				{
					name:     "1",
					priority: 2,
				},
				{
					name:     "1",
					priority: 2,
				},
			},
		},
		{
			name: "#3",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 2,
					},
				},
				less: func(a, b testStr) bool {
					return a.priority < b.priority
				},
			},
			args: args{
				idx: 3,
			},
			want: []testStr{
				{
					name:     "1",
					priority: 1,
				},
				{
					name:     "1",
					priority: 2,
				},
				{
					name:     "1",
					priority: 2,
				},
			},
		},
		{
			name: "#3",
			h: Heap[testStr]{
				heap: []testStr{
					{
						name:     "1",
						priority: 1,
					},
					{
						name:     "1",
						priority: 2,
					},
					{
						name:     "1",
						priority: 2,
					},
				},
				less: func(a, b testStr) bool {
					return a.priority < b.priority
				},
			},
			args: args{
				idx: -1,
			},
			want: []testStr{
				{
					name:     "1",
					priority: 1,
				},
				{
					name:     "1",
					priority: 2,
				},
				{
					name:     "1",
					priority: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.DeleteElement(tt.args.idx)
		})
	}
}
