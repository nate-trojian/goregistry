/*
The goregistry package is an implementation of the
Registry (see here: https://martinfowler.com/eaaCatalog/registry.html) design pattern.
It provides a convenient and centralized data structure for converting a general
JSON message into one of many different output structs.

Registry Interface

The Registery Constructor takes in a KeyFunction.
In most cases, a KeyFunction will decode the JSON string to a base intermediate struct that contains only the key field,
such as

  type base struct {
	  Type string
  }

  var keyFunc = func(msg []byte) (string, error) {
	  b := &base{}
	  err := json.Unmarshal(msg, b)
	  if err != nil {
		  return "", err
	  }
	  return b.Type, nil
  }

When a message is processed using FromJSON, it is first passed through to the KeyFunction.
The output is then used to get the type to try and Unmarshal to.

To register a new type, you need a string key and an instance of the type to Unmarshal to.
This was done to contain all reflect logic within the package, away from the user.

Usage

When you have multiple output structs across many files, it can be useful to create the registry in one central file
and add each struct to the registry in an init function in each struct's file.

  // registry.go
  var (
	  keyFunc = ...
	  registry = goregistry.New(keyFunc)
  )

  // foo.go
  func init() {
	  registry.Register("foo", &Foo{})
  }
  type Foo struct{}
*/
package goregistry
