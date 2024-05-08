package main

import (
	"fmt"
	constructor2 "food-order/constructor"
	"food-order/controller"
	"log"
	"net/http"
	"strings"
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

	//testLesson := setterandgetter.Test{}
	//testLesson.SetQuiz("Data Type in Java")
	//testLesson.SetCount(15)
	//
	//fmt.Printf("My quiz name:  %v\n", testLesson.GetQuiz())
	//fmt.Printf("My quiz count: %v", testLesson.GetCount())

	//myStruct := defaultvalues.MyStruct{}
	//
	//fmt.Println("Integer:", myStruct.Integer)   // 0
	//fmt.Println("Floating:", myStruct.Floating) // 0
	//fmt.Println("String:", myStruct.Str)        // ""
	//fmt.Println("Bool:", myStruct.Boolean)      // false
	//fmt.Println("Pointer:", myStruct.Pointer)   // nil
	//fmt.Println("Slice:", myStruct.Slice)       // []
	//fmt.Println("Map:", myStruct.Map)           // map[]

	cat := constructor2.NewCat("Jessie", "Los Angeles", "lady", 1)
	fmt.Printf("My consructor full result: %v\n", cat)
	fmt.Printf(strings.Repeat("*\n", 3))

	horse := constructor2.NewHorse("Lime", 5)
	fmt.Printf("My horse name: %v\n", horse.Name)
	fmt.Printf("My horse age: %v\n", horse.Age)

	//msg := fmt.Printf("Hello, my name is: %v", horse.Name)
	msgTwo := fmt.Sprintf("Hello my name is %v", horse.Name)
	//fmt.Println(msg)
	fmt.Println(msgTwo)

	fmt.Printf(strings.Repeat("*\n", 3))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		controller.HomePage(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
