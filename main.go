package main

import "fmt"

const (
	NUM_OF_CLIENTS = 2
)

func main() {
	fmt.Println("==============================")
	fmt.Println("G-Counter - Grow only counter")
	fmt.Println("==============================")
	cl1 := NewGCounter(1, NUM_OF_CLIENTS, 0)
	cl2 := NewGCounter(2, NUM_OF_CLIENTS, 0)

	cl1.Increment(1)
	cl2.Increment(3)
	cl2.Increment(3)
	cl1.Increment(1)
	cl1.Increment(1)
	fmt.Println("Before Merge")
	fmt.Printf("client 1: %v\n", cl1.Query())
	fmt.Printf("client 2: %v\n", cl2.Query())

	cl1.Merge(cl2)
	cl2.Merge(cl1)
	fmt.Println("After Merge")
	fmt.Printf("client 1: %v\n", cl1.Query())
	fmt.Printf("client 2: %v\n", cl2.Query())
	
	fmt.Println("========================================")
	fmt.Println("PN-Counter - Positive Negative Counter")	
	fmt.Println("========================================")
	cl3 := NewPNCounter(1, NUM_OF_CLIENTS, 0)
	cl4 := NewPNCounter(2, NUM_OF_CLIENTS, 0)

	cl3.Update(2)
	cl3.Update(3)
	cl4.Update(5)
	cl3.Update(-4)
	cl4.Update(-2)

	fmt.Println("Before Merge")
	fmt.Printf("client 3: %v\n", cl3.Query())
	fmt.Printf("client 4: %v\n", cl4.Query())

	cl3.Merge(cl4)
	cl4.Merge(cl3)
	fmt.Println("After Merge")
	fmt.Printf("client 3: %v\n", cl3.Query())
	fmt.Printf("client 4: %v\n", cl4.Query())
}
