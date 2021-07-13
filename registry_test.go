package goregistry_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/nate-trojian/goregistry"
)

var (
	keyFunc = func(data []byte) (string, error) {
		b := &base{}
		err := json.Unmarshal(data, b)
		if err != nil {
			return "", err
		}
		return b.Type, nil
	}
	internInput = []byte(
		`{
		"type": "intern",
		"name": "Intern 1",
		"hourly_wage": 20.0,
		"hours_worked": 45
	}`)
	salaryInput = []byte(
		`{
		"type": "salary",
		"name": "Alice",
		"salary": 20.0,
		"hours_worked": 45
	}`)
)

type base struct {
	Type string
}

type intern struct {
	Name        string
	HourlyRate  float32
	HoursWorked float32
}

type salary struct {
	Name   string
	Salary float32
}

func TestNewRegistry(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		_, err := goregistry.New(keyFunc)
		if err != nil {
			t.Fatalf("Unexpected error %v", err)
		}
	})
}

func TestRegister(t *testing.T) {
	r, _ := goregistry.New(keyFunc)
	t.Run("Success", func(t *testing.T) {
		r.Register("intern", &intern{})
		r.Register("salary", &salary{})
	})
}

func TestFromJSON(t *testing.T) {
	// Setup
	r, _ := goregistry.New(keyFunc)
	r.Register("intern", &intern{})
	r.Register("salary", &salary{})
	// Error Tests
	tests := []struct {
		Name        string
		Input       []byte
		ExpectedErr error
	}{
		{"Invalid JSON", []byte(`{`), errors.New("unexpected end of JSON input")},
		{"No Value", []byte(`{}`), goregistry.ErrKeyNotFound},
		{"Unknown Key", []byte(`{"type":"a"}`), goregistry.ErrKeyNotFound},
		{"Fail to Unmarshal to Type", []byte(`{"type":"intern", "name": 1}`), errors.New("json: cannot unmarshal number into Go struct field intern.Name of type string")},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := r.FromJSON(test.Input)
			if test.ExpectedErr == nil && err == nil {
				return
			}
			if test.ExpectedErr == nil && err != nil {
				t.Fatalf("Unexpected error received %v", err)
				return
			}
			// Function returns wrapped error
			unwrapped := errors.Unwrap(err)
			if unwrapped == nil || unwrapped.Error() != test.ExpectedErr.Error() {
				t.Fatalf("Expecting error %v but received %v", test.ExpectedErr, err)
			}
		})
	}
	// Successful tests
	t.Run("Successful", func(t *testing.T) {
		t.Run("Intern", func(t *testing.T) {
			i, err := r.FromJSON(internInput)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if _, ok := i.(*intern); !ok {
				t.Fatalf("Failed to cast %v to Intern", i)
			}
		})
		t.Run("Salary", func(t *testing.T) {
			s, err := r.FromJSON(salaryInput)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if _, ok := s.(*salary); !ok {
				t.Fatalf("Failed to cast %v to Salary", s)
			}
		})
	})
}
