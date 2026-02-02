package main

import (
	"fmt"
	phonebook "go_basic/cmd/06_maps"
)

func main() {
	phonebook.Add("Jay", "123-456-7890")
	phonebook.Add("John", "987-654-3210")
	fmt.Println(phonebook.List())
	fmt.Println(phonebook.Search("Jay"))
	phonebook.Delete("Jay")
}
