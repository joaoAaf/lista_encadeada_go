package main

import (
	"fmt"
)

type Person struct {
	name     string
	lastName string
}

type No struct {
	data     interface{}
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

func (l *List) addData(person interface{}) {
	var no No
	no.data = person

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
			"Current position: %d - person: %s - previous: %p - next %p - inactive: %t\n",
			no.position, no.data, no.previous, no.next, no.inactive)
		no = no.next
	}
}

func (l List) sliceList(position1 uint, position2 uint) []interface{} {
	var datas []interface{}
	no := l.findByPositionNo(position1)
	for no.position <= position2 {
		datas = append(datas, no.data)
		no = *no.next
	}
	return datas
}

func (l List) findByData(data interface{}) []uint {
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

func (l List) findByPositions(positions ...uint) []interface{} {
	var datas []interface{}
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

func (l List) findByPosition(position uint) interface{} {
	var data interface{}
	var no *No
	no = l.start
	for no != nil {
		if no.position == position {
			data = no.data
		}
		no = no.next
	}
	return data
}

func (l List) findByPositionNo(position uint) No {
	var no, result *No
	no = l.start
	for no != nil {
		if no.position == position {
			result = no
		}
		no = no.next
	}
	return *result
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

func (l *List) deletionDefinitive() {
	var no, deleted *No
	no = l.start
	for no != nil {
		if no.inactive {
			deleted = no
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
			for deleted != nil {
				deleted.position = deleted.position - 1
				deleted = deleted.next
			}
			deleted = nil
		}
		no = no.next
	}
}

/*func (l *List) deletionDefinitive() {
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
}*/

func (l *List) fastDeletion(positions ...uint) []Deleted {
	deleteds := l.logicalDeletion(positions...)
	l.deletionDefinitive()
	return deleteds
}

func (l *List) update(data interface{}, position uint) {
	var no *No
	no = l.start
	for no != nil {
		if no.position == position {
			no.data = data
			no = nil
		} else {
			no = no.next
		}
	}

}

func main() {
	var list List
	list.addData(Person{"João Anderson", "de Assis Freitas"})
	list.addData(Person{"João Anderson", "de Assis Freitas"})
	list.addData(Person{"João Anderson", "de Assis Freitas"})
	list.addData(Person{"Marcelle", "Tabosa de Souza"})
	list.addData(Person{"João Anderson", "de Assis Freitas"})
	list.addData(Person{"João Anderson", "de Assis Freitas"})
	positions := list.findByData(Person{"João Anderson", "de Assis Freitas"})
	list.logicalDeletion(positions...)
	list.deletionDefinitive()
	list.showList()
	list.update(Person{"João Anderson", "de Assis Freitas"}, 0)
	list.showList()
}
