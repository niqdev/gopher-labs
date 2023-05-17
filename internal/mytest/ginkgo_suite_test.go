package mytest_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// https://onsi.github.io/ginkgo
// useful commands https://ieftimov.com/posts/testing-in-go-go-test

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}
