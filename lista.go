package main

import (
	"fmt"
)

type No struct {
	data     int
	previous *No
	next     *No
	position uint
	inactive bool
}

type Deleted struct {
	position uint
	deleted  bool
}

type List struct {
	start *No
	end   *No
}

func (l List) setLastPosition() uint {
	position := l.end.position + 1
	return position
}

func (l *List) addData(data int) {
	var no No
	no.data = data
	if l.start == nil {
		l.start = &no
		l.end = &no
	} else {
		no.position = uint(l.setLastPosition())
		no.previous = l.end
		l.end.next = &no
		l.end = &no
	}
}

func (l List) showList() {
	var no *No
	no = l.start
	for no != nil {
		fmt.Printf(
			"Current position: %d - data: %d - previous: %p - next %p - inactive: %t\n",
			no.position, no.data, no.previous, no.next, no.inactive)
		no = no.next
	}
}

func (l List) findByData(data int) []uint {
	var result []uint
	var no *No
	no = l.start
	for no != nil {
		if no.data == data {
			result = append(result, no.position)
		}
		no = no.next
	}
	return result
}
func (l List) findByPosition(positions ...uint) []int {
	var datas []int
	var no *No
	no = l.start
	for no != nil {
		for _, p := range positions {
			if no.position == p {
				datas = append(datas, no.data)
			}
		}
		no = no.next
	}
	return datas
}

func (l *List) logicalDeletion(positions ...uint) []Deleted {
	var deleteds []Deleted
	var no *No
	no = l.start
	for no != nil {
		for _, p := range positions {
			if no.position == p {
				no.inactive = true
				var deleted Deleted
				deleted.position = p
				deleted.deleted = true
				deleteds = append(deleteds, deleted)
			}
		}
		no = no.next
	}
	return deleteds
}

func (l *List) deletionDefined() {
	var no, deleted *No
	no = l.start
	for no != nil {
		if no.inactive {
			if deleted == nil {
				deleted = no.next
			}
			if no == l.start {
				l.start = no.next
				l.start.previous = nil
			} else if no.next == nil {
				l.end = no.previous
				l.end.next = nil
			} else {
				no.previous.next = no.next
				no.next.previous = no.previous
			}
		}
		no = no.next
	}
	no = deleted
	for no != nil {
		no.position = no.previous.position + 1
		no = no.next
	}
}

func main() {
	var list List
	list.addData(5)
	list.addData(10)
	list.addData(50)
	list.addData(10)
	list.addData(4)
	list.addData(10)
	positions := list.findByData(10)
	list.logicalDeletion(positions...)
	list.showList()
	list.deletionDefined()
	list.showList()

}
