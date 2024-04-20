package main

import (
	"errors"
	"fmt"
)

type GCounter struct {
	id    int
	state []int
}

func NewGCounter(id int, numOfClients int, initialValue int) *GCounter {
	st := []int{}
	for range numOfClients {
		st = append(st, 0)
	}
	st[id-1] = initialValue
	return &GCounter{
		id: id,
		state: st,
	}
}

func (c *GCounter) Query() int {
	return Sum(c.state)	
}

func (c *GCounter) Increment(num int) error {
	if num < 0 {
		return errors.New("value can't be negative")
	}
	c.state[c.id-1] += num
	return nil
}

func (c *GCounter) Merge(rc *GCounter) {
	zipped, err := Zip(c.state, rc.state)
	if err != nil {
		fmt.Println("Something failed")
	}
	for i, e := range zipped {
		c.state[i] = Max(e.a, e.b)
	}
}
