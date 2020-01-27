package main

import "fmt"

func main() {
	for {
		defer func() {
			fmt.Println("hello")
			for {
			}
		}()
		panic("yolo")
	}
}
