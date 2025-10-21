package main

import "fmt"

func main() {
	var s string = "hello 世界ooo"
	for _, v := range s {
		if string(v) == "o" {
			fmt.Print(string(v))
		}
	}
}