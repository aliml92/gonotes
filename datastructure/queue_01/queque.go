package queue

import "errors"


const minQueueLen = 16


var ErrQueueEmpty = errors.New("queue is empty")
var ErrIndexOutOfRange = errors.New("index is out of range")


type Queue struct {
	buf               []interface{}
	head, tail, count int
}



func New() *Queue {
	return &Queue{
		buf: make([]interface{}, minQueueLen),
	}
}

func (q *Queue) Length() int {
	return q.count
}

func (q *Queue) resize() {
	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

func (q *Queue) Push(elem interface{}) {
	if q.count == len(q.buf) {
		q.resize()
	}

	q.buf[q.tail] = elem

	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++
}

func (q *Queue) Peek() (interface{}, error) {
	if q.count <= 0 {
		return nil, ErrQueueEmpty
	}
	return q.buf[q.head], nil
}

func (q *Queue) Get(i int) (interface{}, error) {
	// If indexing backwards, convert to positive index.
	if i < 0 {
		i += q.count
	}
	if i < 0 || i >= q.count {
		return nil, ErrIndexOutOfRange
		panic("queue: Get() called with index out of range")
	}
	// bitwise modulus
	return q.buf[(q.head+i)&(len(q.buf)-1)], nil
}

func (q *Queue) Pop() (interface{}, error) {
	if q.count <= 0 {
		return nil, ErrQueueEmpty
	}
	ret := q.buf[q.head]
	q.buf[q.head] = nil
	// bitwise modulus
	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--
	// Resize down if buffer 1/4 full.
	if len(q.buf) > minQueueLen && (q.count<<2) == len(q.buf) {
		q.resize()
	}
	return ret, nil
}