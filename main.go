package main

import "file_generator/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
