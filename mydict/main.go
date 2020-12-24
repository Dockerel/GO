package main

import (
	"fmt"

	"github.com/Dockerel/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete Ended")
	}
}
