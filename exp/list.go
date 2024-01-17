package main

import "fmt"

func iterate() {
	l := []int{1, 2, 3, 4, 5, 6}

	for k, v := range l {
		fmt.Println("K", k, " v ", v)
	}
}
