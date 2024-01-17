package queue

type Queue []int

// Pushes the element into the queue_array
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pops element from head
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Return if the queue_array is empty or not
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
