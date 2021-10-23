package main

import (
	"bytes"
	"fmt"
	"ruby/model"
)

func main() {
	var input model.InputBuffer
	for {
		print_prompt()
		read_input(&input)
		if input.Buffer.String() == model.Exit {
			break
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", input.Buffer.String())
		}
	}
}

func print_prompt() {
	fmt.Print("rubyDB > ")
}

func read_input(input *model.InputBuffer) {
	var str string
	fmt.Scanf("%s", &str)
	if len(str) == 0 {
		fmt.Print("Error reading input\n")
		return
	}
	input.Buffer = bytes.NewBufferString(str)
}
