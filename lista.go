package main

import "fmt"

type No struct {
	data     int
	previous *No
	next     *No
	position uint
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
			"Current position: %p - data: %d - previous: %p - next %p\n",
			no, no.data, no.previous, no.next)
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

func main() {
	var list List
	list.addData(5)
	list.addData(10)
	list.addData(50)
	list.addData(10)
	list.addData(4)
	fmt.Println(list.findByPosition(list.findByData(10)...))
}
