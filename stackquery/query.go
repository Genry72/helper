package stackquery

/*
Query
Очередь FIFO (First In First Out) - первым пришел, первым ушел.
*/
type Query[V any] struct {
	query []V
}

func NewQuery[V any](count int) *Query[V] {
	return &Query[V]{
		query: make([]V, 0, count),
	}
}

func NewQueryFromSlice[V any](sl []V) *Query[V] {
	return &Query[V]{
		query: sl,
	}
}

func (q *Query[V]) Len() int {
	return len(q.query)
}

func (q *Query[V]) Push(v V) {
	q.query = append(q.query, v)
}

func (q *Query[V]) Pop() V {
	if q.Len() == 0 {
		panic("empty query")
	}
	v := q.query[0]

	q.query = q.query[1:]

	return v
}
