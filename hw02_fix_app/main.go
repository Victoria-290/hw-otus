package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
)

func main() {
	path := "data.json"

	fmt.Print("Enter data file path: ")
	_, err := fmt.Scanln(&path)
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err := reader.ReadJSON(path, -1)
	fmt.Printf("Error reading JSON: %v\n", err)

	printer.PrintStaff(staff)
}
