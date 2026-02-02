package phonebook

import (
	"maps"
)

var book = make(map[string]string)

func Add(name, phone string) {
	book[name] = phone
}

func Search(name string) (string, bool) {
	value, ok := book[name]

	if !ok {
		return "", false
	}
	return value, true
}

func Delete(name string) {
	delete(book, name)
}

func List() map[string]string {
	return maps.Clone(book)
}
