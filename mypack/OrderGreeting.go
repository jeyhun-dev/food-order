package mypack

import (
	"strconv"
	"strings"
)

func Greeting(name string) string {
	return "Hello " + name
}

func SayHello(count int) string {
	if count == 0 {
		return "my hello count: " + strconv.Itoa(count)
	} else if count < 0 {
		return "my hello count does not negative number: " + strconv.Itoa(count)
	}
	myArray := []string{}
	for i := 0; i < count; i++ {
		myArray = append(myArray, "my hello count: "+strconv.Itoa(i))
	}
	return strings.Join(myArray, ", ")
}
