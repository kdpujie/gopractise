package main

import (
	"fmt"
)

func main() {
	welcome := make([]string, 0, 0)
	fmt.Printf("[%p], size=%d, cap=%d \n", &welcome, len(welcome), cap(welcome))
	welcome = append(welcome, "hello")
	welcome = append(welcome, "world")
	changeSlice(welcome)
	fmt.Println(welcome)
	fmt.Printf("[%p], size=%d, cap=%d \n", &welcome, len(welcome), cap(welcome))
	deleteSlice("B")
}

func changeSlice(s []string) {
	s[0] = "go"
	fmt.Printf("[%p], size=%d, cap=%d \n", &s, len(s), cap(s))
	s = append(s, "playgroud")
	fmt.Printf("[%p], size=%d, cap=%d \n", &s, len(s), cap(s))
	fmt.Println(s)
	fmt.Printf("size=%d, cap=%d \n", len(s), cap(s))
}

func deleteSlice(str string) {
	s := []string{"A", "B", "C", "D", "B"}
	for i := 0; i < len(s); i++ {
		for i := range s {
			if s[i] == str {
				s[i] = s[len(s)-1]
				s = s[:len(s)-1]
				break
			}
		}
		fmt.Printf("s[%d]: %v \n", i, s)
	}
}
