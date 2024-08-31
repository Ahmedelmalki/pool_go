package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(fromto(1, 10))
	fmt.Print(fromto(10, 1))
}

func fromto(from, to int) string {
	rs := ""
	if from < to {
		for i := from; i < to; i++ {
			if i < 10 {
				rs += "0" + strconv.Itoa(i) + ", "
			} else {
				rs += strconv.Itoa(i) + ", "
			}
		}
		if to < 10 {
			rs += "0" + strconv.Itoa(to) + "\n"
		} else {
			rs += strconv.Itoa(to) + "\n"
		}
	} else if to < from{
		for i := from; i > to; i-- {
			if i < 10 {
				rs += "0" + strconv.Itoa(i) + ", "
			} else {
				rs += strconv.Itoa(i) + ", "
			}
		}
		if to < 10 {
			rs += "0" + strconv.Itoa(to) + "\n"
		} else {
			rs += strconv.Itoa(to) + "\n"
		}
	} else {
		rs = strconv.Itoa(to) + "\n"
	}
	if from > 99 || to > 99 || to < 0 || from < 0 {
		rs = "Invalid\n"
	}
	return rs
}