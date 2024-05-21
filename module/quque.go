package module

type RingQueue struct {
	queue    [][]byte
	head     int
	tail     int
	capacity int
}

func NewRingQueue(capacity int) *RingQueue {
	return &RingQueue{
		capacity: capacity,
		queue:    make([][]byte, capacity),
	}
}

func (rq *RingQueue) Enqueue(item []byte) {
	if rq.tail == rq.capacity {

		rq.queue[rq.head] = item
		rq.head = (rq.head + 1) % rq.capacity
		return
		// rq.head += 1
		// rq.head = rq.head % rq.capacity
	}
	rq.queue[rq.tail] = item
	rq.tail = rq.tail + 1
}

func (rq *RingQueue) Dequeue() []byte {
	if rq.head == rq.tail {
		// 队列空
		return nil
	}
	item := rq.queue[rq.head]
	rq.head = (rq.head + 1) % rq.capacity
	return item
}

func (rq *RingQueue) Each(f func([]byte)) {
	if rq.tail != rq.capacity {
		// rq.head = 0
		for i := 0; i < rq.tail; i++ {
			f(rq.queue[i])
		}
	} else {
		for i := rq.head; i < rq.capacity; i++ {
			f(rq.queue[i])
		}
		for i := 0; i < rq.head; i++ {
			f(rq.queue[i])
		}
	}
}
