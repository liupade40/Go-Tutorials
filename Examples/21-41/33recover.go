package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		} else {
			fmt.Println("No panic happened.")
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
