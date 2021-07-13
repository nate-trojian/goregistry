package goregistry_test

import (
	"testing"

	"github.com/nate-trojian/goregistry"
)

func BenchmarkReflect(b *testing.B) {
	r := goregistry.New(keyFunc)
	r.Register("intern", &intern{})
	r.Register("salary", &salary{})
	b.Run("Intern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := r.FromJSON(internInput); err != nil {
				b.Fatalf("Failed to convert intern - %v", err)
			}
		}
	})
	b.Run("Salary", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := r.FromJSON(salaryInput); err != nil {
				b.Fatalf("Failed to convert salary - %v", err)
			}
		}
	})
}
