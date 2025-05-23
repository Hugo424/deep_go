package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

type Animal struct {
	Name    string `properties:"name"`
	Age     int    `properties:"age"`
	Species string `properties:"species,omitempty"`
}

func Serialize[T any](object T) string {
	v := reflect.ValueOf(object)
	t := reflect.TypeOf(object)

	numFields := t.NumField()
	builder := &strings.Builder{}
	for i := 0; i < numFields; i++ {
		tag := t.Field(i).Tag
		field := v.Field(i)

		tagValue, tagPresent := tag.Lookup("properties")
		isOmitempty := strings.Contains(tagValue, "omitempty")
		tagValue = strings.Replace(tagValue, ",omitempty", "", 1)

		if (!tagPresent) || (field.IsZero() && isOmitempty) {
			continue
		}

		builder.WriteString(fmt.Sprintf("%s=%v\n", tagValue, field))
	}

	return strings.TrimSpace(builder.String())
}

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		item   any
		result string
	}{
		"test case with empty fields struct 1": {
			item:   Person{},
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			item: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field struct 1": {
			item: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
		"test case with empty fields struct 2": {
			item:   Animal{},
			result: "name=\nage=0",
		},
		"test case with fields struct 2": {
			item: Animal{
				Name: "Barsique",
				Age:  1,
			},
			result: "name=Barsique\nage=1",
		},
		"test case with omitempty field struct 2": {
			item: Animal{
				Name:    "Sharique",
				Age:     2,
				Species: "dog",
			},
			result: "name=Sharique\nage=2\nspecies=dog",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.item)
			assert.Equal(t, test.result, result)
		})
	}

}
