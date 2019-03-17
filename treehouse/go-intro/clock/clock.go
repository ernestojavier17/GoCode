package clock

import "fmt"

type Minutes struct {
	value int
}

func (m *Minutes) increment() {
	m.value = (m.value + 1) %60 //We use another modulo operator here to ensure that if we set a number of minutes
	//greater than 60, it will wrap around a start at zero.
	//Notice, by the way we did not have to get the value from the pointers using code (*m).value = (*m).value + 1) % 60
	//When you access a struct field through a pointer, the pointer is automatically converted to a value.
}

func (m *Minutes) set(newValue int) {
	m.value = newValue % 60
}

func (m Minutes) display() {
	fmt.Println(m.value)
}

type Clock struct {
	Hours int
	Minutes int
}

func Noon() Clock {
	return Clock{12, 0}
}

