package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	seconds int64
}

func New(h, m int) Clock {
	clock := Clock{}
	clock = clock.AddH(h)
	clock = clock.Add(m)
	return clock
}

func (c Clock) AddH(h int) Clock {
	if h < 0 {
		h = 24 + (h % 24)
	}
	if h > 24 {
		h = h % 24
	}
	c.seconds += int64(h * 3600)
	return c
}

func (c Clock) Add(m int) Clock {

	if m < 0 {
		return c.Subtract(m * -1)
	}

	c.seconds += int64(m * 60)

	c.seconds = c.seconds % 86400

	return c

}

func (c Clock) Subtract(m int) Clock {

	if m < 0 {
		return c.Add(m * -1)
	}
	c.seconds -= int64(m * 60)
	if c.seconds < 0 {
		c.seconds = 86400 + (c.seconds % 86400)
	}
	return c
}

func (c Clock) String() string {
	fmt.Printf("Clock %02d:%02d\n", c.getHours(), c.getMinutes())
	return fmt.Sprintf("%02d:%02d", c.getHours(), c.getMinutes())
}

func (c Clock) getHours() int {
	hoursSeconds := int(c.seconds) / 3600
	hoursBy24 := hoursSeconds % 24
	return hoursBy24
}

func (c Clock) getMinutes() int {
	return (int(c.seconds) % 3600) / 60
}
