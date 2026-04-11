package main

import "fmt"

func Hello(name string) string {
	return "Merhaba, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
