package main

import "fmt"

func ShowHello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(ShowHello("reza"))
}
