package main

import "fmt"

const MAX_SIZE = 100

var stack [MAX_SIZE]*int
var top = -1

func push(data int) {
	if top == MAX_SIZE-1 {
		fmt.Println("Error: Stack overflow")
		return
	}
	ptr := new(int)
	*ptr = data
	top++
	stack[top] = ptr
}

func pop() {
	if top == -1 {
		fmt.Println("Error: Stack underflow")
		return
	}
	ptr := stack[top]
	top--
	*ptr = 0
}

func peek() int {
	if top == -1 {
		fmt.Println("Error: Stack is empty")
		return -1
	}
	ptr := stack[top]
	return *ptr
}

func main() {
	push(10)
	push(20)
	push(30)
	fmt.Printf("Top element: %d\n", peek())
	pop()
	fmt.Printf("Top element: %d\n", peek())
	pop()
	fmt.Printf("Top element: %d\n", peek())
	pop()
	fmt.Printf("Top element: %d\n", peek()) // should print "Error: Stack is empty"
}
