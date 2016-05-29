package linkedlist

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type Linkedlist struct {
	head *node
}

func (l *Linkedlist) Add(value interface{}) {
	new_node := new(node)
    new_node.value = value

	i, j := l.head, l.head
	if i == nil {
        // empty list
		l.head = new_node
		return
	}

	for ; i != nil; i = i.next {
		j = i
	}
	j.next = new_node
	return
}

func (l *Linkedlist) Del() interface{} {
    i, j := l.head, l.head

    if i == nil {
        // Empty list
        return nil
    }

    for ; i != nil ; i = i.next {
        if i.next != nil {
            j = i
            continue
        }
        j.next = nil
        break
    }
    return i
}

func (l Linkedlist) Walk() {
	fmt.Println(l)
	for i := l.head; i != nil; i = i.next {
		fmt.Println(i)
	}
}
