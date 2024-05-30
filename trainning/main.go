package main

import "fmt"

type human struct {
	name string
	age  uint
	iq   int
}

func (a *human) getIq(iq int) {
	a.iq = iq
}

type User struct {
	Username string
	Email    string
	Age      uint
	Xp       float32
	Status   bool
}

func (u User) GetStatus() {
	fmt.Println("is the user active:", u.Status)
}

func main() {
	Allal := User{"allal", "elmalki7777@gmail.com", 21, 23.4, true}
	s := human{"3ab3ali", 22, -27}
	fmt.Println(s)
	s.getIq(20)
	fmt.Println(s)
	Allal.GetStatus()
}

// commit
