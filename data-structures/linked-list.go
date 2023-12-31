package dataStructs

type ListNode[T any] struct {
    Value T
    Next *ListNode[T]
}

func (l *ListNode[T]) Add(value T) {
    if l.Next == nil {
        l.Next = &ListNode[T]{Value: value}
        return
    }
    l.Next.Add(value)
}
