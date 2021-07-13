package goregistry_test

import (
	"encoding/json"
	"fmt"

	"github.com/nate-trojian/goregistry"
)

func ExampleNew() {
	type base struct {
		Type string
	}
	type bar struct{}
	keyFunc := func(data []byte) (string, error) {
		b := &base{}
		err := json.Unmarshal(data, b)
		if err != nil {
			return "", err
		}
		return b.Type, nil
	}
	r := goregistry.New(keyFunc)
	r.Register("bar", &bar{})
}

func Example() {
	// Structs
	type base struct {
		Type string
	}
	type foo struct {
		Name string
	}
	// Key function
	keyFunc := func(data []byte) (string, error) {
		b := &base{}
		err := json.Unmarshal(data, b)
		if err != nil {
			return "", err
		}
		return b.Type, nil
	}
	// Registry Initialization
	r := goregistry.New(keyFunc)
	r.Register("foo", &foo{})

	// Message to decode
	data := []byte(
		`{
			"type": "foo",
			"name": "bar"
		}`)
	b, err := r.FromJSON(data)
	// Did it work?
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	// Of course it did
	fmt.Printf("%+v", b)
}
