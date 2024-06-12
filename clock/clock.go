package main

import (
	"errors"
	"fmt"
	"time"
)

type Counter struct {
	Val   int
	Limit int
}

func (c *Counter) Reached() bool {
	return c.Val == c.Limit
}

func (c *Counter) Increase() {
	c.Val++
}

func (c *Counter) Reset() {
	c.Val = 0
}

// ---
type Clock struct {
	Values []Counter
}

func (c *Clock) Increment() (err error) {
	// position identifier
	length := len(c.Values)
	var ix int
	for i := length - 1; i > -1; i-- {
		if !c.Values[i].Reached() {
			ix = i
			break
		}
		// end case
		if i == 0 {
			err = errors.New("clock reach end of increments")
			return
		}
	}

	// increase
	c.Values[ix].Increase()
	for i := ix + 1; i < length; i++ {
		c.Values[i].Reset()
	}
	return
}

func main() {
	clock := Clock{
		Values: []Counter{
			{Val: 0, Limit: 3},
			{Val: 0, Limit: 4},
			{Val: 0, Limit: 5},
		},
	}

	for {
		// display
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("clock: %d-%d-%d\n", clock.Values[0].Val, clock.Values[1].Val, clock.Values[2].Val)

		// increment
		err := clock.Increment()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
