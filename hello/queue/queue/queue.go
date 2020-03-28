package queue

// A FIFO queue
//type Queue []int
type Queue []interface{}	//interface就能接收任何的参数

//Pushes the element into the queue
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

//Pops element from head
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//Returns if the queue is empty of not
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
