package list

/*
List // тип контейнер
Len() // длинна списка
First() // первый Item
Last() // последний Item
PushFront(v interface{}) // добавить значение в начало
PushBack(v interface{}) // добавить значение в конец
Remove(i Item) // удалить элемент ​
Item // элемент списка
Value() interface{} // возвращает значение
Nex() *Item // следующий Item
Prev() *Item // предыдущий
*/

// Item is an element for List
type Item struct {
	previous *Item
	value    interface{}
	next     *Item
}

// Value returns value of i
func (i *Item) Value() interface{} {
	return i.value
}

// Next returns next element for i
func (i *Item) Next() *Item {
	return i.next
}

// Prev returns previous element for i
func (i *Item) Prev() *Item {
	return i.previous
}

// List is a container type
type List struct {
	first *Item
	last  *Item
}

// NewList is a constructor for an empty List
func NewList() List {
	newList := List{
		first: &Item{},
		last:  &Item{},
	}
	newList.first.next = newList.last
	newList.last.previous = newList.first
	return newList
}

// First returns first element if l
func (l List) First() *Item {
	if l.Len() == 0 {
		return nil
	}
	return l.first.next
}

// Last returns first element if l
func (l List) Last() *Item {
	if l.Len() == 0 {
		return nil
	}
	return l.last.previous
}

// Len returns length of l
func (l List) Len() int {
	len := 0
	current := l.first.next

	for current != l.last {
		len++
		current = current.next
	}

	return len
}

// PushFront adds an item as a first element for l
func (l List) PushFront(v interface{}) {
	newItem := &Item{
		previous: l.first,
		value:    v,
		next:     l.first.next,
	}
	l.first.next.previous = newItem
	l.first.next = newItem
}

// PushBack adds an item as a last element for l
func (l List) PushBack(v interface{}) {
	newItem := &Item{
		previous: l.last.previous,
		value:    v,
		next:     l.last,
	}
	l.last.previous.next = newItem
	l.last.previous = newItem
}

// Remove removes i from l, if not found do nothing
func (l List) Remove(i *Item) {
	current := l.first.next
	for {
		if current == l.last {
			return
		}
		if current == i {
			current.previous.next = current.next
			current.next.previous = current.previous
			return
		}
		current = current.next
	}
}
