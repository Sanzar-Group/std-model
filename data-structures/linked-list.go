package dataStructs

type ListNode[T any] struct {
    value T
    next *ListNode[T]
}

func (l *ListNode[T]) add(value T) {
    if l.next == nil {
        l.next = &ListNode[T]{value: value}
        return
    }
    l.next.add(value)
}
