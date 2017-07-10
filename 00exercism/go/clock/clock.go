package clock

import "fmt"

const testVersion = 4

// Clock : Complete the type definition.  Pick a suitable data type.
type Clock struct {
	hour   int
	minute int
}

// New creates a new clock
func New(hour, minute int) Clock {
	c := Clock{0, 0}
	if hour >= 0 && minute >= 0 {
		min := minute + (hour * 60)
		c = c.Add(min)
	} else if minute >= 0 {
		h := 24 - (24 % (hour * -1))
		fmt.Println("****------- HOUR -----****", h)
		min := minute + (h * 60)

		c = c.Add(min)
	}
	return c
}

// Returns the time of a clock as a string
func (c Clock) String() string {
	h := fmt.Sprintf("%v", c.hour)
	min := fmt.Sprintf("%v", c.minute)
	if len(h) == 1 {
		h = "0" + h
	}
	if len(min) == 1 {
		min = "0" + min
	}
	// fmt.Printf("TEST :::::: %v:%v\n", h, min) // TEST
	return fmt.Sprintf("%v:%v", h, min)
}

// Add some minutes to the clock
func (c Clock) Add(minutes int) Clock {
	if c.minute >= 0 {
		min := c.minute + (c.hour * 60) + minutes
		h := min / 60
		fmt.Println("******** ADD *** ", min)
		if h > 24 {
			h %= 24
		}
		if h == 24 {
			h = 0
		}
		min %= 60
		fmt.Println("HOUR \t", h, "\t MINUTES \t", min)
		return Clock{h, min}
	}
	return Clock{c.hour, c.minute}
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
