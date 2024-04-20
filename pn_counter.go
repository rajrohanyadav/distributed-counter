package main

import "fmt"

type PNCounter struct {
	id   int
	incr []int
	decr []int
}

func NewPNCounter(id, numOfClients, initialValue int) *PNCounter {
	incr := make([]int, numOfClients)
	decr := make([]int, numOfClients)
	incr[id-1] = initialValue
	decr[id-1] = initialValue
	return &PNCounter{
		id:   id,
		incr: incr,
		decr: decr,
	}
}

func (pn *PNCounter) Query() int {
	return Sum(pn.incr) - Sum(pn.decr)
}

func (pn *PNCounter) Update(num int) {
	if num < 0 {
		pn.decr[pn.id-1] -= num
	} else {
		pn.incr[pn.id-1] += num
	}
}

func (pn *PNCounter) Merge(rc *PNCounter) {
	zippedIncr, err := Zip(pn.incr, rc.incr)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	zippedDecr, err := Zip(pn.decr, rc.decr)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	for i, e := range zippedIncr {
		pn.incr[i] = Max(e.a, e.b)
	}

	for i, e := range zippedDecr {
		pn.decr[i] = Max(e.a, e.b)
	}
}
