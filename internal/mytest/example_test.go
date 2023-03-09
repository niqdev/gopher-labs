package mytest

import (
	"testing"

	example "github.com/niqdev/gopher-labs/pkg"
	"github.com/stretchr/testify/assert"
)

func TestHelloEmpty(t *testing.T) {
	_, err := example.Hello("")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestHello(t *testing.T) {
	result, err := example.Hello("test")

	if err != nil {
		t.Fatal("expected success")
	}
	if result != "hello test" {
		t.Fatal("unexpected value")
	}
}

// https://github.com/stretchr/testify
func TestAssertions(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
}

// TODO https://onsi.github.io/ginkgo
