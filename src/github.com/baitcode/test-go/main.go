package main

import "fmt"

type ExportedStruct struct {
	shit string
}

func (e *ExportedStruct) SayShit() {
	fmt.Println(e.shit)
}

func main() {
	s := ExportedStruct{
		shit: "shit",
	}
	s.SayShit()
}
