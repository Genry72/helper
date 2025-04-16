package stackquery_test

import (
	"github.com/Genry72/helper/stackquery"
	"testing"
)

type stackQueryer[V any] interface {
	Len() int
	Push(v V)
	Pop() V
}

func BenchmarkStackquery(b *testing.B) {
	const count = 100000

	b.Run("stackOneElement", func(b *testing.B) {
		stack := stackquery.NewStack[int](count)
		for n := 0; n < b.N; n++ {
			checkOneElement(stack, count)
		}
	})

	b.Run("queryOneElement", func(b *testing.B) {
		query := stackquery.NewQuery[int](0)
		for n := 0; n < b.N; n++ {
			checkOneElement(query, count)
		}
	})

	b.Run("stackAllElement", func(b *testing.B) {
		stack := stackquery.NewStack[int](count)
		for n := 0; n < b.N; n++ {
			checkAllElement(stack, count)
		}
	})

	b.Run("queryAllElement", func(b *testing.B) {
		query := stackquery.NewQuery[int](count)
		for n := 0; n < b.N; n++ {
			checkAllElement(query, count)
		}
	})
}

// Добавление и извлечение одного элемента
func checkOneElement[V int](sq stackQueryer[V], count int) {
	for i := 0; i < count; i++ {
		sq.Push(V(i))
		sq.Pop()
	}
}

// Добавление и извлечение всех элементов
func checkAllElement[V int](sq stackQueryer[V], count int) {
	for i := 0; i < count; i++ {
		sq.Push(V(i))
	}

	for sq.Len() > 0 {
		sq.Pop()
	}
}
