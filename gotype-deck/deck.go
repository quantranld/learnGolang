package main

import "fmt"

type deck []string

func (d deck) handsome() {
	for index, value := range d {
		fmt.Println(index, value)
	}
}
