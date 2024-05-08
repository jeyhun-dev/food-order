package controller

import (
	"fmt"
	"net/http"
)

var num int

func HomePage(w http.ResponseWriter, r *http.Request) {
	num++
	fmt.Fprintf(w, "Welcome to the Azerbaijan!\n")
	fmt.Fprintf(w, "Welcome to the Baku!")

	fmt.Printf("Endpoint Hit: homePage %v\n", num)
}
