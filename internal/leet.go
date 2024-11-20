package internal

import (
	"fmt"
	"reflect"
)

type Leet struct{}

// todo make type for leetchars = https://en.wikipedia.org/wiki/Leet#Orthography
func leetIterator() {
	fields := reflect.VisibleFields(reflect.TypeOf(Chars{}))
	for _, field := range fields {
		fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
	}
}

func (Leet) Encode(str string) string {
	leetIterator()
	return "todo return str->leet cipher"
}

func (Leet) Decode(str string) string {
	return "todo return leet->str msg"
}
