package main

import (
	"fmt"
	"food-order/src/mypack"
	"strings"
)

func main() {
	var count int
	fmt.Printf("Hello World\n")
	var greetingResponse = mypack.Greeting("Jeyhun Mammadsaidov")
	fmt.Println(greetingResponse + " \n" + strings.Repeat("-", 10))
	fmt.Printf("enter your count: \n")
	fmt.Scan(&count)
	var myHelloCount = mypack.SayHello(count)
	fmt.Printf("result: %s\n", myHelloCount)
}
