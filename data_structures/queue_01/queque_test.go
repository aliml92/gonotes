package queue

import "testing"



func TestQuequeSimple(t *testing.T) {
	q := New()


	for i := 0; i < minQueueLen; i++ {
		q.Push(i)
	}

	for i := 0; i < minQueueLen; i++ {
		if v, _ := q.Peek(); v.(int) != i {
			t.Error("peek", i, "had value", v)
		}
		x, _ := q.Pop()
		if x != i {
			t.Error("remove", i, "had value", x)
		}
	}
}



func BenchmarkQueueGet(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_,_  = q.Get(i)
	}
}

func BenchmarkQueueTickTock(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Push(nil)
		q.Peek()
		q.Pop()
	}
}