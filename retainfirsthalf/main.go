package main

import "fmt"

func main() {
	fmt.Println(retainfirsthalf("This is the 1st halfThis is the 2nd half"))
}

func retainfirsthalf(str string) string {
	str = str[:len(str)/2]
	return str
}
