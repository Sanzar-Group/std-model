package dataStructs

type listNode[T any] struct {
    value T
    next *listNode[T]
}

func (l *listNode[T]) add(value T) {
    if l.next == nil {
        l.next = &listNode[T]{value: value}
        return
    }
    l.next.add(value)
}
