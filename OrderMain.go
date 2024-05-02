package main

import (
	"fmt"
	"food-order/src/defaultvalues"
)

func main() {
	//var count int
	//fmt.Printf("Hello World\n")
	//var greetingResponse = mypack.Greeting("Jeyhun Mammadsaidov")
	//fmt.Println(greetingResponse + " \n" + strings.Repeat("-", 10))
	//fmt.Printf("enter your count: \n")
	//fmt.Scan(&count)
	//var myHelloCount = mypack.SayHello(count)
	//fmt.Printf("result: %s\n", myHelloCount)

	/*	person := setterandgetter.Person{}

		person.SetName("Jeyhun")
		person.SetAge(25)
		person.SetAddress("Baku")
		time.Sleep(10000)

		fmt.Printf("My name is %v, \nand "+
			"I'm %v years old, \nand "+
			"my address is %v,\n ",
			person.GetName(),
			person.GetAge(),
			person.GetAddress())*/

	myStruct := defaultvalues.MyStruct{}

	fmt.Println("Integer:", myStruct.Integer)   // 0
	fmt.Println("Floating:", myStruct.Floating) // 0
	fmt.Println("String:", myStruct.Str)        // ""
	fmt.Println("Bool:", myStruct.Boolean)      // false
	fmt.Println("Pointer:", myStruct.Pointer)   // nil
	fmt.Println("Slice:", myStruct.Slice)       // []
	fmt.Println("Map:", myStruct.Map)           // map[]
}
