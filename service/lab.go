package main

import "fmt"

func main() {
	m := make(map[string]bool)
	m["jack"] = true
	m["john"] = false

	l := []string{"a", "b", "c"}
	for index, v := range l {
		fmt.Println(index, v)
	}

	for k,v := range m {
		fmt.Println(k, v)
	}
}
