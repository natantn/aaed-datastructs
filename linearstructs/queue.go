package linearstructs

type QueueMember struct {
	Value int
	Next *QueueMember
}

func CreateMember(value int) *QueueMember{
	return &QueueMember{Value: value}
}

type Queue struct {
	First, Last *QueueMember
}

func InitQueue() *Queue {
	return &Queue{}
}

func (q *Queue) EnterQueue(value int) {
	m := CreateMember(value)
	if q.First == nil {
		q.First = m
	} else {
		last := q.Last
		last.Next = m
	}	
	q.Last = m
	return
}

func (q *Queue) GetFromQueue() (m *QueueMember){
	if q.First == nil {
		return nil
	}

	m = q.First
	q.First = m.Next
	if q.Last == m {
		q.Last = nil
	}
	return
}