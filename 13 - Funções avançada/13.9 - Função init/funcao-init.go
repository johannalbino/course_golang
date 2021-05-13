package main

import "fmt"

var n int

func init() {
	fmt.Println("Eu executo antes de todo mundo.")
	n = 10
}

func main() {
	fmt.Println("Eu sou o main.")
	fmt.Println(n)
}
