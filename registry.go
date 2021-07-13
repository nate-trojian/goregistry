package goregistry

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	// ErrKeyNotFound is returned in FromJSON when keyFunc returns a key not in the Registry
	ErrKeyNotFound = errors.New("key not found in registry")
)

// New creates a new Registry type using a KeyFunction
func New(keyFunc KeyFunction) (Registry, error) {
	return &registry{
		keyFunc: keyFunc,
		mapping: make(map[string]reflect.Type),
	}, nil
}

// KeyFunction provides the definition for the function to return a key from a JSON string as a byte array
type KeyFunction func([]byte) (string, error)

// Registry interface defines the base Registry functionality
type Registry interface {
	Register(string, interface{})
	FromJSON([]byte) (interface{}, error)
}

// registry is the basic implementation of the Registry interface
type registry struct {
	keyFunc KeyFunction
	mapping map[string]reflect.Type
}

// Register adds a new type to the mapping
func (r *registry) Register(key string, v interface{}) {
	t := reflect.TypeOf(v)
	r.mapping[key] = t
}

// FromJSON converts a JSON string as a byte array into one of the types in the registry.
// It first passes the input JSON to the key function to determine which type to try to Unmarshal to
func (r *registry) FromJSON(data []byte) (interface{}, error) {
	key, err := r.keyFunc(data)
	if err != nil {
		return nil, fmt.Errorf("key function encountered an error - %w", err)
	}
	t, ok := r.mapping[key]
	if !ok {
		// Keep consistency with returning Wrapped Error
		return nil, fmt.Errorf("key %s not found - %w", key, ErrKeyNotFound)
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	ret := reflect.New(t).Interface()
	if err := json.Unmarshal(data, ret); err != nil {
		return nil, fmt.Errorf("failed to parse data into type with key %s - %w", key, err)
	}
	return ret, nil
}
