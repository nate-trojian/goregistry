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
	r, _ := goregistry.New(keyFunc)
	r.Register("bar", &bar{})
}

func ExampleRegistry_FromJSON() {
	type base struct {
		Type string
	}
	type bar struct {
		Name string
	}
	keyFunc := func(data []byte) (string, error) {
		b := &base{}
		err := json.Unmarshal(data, b)
		if err != nil {
			return "", err
		}
		return b.Type, nil
	}
	r, _ := goregistry.New(keyFunc)
	r.Register("bar", &bar{})

	data := []byte(
		`{
			"type": "bar",
			"name": "BAR"
		}`)
	b, err := r.FromJSON(data)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%v", b)
}
